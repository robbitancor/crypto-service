package rsql

import (
	"database/sql"
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewEtheriumMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (e *MySQLRepository) Create(eth etherium.Etherium) error {
	e.db.Query("")
	return nil
}

func (e *MySQLRepository) Read(eth etherium.Etherium) (etherium.Etherium, error) {
	return etherium.Etherium{}, nil
}

func (e *MySQLRepository) Update(eth etherium.Etherium) error {
	return nil
}

func (e *MySQLRepository) Delete(eth etherium.Etherium) error {
	return nil
}
