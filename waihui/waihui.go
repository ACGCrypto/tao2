package waihui

import (
	"errors"
	"fmt"
	"github.com/taoblockchain/tao2/consensus"
	"github.com/taoblockchain/tao2/core/types"
	"github.com/taoblockchain/tao2/p2p"
	"github.com/taoblockchain/tao2/waihui/waihui_state"
	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
	"math/big"
	"strconv"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/taoblockchain/tao2/common"
	"github.com/taoblockchain/tao2/core/state"
	"github.com/taoblockchain/tao2/log"
	"github.com/taoblockchain/tao2/rpc"
	"golang.org/x/sync/syncmap"
)

const (
	ProtocolName       = "waihui"
	ProtocolVersion    = uint64(1)
	ProtocolVersionStr = "1.0"
	overflowIdx        // Indicator of message queue overflow
)

var (
	ErrNonceTooHigh = errors.New("nonce too high")
	ErrNonceTooLow  = errors.New("nonce too low")
)

type Config struct {
	DataDir        string `toml:",omitempty"`
	DBEngine       string `toml:",omitempty"`
	DBName         string `toml:",omitempty"`
	ConnectionUrl  string `toml:",omitempty"`
	ReplicaSetName string `toml:",omitempty"`
}

// DefaultConfig represents (shocker!) the default configuration.
var DefaultConfig = Config{
	DataDir: "",
}

type Waihui struct {
	// Order related
	db         OrderDao
	mongodb    OrderDao
	Triegc     *prque.Prque         // Priority queue mapping block numbers to tries to gc
	StateCache waihui_state.Database // State database to reuse between imports (contains state cache)    *waihui_state.WaihuiStateDB

	orderNonce map[common.Address]*big.Int

	sdkNode           bool
	settings          syncmap.Map // holds configuration settings that can be dynamically changed
	tokenDecimalCache *lru.Cache
	orderCache        *lru.Cache
}

func (waihui *Waihui) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (waihui *Waihui) Start(server *p2p.Server) error {
	return nil
}

func (waihui *Waihui) Stop() error {
	return nil
}

func NewLDBEngine(cfg *Config) *BatchDatabase {
	datadir := cfg.DataDir
	batchDB := NewBatchDatabaseWithEncode(datadir, 0)
	return batchDB
}

func NewMongoDBEngine(cfg *Config) *MongoDatabase {
	mongoDB, err := NewMongoDatabase(nil, cfg.DBName, cfg.ConnectionUrl, cfg.ReplicaSetName, 0)

	if err != nil {
		log.Crit("Failed to init mongodb engine", "err", err)
	}

	return mongoDB
}

func New(cfg *Config) *Waihui {
	tokenDecimalCache, _ := lru.New(defaultCacheLimit)
	orderCache, _ := lru.New(waihui_state.OrderCacheLimit)
	waihui := &Waihui{
		orderNonce:        make(map[common.Address]*big.Int),
		Triegc:            prque.New(),
		tokenDecimalCache: tokenDecimalCache,
		orderCache:        orderCache,
	}

	// default DBEngine: levelDB
	waihui.db = NewLDBEngine(cfg)
	waihui.sdkNode = false

	if cfg.DBEngine == "mongodb" { // this is an add-on DBEngine for SDK nodes
		waihui.mongodb = NewMongoDBEngine(cfg)
		waihui.sdkNode = true
	}

	waihui.StateCache = waihui_state.NewDatabase(waihui.db)
	waihui.settings.Store(overflowIdx, false)

	return waihui
}

// Overflow returns an indication if the message queue is full.
func (waihui *Waihui) Overflow() bool {
	val, _ := waihui.settings.Load(overflowIdx)
	return val.(bool)
}

func (waihui *Waihui) IsSDKNode() bool {
	return waihui.sdkNode
}

func (waihui *Waihui) GetDB() OrderDao {
	return waihui.db
}

func (waihui *Waihui) GetMongoDB() OrderDao {
	return waihui.mongodb
}

// APIs returns the RPC descriptors the Waihui implementation offers
func (waihui *Waihui) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewPublicWaihuiAPI(waihui),
			Public:    true,
		},
	}
}

// Version returns the Waihui sub-protocols version number.
func (waihui *Waihui) Version() uint64 {
	return ProtocolVersion
}

