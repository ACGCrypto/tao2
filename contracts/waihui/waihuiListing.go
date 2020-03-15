package waihui

import (
	"github.com/taoblockchain/tao2/accounts/abi/bind"
	"github.com/taoblockchain/tao2/common"
	"github.com/taoblockchain/tao2/contracts/waihui/contract"
)

type WAIHUIListing struct {
	*contract.WAIHUIListingSession
	contractBackend bind.ContractBackend
}

func NewMyWAIHUIListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*WAIHUIListing, error) {
	smartContract, err := contract.NewWAIHUIListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &WAIHUIListing{
		&contract.WAIHUIListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployWAIHUIListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *WAIHUIListing, error) {
	contractAddr, _, _, err := contract.DeployWAIHUIListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyWAIHUIListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
