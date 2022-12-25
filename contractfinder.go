package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractFinder struct {
	client      *ethclient.Client
	latestBlock int64
}

func NewContractFinder(provider string) (*ContractFinder, error) {
	conn, err := ethclient.DialContext(context.Background(), provider)
	if err != nil {
		return nil, err
	}
	latestBlock, err := conn.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	cf := new(ContractFinder)
	cf.latestBlock = latestBlock.Number().Int64()
	cf.client = conn
	return cf, nil
}

func (cf *ContractFinder) codeLen(contractAddr string, blockNumber int64) (int, error) {
	log.Println("Buscando o contrato", contractAddr, "no bloco", blockNumber)
	dados, err := cf.client.CodeAt(context.Background(), common.HexToAddress(contractAddr), big.NewInt(blockNumber))
	if err != nil {
		return 0, err
	}
	return len(dados), nil
}

func (cf *ContractFinder) getCreationBlock(contractAddr string, startBlock, endBlock int64) (int64, error) {
	if startBlock == endBlock {
		return startBlock, nil
	}

	midBlock := (startBlock + endBlock) / 2
	codeLen, err := cf.codeLen(contractAddr, midBlock)
	if err != nil {
		return 0, err
	}
	if codeLen > 2 {
		return cf.getCreationBlock(contractAddr, startBlock, midBlock)
	} else {
		return cf.getCreationBlock(contractAddr, midBlock+1, endBlock)
	}
}

func (cf *ContractFinder) GetContractCreationBlock(contractAddr string) (int64, error) {
	return cf.getCreationBlock(contractAddr, 1, cf.latestBlock)
}