func (waihui *Waihui) ProcessOrderPending(coinbase common.Address, chain consensus.ChainContext, pending map[common.Address]types.OrderTransactions, statedb *state.StateDB, waihuistatedb *waihui_state.WaihuiStateDB) ([]waihui_state.TxDataMatch, map[common.Hash]waihui_state.MatchingResult) {
	txMatches := []waihui_state.TxDataMatch{}
	matchingResults := map[common.Hash]waihui_state.MatchingResult{}

	txs := types.NewOrderTransactionByNonce(types.OrderTxSigner{}, pending)
	for {
		tx := txs.Peek()
		if tx == nil {
			break
		}
		log.Debug("ProcessOrderPending start", "len", len(pending))
		log.Debug("Get pending orders to process", "address", tx.UserAddress(), "nonce", tx.Nonce())
		V, R, S := tx.Signature()

		bigstr := V.String()
		n, e := strconv.ParseInt(bigstr, 10, 8)
		if e != nil {
			continue
		}

		order := &waihui_state.OrderItem{
			Nonce:           big.NewInt(int64(tx.Nonce())),
			Quantity:        tx.Quantity(),
			Price:           tx.Price(),
			ExchangeAddress: tx.ExchangeAddress(),
			UserAddress:     tx.UserAddress(),
			BaseToken:       tx.BaseToken(),
			QuoteToken:      tx.QuoteToken(),
			Status:          tx.Status(),
			Side:            tx.Side(),
			Type:            tx.Type(),
			Hash:            tx.OrderHash(),
			OrderID:         tx.OrderID(),
			Signature: &waihui_state.Signature{
				V: byte(n),
				R: common.BigToHash(R),
				S: common.BigToHash(S),
			},
			PairName: tx.PairName(),
		}
		// make sure order is valid before running matching engine
		if err := order.VerifyOrder(statedb); err != nil {
			log.Error("waihui processOrderPending: invalid order", "err", err)
			continue
		}
		cancel := false
		if order.Status == OrderStatusCancelled {
			cancel = true
		}

		log.Info("Process order pending", "orderPending", order, "BaseToken", order.BaseToken.Hex(), "QuoteToken", order.QuoteToken)
		originalOrder := &waihui_state.OrderItem{}
		*originalOrder = *order
		originalOrder.Quantity = waihui_state.CloneBigInt(order.Quantity)

		if cancel {
			order.Status = OrderStatusCancelled
		}

		newTrades, newRejectedOrders, err := waihui.CommitOrder(coinbase, chain, statedb, waihuistatedb, waihui_state.GetOrderBookHash(order.BaseToken, order.QuoteToken), order)

		for _, reject := range newRejectedOrders {
			log.Debug("Reject order", "reject", *reject)
		}

		switch err {
		case ErrNonceTooLow:
			// New head notification data race between the transaction pool and miner, shift
			log.Debug("Skipping order with low nonce", "sender", tx.UserAddress(), "nonce", tx.Nonce())
			txs.Shift()
			continue

		case ErrNonceTooHigh:
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Debug("Skipping order account with high nonce", "sender", tx.UserAddress(), "nonce", tx.Nonce())
			txs.Pop()
			continue

		case nil:
			// everything ok
			txs.Shift()

		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
			txs.Shift()
			continue
		}

		// orderID has been updated
		originalOrder.OrderID = order.OrderID
		originalOrderValue, err := waihui_state.EncodeBytesItem(originalOrder)
		if err != nil {
			log.Error("Can't encode", "order", originalOrder, "err", err)
			continue
		}
		txMatch := waihui_state.TxDataMatch{
			Order: originalOrderValue,
		}
		txMatches = append(txMatches, txMatch)
		matchingResults[order.Hash] = waihui_state.MatchingResult{
			Trades:  newTrades,
			Rejects: newRejectedOrders,
		}
	}
	return txMatches, matchingResults
}

