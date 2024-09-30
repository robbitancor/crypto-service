package etherium

import (
	"encoding/json"
	"fmt"
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	"gihub.com/robbitancor/simple-microservice/internal/server"
	repository "gihub.com/robbitancor/simple-microservice/internal/storage"
	"gihub.com/robbitancor/simple-microservice/internal/storage/redis"
	"github.com/labstack/gommon/log"
)

type Service struct {
	c  server.Client                 // http client
	r  repository.EtheriumRepository // db client (e.g.,mysql or mongodb)
	rc redis.RedisObject             //redis client
}

func (e Service) SetRepo(repo repository.EtheriumRepository) {
	e.r = repo
}

func NewEtheriumService(client server.Client, repo repository.EtheriumRepository, rc redis.RedisObject) Service {
	return Service{r: repo, c: client, rc: rc}
}

func (e Service) GetBalance(uri, address, module, action, apiKey string) etherium.Balance {
	var balance etherium.Balance
	val, err := e.c.GetRequest(fmt.Sprintf("%s/api?module=%s&action=%s&address=%s&apiKey=%s", uri, module, action, address, apiKey))
	if err != nil {
		log.Error(err)
		return balance
	}

	err = json.Unmarshal(val, &balance)

	if err != nil {
		log.Error(err)
		return balance
	}

	log.Printf("Retrieved value: %s", val)
	return balance
}

func (e Service) GetBlockNumber(uri, module, action, apiKey string) etherium.Block {
	val, err := e.c.GetRequest(fmt.Sprintf("%s/api?module=%s&action=%s&apiKey=%s", uri, module, action, apiKey))
	if err != nil {
		log.Error(err)
		return etherium.Block{}
	}

	var block etherium.Block
	err = json.Unmarshal(val, &block)

	if err != nil {
		log.Error(err)
		return etherium.Block{}
	}
	return block
}

func (e Service) GetGasPrice(uri, module, action, apiKey string) etherium.Gas {
	var gas etherium.Gas
	val, err := e.c.GetRequest(fmt.Sprintf("%s/api?module=%s&action=%s&apiKey=%s", uri, module, action, apiKey))
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(val, &gas)

	if err != nil {
		log.Error(err)
	}

	return gas
}
func (e Service) SaveEtherium(eth etherium.Etherium) error {
	err := e.r.Create(eth)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e Service) CacheEtherium(eth etherium.Etherium) error {
	err := e.rc.Create(eth)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
