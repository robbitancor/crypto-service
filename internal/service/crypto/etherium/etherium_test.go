package etherium

import (
	mock_service "gihub.com/robbitancor/simple-microservice/internal/service/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFindByAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_service.NewMockEtheriumService(ctrl)
	uri := "https://api.etherscan.io"
	module := "proxy"
	action := "eth_gasPrice"
	apiKey := "84TFIGZS8D2FH8Y66VAGS895PGPRTAR2VZ"

	m.EXPECT().GetGasPrice(uri+"/api", module, action, apiKey).Return("")
}
