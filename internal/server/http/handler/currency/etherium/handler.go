package etherium

import (
	"gihub.com/robbitancor/simple-microservice/internal/config"
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	"gihub.com/robbitancor/simple-microservice/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"sync"
)

// Controller

type Handler struct {
	s service.EtheriumService
	a config.ApiConfig
}

func NewHandler(cfg config.ApiConfig, service service.EtheriumService) *Handler {
	return &Handler{s: service, a: cfg}
}

func (h *Handler) GetEtheriumBalance(c echo.Context) error {
	var reqDetails map[string]interface{}

	var balance etherium.Balance
	var block etherium.Block
	var gas etherium.Gas
	var ether etherium.Etherium

	var wg sync.WaitGroup

	_ = c.Bind(&reqDetails)

	address := reqDetails["address"].(string)
	wg.Add(1)
	go func(group *sync.WaitGroup) {
		defer wg.Done()
		balance = h.s.GetBalance(h.a.BaseUri, address, "account", "balance", h.a.Token)
		log.Printf("finished get balance.")
	}(&wg)

	wg.Add(1)
	go func(group *sync.WaitGroup) {
		defer wg.Done()
		block = h.s.GetBlockNumber(h.a.BaseUri, "proxy", "eth_blockNumber", h.a.Token)
		log.Printf("finished get block number.")

	}(&wg)

	wg.Add(1)
	go func(group *sync.WaitGroup) {
		defer wg.Done()
		gas = h.s.GetGasPrice(h.a.BaseUri, "proxy", "eth_gasPrice", h.a.Token)
		log.Printf("finished get gas price.")
	}(&wg)

	wg.Wait()

	ether.Gas = gas.Result
	ether.Balance = balance.Result
	ether.Block = block.Result

	// save to mongodb
	err := h.s.SaveEtherium(ether)

	if err != nil {
		log.Error(err)
		log.Error("error while saving to db", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// save to redis
	err = h.s.CacheEtherium(ether)
	if err != nil {
		log.Error("error while saving to cache", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ether)
}
