package storage

import (
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
)

type EtheriumRepository interface {
	Create(eth etherium.Etherium) error
	Read(eth etherium.Etherium) (etherium.Etherium, error)
	Update(eth etherium.Etherium) error
	Delete(eth etherium.Etherium) error
}
