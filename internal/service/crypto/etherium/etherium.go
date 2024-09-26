package etherium

import (
	"encoding/json"
	"fmt"
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	"gihub.com/robbitancor/simple-microservice/internal/server"
	storage2 "gihub.com/robbitancor/simple-microservice/internal/storage"
	"github.com/labstack/gommon/log"
)

type Service struct {
	c server.Client
	r *storage2.EtheriumRepository
}

func (e Service) SetRepo(repo storage2.EtheriumRepository) {
	e.r = &repo
}

func NewEtheriumService(client server.Client, repo storage2.EtheriumRepository) Service {
	return Service{r: &repo, c: client}
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

	return nil
}
