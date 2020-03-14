package waihui

import (
	"github.com/tao2-core/tao2-core/accounts/abi/bind"
	"github.com/tao2-core/tao2-core/common"
	"github.com/tao2-core/tao2-core/contracts/waihui/contract"
	"math/big"
)

type MyTRC2 struct {
	*contract.MyTRC2Session
	contractBackend bind.ContractBackend
}

func NewTRC2(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*MyTRC2, error) {
	smartContract, err := contract.NewMyTRC2(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &MyTRC2{
		&contract.MyTRC2Session{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployTRC2(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, name string, symbol string, decimals uint8, cap, fee *big.Int) (common.Address, *MyTRC2, error) {
	contractAddr, _, _, err := contract.DeployMyTRC2(transactOpts, contractBackend, name, symbol, decimals, cap, fee)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewTRC2(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
