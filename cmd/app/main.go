package main

import (
	"database/sql"
	_ "database/sql/driver"
	"gihub.com/robbitancor/simple-microservice/internal/config"
	eth_http "gihub.com/robbitancor/simple-microservice/internal/server/http"
	etherium2 "gihub.com/robbitancor/simple-microservice/internal/server/http/handler/currency/etherium"
	"gihub.com/robbitancor/simple-microservice/internal/service/crypto/etherium"
	rmongo "gihub.com/robbitancor/simple-microservice/internal/storage/mongodb/crypto/etherium"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	c := http.Client{}
	cfg := loadConfigs()
	initRedis()
	initLogger(e)
	cc := initMongo(cfg.DbConfig)
	nosql := rmongo.NewEtheriumMongoRepository(cc)
	client := eth_http.NewClient(cfg.ApiConfig, &c)
	s := etherium.NewEtheriumService(client, nosql)
	h := etherium2.NewHandler(cfg.ApiConfig, s)

	ethGroup := e.Group("/etherium")

	// Secure access to api
	ethGroup.Use(middleware.Secure())

	ethGroup.POST("/getBalance", h.GetEtheriumBalance)

	port := ":" + cfg.ServerConfig.Port
	e.Logger.Fatal(e.Start(port))
}

func initDb(dbConfig config.DbConfig) *sql.DB {
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Error(err)
	}

	return db
}

func initMongo(dbConfig config.DbConfig) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Error(err)
	}

	return client
}

func initLogger(e *echo.Echo) {
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
}

func loadConfigs() config.ConfigRoot {
	yamlData, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Error(err)
	}

	var configRoot config.ConfigRoot

	err = yaml.Unmarshal(yamlData, &configRoot)
	if err != nil {
		log.Error(err)
	}

	return configRoot
}

func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return client
}
