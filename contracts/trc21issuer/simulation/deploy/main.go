package main

import (
	"context"
	"fmt"
	"github.com/taoblockchain/tao2/accounts/abi/bind"
	"github.com/taoblockchain/tao2/common"
	"github.com/taoblockchain/tao2/contracts/trc2issuer"
	"github.com/taoblockchain/tao2/contracts/trc2issuer/simulation"
	"github.com/taoblockchain/tao2/ethclient"
	"log"
	"math/big"
	"time"
)

func main() {
	fmt.Println("========================")
	fmt.Println("mainAddr", simulation.MainAddr.Hex())
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)

	// init trc2 issuer
	auth := bind.NewKeyedTransactor(simulation.MainKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = big.NewInt(210000000000000)
	trc2IssuerAddr, trc2Issuer, err := trc2issuer.DeployTRC2Issuer(auth, client, simulation.MinApply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main address", simulation.MainAddr.Hex(), "nonce", nonce)
	fmt.Println("===> trc2 issuer address", trc2IssuerAddr.Hex())

	auth.Nonce = big.NewInt(int64(nonce + 1))

	// init trc1
	trc2TokenAddr, trc2Token, err := trc2issuer.DeployTRC2(auth, client, "TEST", "TOMO", 18, simulation.Cap, simulation.Fee)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===>  trc2 token address", trc2TokenAddr.Hex(), "cap", simulation.Cap)

	fmt.Println("wait 10s to execute init smart contract")
	time.Sleep(10 * time.Second)

	trc2Issuer.TransactOpts.Nonce = big.NewInt(int64(nonce + 2))
	trc2Issuer.TransactOpts.Value = simulation.MinApply
	trc2Issuer.TransactOpts.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	trc2Token.TransactOpts.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	trc2Token.TransactOpts.GasLimit = uint64(4000000)
	auth.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	// get balance init trc2 smart contract
	balance, err := trc2Token.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(simulation.Cap) != 0 {
		log.Fatal(err, "\tget\t", balance, "\twant\t", simulation.Cap)
	}
	fmt.Println("balance", balance, "mainAddr", simulation.MainAddr.Hex())

	// add trc1 list token trc2 issuer
	_, err = trc2Issuer.Apply(trc2TokenAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("wait 10s to add token to list issuer")
	time.Sleep(10 * time.Second)

	//check trc2 SMC balance
	balance, err = client.BalanceAt(context.Background(), trc2IssuerAddr, nil)
	if err != nil || balance.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance  in trc2Issuer SMC: ", err, "got", balance, "wanted", simulation.MinApply)
	}

	//check balance fee
	balanceIssuerFee, err := trc2Issuer.GetTokenCapacity(trc2TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", simulation.MinApply)
	}
}
