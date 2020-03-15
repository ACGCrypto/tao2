package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/taoblockchain/tao2/accounts/abi/bind"
	"github.com/taoblockchain/tao2/common"
	"github.com/taoblockchain/tao2/common/hexutil"
	"github.com/taoblockchain/tao2/contracts/trc2issuer"
	"github.com/taoblockchain/tao2/contracts/trc2issuer/simulation"
	"github.com/taoblockchain/tao2/ethclient"
	"log"
	"math/big"
	"time"
)

var (
	trc2TokenAddr = common.HexToAddress("0x80430A33EaB86890a346bCf64F86CFeAC73287f3")
)

func airDropTokenToAccountNoTao() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)
	mainAccount := bind.NewKeyedTransactor(simulation.MainKey)
	mainAccount.Nonce = big.NewInt(int64(nonce))
	mainAccount.Value = big.NewInt(0)      // in wei
	mainAccount.GasLimit = uint64(4000000) // in units
	mainAccount.GasPrice = big.NewInt(0).Mul(common.TRC2GasPrice,big.NewInt(2))
	trc2Instance, _ := trc2issuer.NewTRC2(mainAccount, trc2TokenAddr, client)
	trc2IssuerInstance, _ := trc2issuer.NewTRC2Issuer(mainAccount, common.TRC2IssuerSMC, client)
	// air drop token
	remainFee, _ := trc2IssuerInstance.GetTokenCapacity(trc2TokenAddr)
	tx, err := trc2Instance.Transfer(simulation.AirdropAddr, simulation.AirDropAmount)
	if err != nil {
		log.Fatal("can't air drop to ", err)
	}
	// check balance after transferAmount
	fmt.Println("wait 10s to airdrop success ", tx.Hash().Hex())
	time.Sleep(10 * time.Second)

	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPTRC2Fee.Uint64() {
		fee = fee.Mul(fee, common.TRC2GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err := trc2IssuerInstance.GetTokenCapacity(trc2TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	if err != nil {
		log.Fatal("can't execute transferAmount in tr21:", err)
	}
}
func testTransferTRC2TokenWithAccountNoTao() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}

	// access to address which received token trc1 but dont have tao
	nonce, _ := client.NonceAt(context.Background(), simulation.AirdropAddr, nil)
	airDropAccount := bind.NewKeyedTransactor(simulation.AirdropKey)
	airDropAccount.Nonce = big.NewInt(int64(nonce))
	airDropAccount.Value = big.NewInt(0)      // in wei
	airDropAccount.GasLimit = uint64(4000000) // in units
	airDropAccount.GasPrice = big.NewInt(0).Mul(common.TRC2GasPrice,big.NewInt(2))
	trc2Instance, _ := trc2issuer.NewTRC2(airDropAccount, trc2TokenAddr, client)
	trc2IssuerInstance, _ := trc2issuer.NewTRC2Issuer(airDropAccount, common.TRC2IssuerSMC, client)

	remainFee, _ := trc2IssuerInstance.GetTokenCapacity(trc2TokenAddr)
	airDropBalanceBefore, err := trc2Instance.BalanceOf(simulation.AirdropAddr)
	receiverBalanceBefore, err := trc2Instance.BalanceOf(simulation.ReceiverAddr)
	// execute transferAmount trc to other address
	tx, err := trc2Instance.Transfer(simulation.ReceiverAddr, simulation.TransferAmount)
	if err != nil {
		log.Fatal("can't execute transferAmount in tr21:", err)
	}

	// check balance after transferAmount
	fmt.Println("wait 10s to transferAmount success ")
	time.Sleep(10 * time.Second)

	balance, err := trc2Instance.BalanceOf(simulation.ReceiverAddr)
	wantedBalance := big.NewInt(0).Add(receiverBalanceBefore, simulation.TransferAmount)
	if err != nil || balance.Cmp(wantedBalance) != 0 {
		log.Fatal("check balance after fail receiverAmount in tr21: ", err, "get", balance, "wanted", wantedBalance)
	}

	remainAirDrop := big.NewInt(0).Sub(airDropBalanceBefore, simulation.TransferAmount)
	remainAirDrop = remainAirDrop.Sub(remainAirDrop, simulation.Fee)
	// check balance trc2 again
	balance, err = trc2Instance.BalanceOf(simulation.AirdropAddr)
	if err != nil || balance.Cmp(remainAirDrop) != 0 {
		log.Fatal("check balance after fail transferAmount in tr21: ", err, "get", balance, "wanted", remainAirDrop)
	}
	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPTRC2Fee.Uint64() {
		fee = fee.Mul(fee, common.TRC2GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err := trc2IssuerInstance.GetTokenCapacity(trc2TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check trc2 SMC balance
	balance, err = client.BalanceAt(context.Background(), common.TRC2IssuerSMC, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
}
func testTransferTrc21Fail() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.AirdropAddr, nil)
	airDropAccount := bind.NewKeyedTransactor(simulation.AirdropKey)
	airDropAccount.Nonce = big.NewInt(int64(nonce))
	airDropAccount.Value = big.NewInt(0)      // in wei
	airDropAccount.GasLimit = uint64(4000000) // in units
	airDropAccount.GasPrice = big.NewInt(0).Mul(common.TRC2GasPrice,big.NewInt(2))
	trc2Instance, _ := trc2issuer.NewTRC2(airDropAccount, trc2TokenAddr, client)
	trc2IssuerInstance, _ := trc2issuer.NewTRC2Issuer(airDropAccount, common.TRC2IssuerSMC, client)
	balanceIssuerFee, err := trc2IssuerInstance.GetTokenCapacity(trc2TokenAddr)

	minFee, err := trc2Instance.MinFee()
	if err != nil {
		log.Fatal("can't get minFee of trc2 smart contract:", err)
	}
	ownerBalance, err := trc2Instance.BalanceOf(simulation.MainAddr)
	remainFee, err := trc2IssuerInstance.GetTokenCapacity(trc2TokenAddr)
	airDropBalanceBefore, err := trc2Instance.BalanceOf(simulation.AirdropAddr)

	tx, err := trc2Instance.Transfer(common.Address{}, big.NewInt(1))
	if err != nil {
		log.Fatal("can't execute test transfer to zero address in tr21:", err)
	}
	fmt.Println("wait 10s to transfer to zero address")
	time.Sleep(10 * time.Second)

	fmt.Println("airDropBalanceBefore", airDropBalanceBefore)
	// check balance trc2 again
	airDropBalanceBefore = big.NewInt(0).Sub(airDropBalanceBefore, minFee)
	balance, err := trc2Instance.BalanceOf(simulation.AirdropAddr)
	if err != nil || balance.Cmp(airDropBalanceBefore) != 0 {
		log.Fatal("check balance after fail transferAmount in tr21: ", err, "get", balance, "wanted", airDropBalanceBefore)
	}

	ownerBalance = big.NewInt(0).Add(ownerBalance, minFee)
	//check balance fee
	balance, err = trc2Instance.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(ownerBalance) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPTRC2Fee.Uint64() {
		fee = fee.Mul(fee, common.TRC2GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err = trc2IssuerInstance.GetTokenCapacity(trc2TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check trc2 SMC balance
	balance, err = client.BalanceAt(context.Background(), common.TRC2IssuerSMC, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}

}
func main() {
	fmt.Println("========================")
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")

	start := time.Now()
	for i := 0; i < 10000000; i++ {
		airDropTokenToAccountNoTao()
		fmt.Println("Finish airdrop token to a account")
		testTransferTRC2TokenWithAccountNoTao()
		fmt.Println("Finish transfer trc2 token with a account no tao")
		testTransferTrc21Fail()
		fmt.Println("Finish testing ! Success transferAmount token trc1 with a account no tao")
	}
	fmt.Println(common.PrettyDuration(time.Since(start)))
}
