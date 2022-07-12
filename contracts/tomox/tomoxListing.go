package janot

import (
	"github.com/janetachain/janetachain/accounts/abi/bind"
	"github.com/janetachain/janetachain/common"
	"github.com/janetachain/janetachain/contracts/janot/contract"
)

type JanoTListing struct {
	*contract.JanoTListingSession
	contractBackend bind.ContractBackend
}

func NewMyJanoTListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*JanoTListing, error) {
	smartContract, err := contract.NewJanoTListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &JanoTListing{
		&contract.JanoTListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployJanoTListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *JanoTListing, error) {
	contractAddr, _, _, err := contract.DeployJanoTListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyJanoTListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