// there are 3 tasks need to complete to update data in SDK nodes after matching
// 1. txMatchData.Order: order has been processed. This order should be put to `orders` collection with status sdktypes.OrderStatusOpen
// 2. txMatchData.Trades: includes information of matched orders.
// 		a. PutObject them to `trades` collection
// 		b. Update status of regrading orders to sdktypes.OrderStatusFilled
func (waihui *Waihui) SyncDataToSDKNode(takerOrderInTx *waihui_state.OrderItem, txHash common.Hash, txMatchTime time.Time, statedb *state.StateDB, trades []map[string]string, rejectedOrders []*waihui_state.OrderItem, dirtyOrderCount *uint64) error {
	var (
		// originTakerOrder: order get from db, nil if it doesn't exist
		// takerOrderInTx: order decoded from txdata
		// updatedTakerOrder: order with new status, filledAmount, CreatedAt, UpdatedAt. This will be inserted to db
		originTakerOrder, updatedTakerOrder *waihui_state.OrderItem
		makerDirtyHashes                    []string
		makerDirtyFilledAmount              map[string]*big.Int
		err                                 error
	)
	db := waihui.GetMongoDB()
	sc := db.InitBulk()
	defer sc.Close()
	// 1. put processed takerOrderInTx to db
	lastState := waihui_state.OrderHistoryItem{}
	val, err := db.GetObject(takerOrderInTx.Hash, &waihui_state.OrderItem{})
	if err == nil && val != nil {
		originTakerOrder = val.(*waihui_state.OrderItem)
		lastState = waihui_state.OrderHistoryItem{
			TxHash:       originTakerOrder.TxHash,
			FilledAmount: waihui_state.CloneBigInt(originTakerOrder.FilledAmount),
			Status:       originTakerOrder.Status,
			UpdatedAt:    originTakerOrder.UpdatedAt,
		}
	}
	if originTakerOrder != nil {
		updatedTakerOrder = originTakerOrder
	} else {
		updatedTakerOrder = takerOrderInTx
	}

	if takerOrderInTx.Status != OrderStatusCancelled {
		updatedTakerOrder.Status = OrderStatusOpen
	} else {
		updatedTakerOrder.Status = OrderStatusCancelled
	}
	updatedTakerOrder.TxHash = txHash
	if updatedTakerOrder.CreatedAt.IsZero() {
		updatedTakerOrder.CreatedAt = txMatchTime
	}
	if txMatchTime.Before(updatedTakerOrder.UpdatedAt) || (txMatchTime.Equal(updatedTakerOrder.UpdatedAt) && *dirtyOrderCount == 0) {
		log.Debug("Ignore old orders/trades taker", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", updatedTakerOrder.UpdatedAt.UnixNano())
		return nil
	}
	*dirtyOrderCount++

	waihui.UpdateOrderCache(updatedTakerOrder.BaseToken, updatedTakerOrder.QuoteToken, updatedTakerOrder.Hash, txHash, lastState)
	updatedTakerOrder.UpdatedAt = txMatchTime

	// 2. put trades to db and update status to FILLED
	log.Debug("Got trades", "number", len(trades), "txhash", txHash.Hex())
	makerDirtyFilledAmount = make(map[string]*big.Int)
	for _, trade := range trades {
		// 2.a. put to trades
		tradeRecord := &Trade{}
		quantity := waihui_state.ToBigInt(trade[TradeQuantity])
		price := waihui_state.ToBigInt(trade[TradePrice])
		if price.Cmp(big.NewInt(0)) <= 0 || quantity.Cmp(big.NewInt(0)) <= 0 {
			return fmt.Errorf("trade misses important information. tradedPrice %v, tradedQuantity %v", price, quantity)
		}
		tradeRecord.Amount = quantity
		tradeRecord.PricePoint = price
		tradeRecord.PairName = updatedTakerOrder.PairName
		tradeRecord.BaseToken = updatedTakerOrder.BaseToken
		tradeRecord.QuoteToken = updatedTakerOrder.QuoteToken
		tradeRecord.Status = TradeStatusSuccess
		tradeRecord.Taker = updatedTakerOrder.UserAddress
		tradeRecord.Maker = common.HexToAddress(trade[TradeMaker])
		tradeRecord.TakerOrderHash = updatedTakerOrder.Hash
		tradeRecord.MakerOrderHash = common.HexToHash(trade[TradeMakerOrderHash])
		tradeRecord.TxHash = txHash
		tradeRecord.TakerOrderSide = updatedTakerOrder.Side
		tradeRecord.TakerExchange = updatedTakerOrder.ExchangeAddress
		tradeRecord.MakerExchange = common.HexToAddress(trade[TradeMakerExchange])

		// feeAmount: all fees are calculated in quoteToken
		quoteTokenQuantity := big.NewInt(0).Mul(quantity, price)
		quoteTokenQuantity = big.NewInt(0).Div(quoteTokenQuantity, common.BasePrice)
		takerFee := big.NewInt(0).Mul(quoteTokenQuantity, waihui_state.GetExRelayerFee(updatedTakerOrder.ExchangeAddress, statedb))
		takerFee = big.NewInt(0).Div(takerFee, common.WaihuiBaseFee)
		tradeRecord.TakeFee = takerFee

		makerFee := big.NewInt(0).Mul(quoteTokenQuantity, waihui_state.GetExRelayerFee(common.HexToAddress(trade[TradeMakerExchange]), statedb))
		makerFee = big.NewInt(0).Div(makerFee, common.WaihuiBaseFee)
		tradeRecord.MakeFee = makerFee

		// set makerOrderType, takerOrderType
		tradeRecord.MakerOrderType = trade[MakerOrderType]
		tradeRecord.TakerOrderType = updatedTakerOrder.Type

		if tradeRecord.CreatedAt.IsZero() {
			tradeRecord.CreatedAt = txMatchTime
		}
		tradeRecord.UpdatedAt = txMatchTime
		tradeRecord.Hash = tradeRecord.ComputeHash()

		log.Debug("TRADE history", "pairName", tradeRecord.PairName, "amount", tradeRecord.Amount, "pricepoint", tradeRecord.PricePoint,
			"taker", tradeRecord.Taker.Hex(), "maker", tradeRecord.Maker.Hex(), "takerOrder", tradeRecord.TakerOrderHash.Hex(), "makerOrder", tradeRecord.MakerOrderHash.Hex(),
			"takerFee", tradeRecord.TakeFee, "makerFee", tradeRecord.MakeFee)
		if err := db.PutObject(tradeRecord.Hash, tradeRecord); err != nil {
			return fmt.Errorf("SDKNode: failed to store tradeRecord %s", err.Error())
		}

		// 2.b. update status and filledAmount
		filledAmount := quantity
		// maker dirty order
		makerFilledAmount := big.NewInt(0)
		if amount, ok := makerDirtyFilledAmount[trade[TradeMakerOrderHash]]; ok {
			makerFilledAmount = waihui_state.CloneBigInt(amount)
		}
		makerFilledAmount.Add(makerFilledAmount, filledAmount)
		makerDirtyFilledAmount[trade[TradeMakerOrderHash]] = makerFilledAmount
		makerDirtyHashes = append(makerDirtyHashes, trade[TradeMakerOrderHash])

		//updatedTakerOrder = waihui.updateMatchedOrder(updatedTakerOrder, filledAmount, txMatchTime, txHash)
		//  update filledAmount, status of takerOrder
		updatedTakerOrder.FilledAmount.Add(updatedTakerOrder.FilledAmount, filledAmount)
		if updatedTakerOrder.FilledAmount.Cmp(updatedTakerOrder.Quantity) < 0 && updatedTakerOrder.Type == waihui_state.Limit {
			updatedTakerOrder.Status = OrderStatusPartialFilled
		} else {
			updatedTakerOrder.Status = OrderStatusFilled
		}
	}

	// update status for Market orders
	if updatedTakerOrder.Type == waihui_state.Market {
		if updatedTakerOrder.FilledAmount.Cmp(big.NewInt(0)) > 0 {
			updatedTakerOrder.Status = OrderStatusFilled
		} else {
			updatedTakerOrder.Status = OrderStatusRejected
		}
	}
	log.Debug("PutObject processed takerOrder",
		"pairName", updatedTakerOrder.PairName, "userAddr", updatedTakerOrder.UserAddress.Hex(), "side", updatedTakerOrder.Side,
		"price", updatedTakerOrder.Price, "quantity", updatedTakerOrder.Quantity, "filledAmount", updatedTakerOrder.FilledAmount, "status", updatedTakerOrder.Status,
		"hash", updatedTakerOrder.Hash.Hex(), "txHash", updatedTakerOrder.TxHash.Hex())
	if err := db.PutObject(updatedTakerOrder.Hash, updatedTakerOrder); err != nil {
		return fmt.Errorf("SDKNode: failed to put processed takerOrder. Hash: %s Error: %s", updatedTakerOrder.Hash.Hex(), err.Error())
	}
	makerOrders := db.GetListOrderByHashes(makerDirtyHashes)
	log.Debug("Maker dirty orders", "len", len(makerOrders), "txhash", txHash.Hex())
	for _, o := range makerOrders {
		if txMatchTime.Before(o.UpdatedAt) {
			log.Debug("Ignore old orders/trades maker", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", updatedTakerOrder.UpdatedAt.UnixNano())
			continue
		}
		lastState = waihui_state.OrderHistoryItem{
			TxHash:       o.TxHash,
			FilledAmount: waihui_state.CloneBigInt(o.FilledAmount),
			Status:       o.Status,
			UpdatedAt:    o.UpdatedAt,
		}
		waihui.UpdateOrderCache(o.BaseToken, o.QuoteToken, o.Hash, txHash, lastState)
		o.TxHash = txHash
		o.UpdatedAt = txMatchTime
		o.FilledAmount.Add(o.FilledAmount, makerDirtyFilledAmount[o.Hash.Hex()])
		if o.FilledAmount.Cmp(o.Quantity) < 0 {
			o.Status = OrderStatusPartialFilled
		} else {
			o.Status = OrderStatusFilled
		}
		log.Debug("PutObject processed makerOrder",
			"pairName", o.PairName, "userAddr", o.UserAddress.Hex(), "side", o.Side,
			"price", o.Price, "quantity", o.Quantity, "filledAmount", o.FilledAmount, "status", o.Status,
			"hash", o.Hash.Hex(), "txHash", o.TxHash.Hex())
		if err := db.PutObject(o.Hash, o); err != nil {
			return fmt.Errorf("SDKNode: failed to put processed makerOrder. Hash: %s Error: %s", o.Hash.Hex(), err.Error())
		}
	}

	// 3. put rejected orders to db and update status REJECTED
	log.Debug("Got rejected orders", "number", len(rejectedOrders), "rejectedOrders", rejectedOrders)

	if len(rejectedOrders) > 0 {
		var rejectedHashes []string
		// updateRejectedOrders
		for _, rejectedOrder := range rejectedOrders {
			rejectedHashes = append(rejectedHashes, rejectedOrder.Hash.Hex())
			if updatedTakerOrder.Hash == rejectedOrder.Hash && !txMatchTime.Before(updatedTakerOrder.UpdatedAt) {
				// cache order history for handling reorg
				orderHistoryRecord := waihui_state.OrderHistoryItem{
					TxHash:       updatedTakerOrder.TxHash,
					FilledAmount: waihui_state.CloneBigInt(updatedTakerOrder.FilledAmount),
					Status:       updatedTakerOrder.Status,
					UpdatedAt:    updatedTakerOrder.UpdatedAt,
				}
				waihui.UpdateOrderCache(updatedTakerOrder.BaseToken, updatedTakerOrder.QuoteToken, updatedTakerOrder.Hash, txHash, orderHistoryRecord)

				updatedTakerOrder.Status = OrderStatusRejected
				updatedTakerOrder.TxHash = txHash
				updatedTakerOrder.UpdatedAt = txMatchTime
				if err := db.PutObject(updatedTakerOrder.Hash, updatedTakerOrder); err != nil {
					return fmt.Errorf("SDKNode: failed to reject takerOrder. Hash: %s Error: %s", updatedTakerOrder.Hash.Hex(), err.Error())
				}
			}
		}
		dirtyRejectedOrders := db.GetListOrderByHashes(rejectedHashes)
		for _, order := range dirtyRejectedOrders {
			if txMatchTime.Before(order.UpdatedAt) {
				log.Debug("Ignore old orders/trades reject", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", updatedTakerOrder.UpdatedAt.UnixNano())
				continue
			}
			// cache order history for handling reorg
			orderHistoryRecord := waihui_state.OrderHistoryItem{
				TxHash:       order.TxHash,
				FilledAmount: waihui_state.CloneBigInt(order.FilledAmount),
				Status:       order.Status,
				UpdatedAt:    order.UpdatedAt,
			}
			waihui.UpdateOrderCache(order.BaseToken, order.QuoteToken, order.Hash, txHash, orderHistoryRecord)
			dirtyFilledAmount, ok := makerDirtyFilledAmount[order.Hash.Hex()]
			if ok && dirtyFilledAmount != nil {
				order.FilledAmount.Add(order.FilledAmount, dirtyFilledAmount)
			}
			order.Status = OrderStatusRejected
			order.TxHash = txHash
			order.UpdatedAt = txMatchTime
			if err = db.PutObject(order.Hash, order); err != nil {
				return fmt.Errorf("SDKNode: failed to update rejectedOder to sdkNode %s", err.Error())
			}
		}
	}

	if err := db.CommitBulk(); err != nil {
		return fmt.Errorf("SDKNode fail to commit bulk update orders, trades at txhash %s . Error: %s", txHash.Hex(), err.Error())
	}
	return nil
}

func (waihui *Waihui) GetWaihuiState(block *types.Block) (*waihui_state.WaihuiStateDB, error) {
	root, err := waihui.GetWaihuiStateRoot(block)
	if err != nil {
		return nil, err
	}
	if waihui.StateCache == nil {
		return nil, errors.New("Not initialized waihui")
	}
	return waihui_state.New(root, waihui.StateCache)
}

func (waihui *Waihui) GetStateCache() waihui_state.Database {
	return waihui.StateCache
}

func (waihui *Waihui) GetTriegc() *prque.Prque {
	return waihui.Triegc
}

func (waihui *Waihui) GetWaihuiStateRoot(block *types.Block) (common.Hash, error) {
	for _, tx := range block.Transactions() {
		if tx.To() != nil && tx.To().Hex() == common.WaihuiStateAddr {
			if len(tx.Data()) > 0 {
				return common.BytesToHash(tx.Data()), nil
			}
		}
	}
	return waihui_state.EmptyRoot, nil
}

func (waihui *Waihui) UpdateOrderCache(baseToken, quoteToken common.Address, orderHash common.Hash, txhash common.Hash, lastState waihui_state.OrderHistoryItem) {
	var orderCacheAtTxHash map[common.Hash]waihui_state.OrderHistoryItem
	c, ok := waihui.orderCache.Get(txhash)
	if !ok || c == nil {
		orderCacheAtTxHash = make(map[common.Hash]waihui_state.OrderHistoryItem)
	} else {
		orderCacheAtTxHash = c.(map[common.Hash]waihui_state.OrderHistoryItem)
	}
	orderKey := waihui_state.GetOrderHistoryKey(baseToken, quoteToken, orderHash)
	_, ok = orderCacheAtTxHash[orderKey]
	if !ok {
		orderCacheAtTxHash[orderKey] = lastState
	}
	waihui.orderCache.Add(txhash, orderCacheAtTxHash)
}

func (waihui *Waihui) RollbackReorgTxMatch(txhash common.Hash) {
	db := waihui.GetMongoDB()
	defer waihui.orderCache.Remove(txhash)

	for _, order := range db.GetOrderByTxHash(txhash) {
		c, ok := waihui.orderCache.Get(txhash)
		log.Debug("Waihui reorg: rollback order", "txhash", txhash.Hex(), "order", waihui_state.ToJSON(order))
		if !ok {
			log.Debug("Waihui reorg: remove order due to no orderCache", "order", waihui_state.ToJSON(order))
			if err := db.DeleteObject(order.Hash); err != nil {
				log.Error("SDKNode: failed to remove reorg order", "err", err.Error(), "order", waihui_state.ToJSON(order))
			}
			continue
		}
		orderCacheAtTxHash := c.(map[common.Hash]waihui_state.OrderHistoryItem)
		orderHistoryItem, _ := orderCacheAtTxHash[waihui_state.GetOrderHistoryKey(order.BaseToken, order.QuoteToken, order.Hash)]
		if (orderHistoryItem == waihui_state.OrderHistoryItem{}) {
			log.Debug("Waihui reorg: remove order due to empty orderHistory", "order", waihui_state.ToJSON(order))
			if err := db.DeleteObject(order.Hash); err != nil {
				log.Error("SDKNode: failed to remove reorg order", "err", err.Error(), "order", waihui_state.ToJSON(order))
			}
			continue
		}
		order.TxHash = orderHistoryItem.TxHash
		order.Status = orderHistoryItem.Status
		order.FilledAmount = waihui_state.CloneBigInt(orderHistoryItem.FilledAmount)
		order.UpdatedAt = orderHistoryItem.UpdatedAt
		log.Debug("Waihui reorg: update order to the last orderHistoryItem", "order", waihui_state.ToJSON(order), "orderHistoryItem", waihui_state.ToJSON(orderHistoryItem))
		if err := db.PutObject(order.Hash, order); err != nil {
			log.Error("SDKNode: failed to update reorg order", "err", err.Error(), "order", waihui_state.ToJSON(order))
		}
	}
	log.Debug("Waihui reorg: DeleteTradeByTxHash", "txhash", txhash.Hex())
	db.DeleteTradeByTxHash(txhash)

}
