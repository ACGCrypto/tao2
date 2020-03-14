package waihui

import (
	"context"
	"errors"
	"github.com/tao2-core/tao2-core/waihui/waihui_state"
	"math/big"
	"sync"
	"time"

	"github.com/tao2-core/tao2-core/common"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicWaihuiAPI provides the waihui RPC service that can be
// use publicly without security implications.
type PublicWaihuiAPI struct {
	t        *Waihui
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicWaihuiAPI create a new RPC waihui service.
func NewPublicWaihuiAPI(t *Waihui) *PublicWaihuiAPI {
	api := &PublicWaihuiAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Waihui sub-protocol version.
func (api *PublicWaihuiAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}

// GetOrderNonce returns the latest orderNonce of the given address
func (api *PublicWaihuiAPI) GetOrderNonce(address common.Address) (*big.Int, error) {
	//TODO: getOrderNonce from state
	return big.NewInt(0), nil
}

// GetPendingOrders returns pending orders of the given pair
func (api *PublicWaihuiAPI) GetPendingOrders(pairName string) ([]*waihui_state.OrderItem, error) {
	result := []*waihui_state.OrderItem{}
	//TODO: get pending orders from orderpool
	return result, nil
}
