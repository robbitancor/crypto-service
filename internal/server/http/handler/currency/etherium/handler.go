package etherium

import (
	"gihub.com/robbitancor/simple-microservice/internal/config"
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	"gihub.com/robbitancor/simple-microservice/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
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
	var eth etherium.Etherium

	c.Bind(&reqDetails)
	address := reqDetails["address"].(string)

	balance = h.s.GetBalance(h.a.BaseUri, address, "account", "balance", h.a.Token)
	block = h.s.GetBlockNumber(h.a.BaseUri, "proxy", "eth_blockNumber", h.a.Token)
	gas = h.s.GetGasPrice(h.a.BaseUri, "proxy", "eth_gasPrice", h.a.Token)

	eth.Gas = gas.Result
	eth.Balance = balance.Result
	eth.Block = block.Result

	return c.JSON(http.StatusOK, eth)
}
