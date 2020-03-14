package trc2issuer

import (
	"github.com/tao2-core/tao2-core/accounts/abi/bind"
	"github.com/tao2-core/tao2-core/accounts/abi/bind/backends"
	"github.com/tao2-core/tao2-core/common"
	"github.com/tao2-core/tao2-core/core"
	"github.com/tao2-core/tao2-core/crypto"
	"math/big"
	"testing"
)

var (
	mainKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	mainAddr   = crypto.PubkeyToAddress(mainKey.PublicKey)

	airdropKey, _ = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	airdropAddr   = crypto.PubkeyToAddress(airdropKey.PublicKey)

	subKey, _ = crypto.HexToECDSA("5bb98c5f937d176aa399ea6e6541f4db8f8db5a4ee1a8b56fb8beb41f2d755e3")
	subAddr   = crypto.PubkeyToAddress(subKey.PublicKey) //0x21292d56E2a8De3cC4672dB039AAA27f9190B1f6

	token = common.HexToAddress("0000000000000000000000000000000000000089")

	delay    = big.NewInt(30 * 48)
	minApply = big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(100000000000000000)) // 100 TOMO
)

func TestFeeTxWithTRC2Token(t *testing.T) {

	// init genesis
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		mainAddr: {Balance: big.NewInt(0).Mul(big.NewInt(10000000000000), big.NewInt(10000000000000))},
	})
	transactOpts := bind.NewKeyedTransactor(mainKey)
	// deploy payer swap SMC
	trc2IssuerAddr, trc2Issuer, err := DeployTRC2Issuer(transactOpts, contractBackend, minApply)

	//set contract address to config
	common.TRC2IssuerSMC = trc2IssuerAddr
	if err != nil {
		t.Fatal("can't deploy smart contract: ", err)
	}
	contractBackend.Commit()
	cap := big.NewInt(0).Mul(big.NewInt(10000000), big.NewInt(10000000000000))
	TRC2fee := big.NewInt(100)
	//  deploy a TRC2 SMC
	trc2TokenAddr, trc2, err := DeployTRC2(transactOpts, contractBackend, "TEST", "TOMO", 18, cap, TRC2fee)
	if err != nil {
		t.Fatal("can't deploy smart contract: ", err)
	}
	contractBackend.Commit()
	// add trc2 address to list token trc2Issuer
	trc2Issuer.TransactOpts.Value = minApply
	_, err = trc2Issuer.Apply(trc2TokenAddr)
	if err != nil {
		t.Fatal("can't add a token in  smart contract pay swap: ", err)
	}
	contractBackend.Commit()

	//check trc2 SMC balance
	balance, err := contractBackend.BalanceAt(nil, trc2IssuerAddr, nil)
	if err != nil || balance.Cmp(minApply) != 0 {
		t.Fatal("can't get balance  in trc2Issuer SMC: ", err, "got", balance, "wanted", minApply)
	}

	//check balance fee
	balanceIssuerFee, err := trc2Issuer.GetTokenCapacity(trc2TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(minApply) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", minApply)
	}
	trc2Issuer.TransactOpts.Value = big.NewInt(0)
	airDropAmount := big.NewInt(1000000000)
	// airdrop token trc2 to a address no tao
	tx, err := trc2.Transfer(airdropAddr, airDropAmount)
	if err != nil {
		t.Fatal("can't execute transfer in tr20: ", err)
	}
	contractBackend.Commit()
	receipt, err := contractBackend.TransactionReceipt(nil, tx.Hash())
	if err != nil {
		t.Fatal("can't transaction's receipt ", err, "hash", tx.Hash())
	}
	fee := big.NewInt(0).SetUint64(receipt.GasUsed)
	if receipt.Logs[0].BlockNumber > common.TIPTRC2Fee.Uint64() {
		fee = fee.Mul(fee, common.TRC2GasPrice)
	}
	remainFee := big.NewInt(0).Sub(minApply, fee)

	// check balance trc2 again
	balance, err = trc2.BalanceOf(airdropAddr)
	if err != nil || balance.Cmp(airDropAmount) != 0 {
		t.Fatal("check balance after fail transfer in tr20: ", err, "get", balance, "transfer", airDropAmount)
	}

	//check balance fee
	balanceIssuerFee, err = trc2Issuer.GetTokenCapacity(trc2TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check trc2 SMC balance
	balance, err = contractBackend.BalanceAt(nil, trc2IssuerAddr, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}

	// access to address which received token trc2 but dont have tao
	key1TransactOpts := bind.NewKeyedTransactor(airdropKey)
	key1Trc20, _ := NewTRC2(key1TransactOpts, trc2TokenAddr, contractBackend)

	transferAmount := big.NewInt(100000)
	// execute transfer trc to other address
	tx, err = key1Trc20.Transfer(subAddr, transferAmount)
	if err != nil {
		t.Fatal("can't execute transfer in tr20:", err)
	}
	contractBackend.Commit()

	balance, err = trc2.BalanceOf(subAddr)
	if err != nil || balance.Cmp(transferAmount) != 0 {
		t.Fatal("check balance after fail transfer in tr20: ", err, "get", balance, "transfer", transferAmount)
	}

	remainAirDrop := big.NewInt(0).Sub(airDropAmount, transferAmount)
	remainAirDrop = remainAirDrop.Sub(remainAirDrop, TRC2fee)
	// check balance trc2 again
	balance, err = trc2.BalanceOf(airdropAddr)
	if err != nil || balance.Cmp(remainAirDrop) != 0 {
		t.Fatal("check balance after fail transfer in tr20: ", err, "get", balance, "wanted", remainAirDrop)
	}

	receipt, err = contractBackend.TransactionReceipt(nil, tx.Hash())
	if err != nil {
		t.Fatal("can't transaction's receipt ", err, "hash", tx.Hash())
	}
	fee = big.NewInt(0).SetUint64(receipt.GasUsed)
	if receipt.Logs[0].BlockNumber > common.TIPTRC2Fee.Uint64() {
		fee = fee.Mul(fee, common.TRC2GasPrice)
	}
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err = trc2Issuer.GetTokenCapacity(trc2TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check trc2 SMC balance
	balance, err = contractBackend.BalanceAt(nil, trc2IssuerAddr, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		t.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
}
