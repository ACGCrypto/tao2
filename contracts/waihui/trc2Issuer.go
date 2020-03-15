package waihui

import (
	"github.com/taoblockchain/tao2/accounts/abi/bind"
	"github.com/taoblockchain/tao2/common"
	"github.com/taoblockchain/tao2/contracts/waihui/contract"
	"math/big"
)

type TRC2Issuer struct {
	*contract.TRC2IssuerSession
	contractBackend bind.ContractBackend
}

func NewTRC2Issuer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*TRC2Issuer, error) {
	contractObject, err := contract.NewTRC2Issuer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &TRC2Issuer{
		&contract.TRC2IssuerSession{
			Contract:     contractObject,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployTRC2Issuer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, minApply *big.Int) (common.Address, *TRC2Issuer, error) {
	contractAddr, _, _, err := contract.DeployTRC2Issuer(transactOpts, contractBackend, minApply)
	if err != nil {
		return contractAddr, nil, err
	}
	contractObject, err := NewTRC2Issuer(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, contractObject, nil
}
