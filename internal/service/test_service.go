package service

import (
	mock_service "gihub.com/robbitancor/simple-microservice/internal/service/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	_ = mock_service.NewMockEtheriumService(ctrl)
	//m.EXPECT().GetGasPrice(gomock)
}
