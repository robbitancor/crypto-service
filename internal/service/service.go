package service

import (
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	"gihub.com/robbitancor/simple-microservice/internal/storage"
)

type EtheriumService interface {
	GetBalance(uri, address, module, action, apiKey string) etherium.Balance
	GetBlockNumber(uri, module, action, apiKey string) etherium.Block
	GetGasPrice(uri, module, action, apiKey string) etherium.Gas
	SaveEtherium(eth etherium.Etherium) error
	CacheEtherium(eth etherium.Etherium) error
	SetRepo(repo storage.EtheriumRepository)
}
