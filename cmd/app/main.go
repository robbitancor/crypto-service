package main

import (
	"context"
	"database/sql"
	_ "database/sql/driver"
	"fmt"
	"gihub.com/robbitancor/simple-microservice/internal/config"
	eth_http "gihub.com/robbitancor/simple-microservice/internal/server/http"
	handler "gihub.com/robbitancor/simple-microservice/internal/server/http/handler/currency/etherium"
	service "gihub.com/robbitancor/simple-microservice/internal/service/crypto/etherium"
	rmongo "gihub.com/robbitancor/simple-microservice/internal/storage/mongodb/crypto/etherium"
	rredis "gihub.com/robbitancor/simple-microservice/internal/storage/redis"
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
	ctx := context.Background()
	e := echo.New()
	c := http.Client{}

	initServer(e)
	cfg := loadConfigs()
	r := initRedis(ctx, cfg.RedisConfig)
	cc := initMongo(cfg.DbConfig)
	client := eth_http.NewClient(cfg.ApiConfig, &c)
	rc := rredis.NewRedisRepository(ctx, r)
	nosql := rmongo.NewEtheriumMongoRepository(cc)
	s := service.NewEtheriumService(client, nosql, rc)
	h := handler.NewHandler(cfg.ApiConfig, s)

	ethGroup := e.Group("/etherium")

	// add security
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

func initMongo(cfg config.DbConfig) *mongo.Client {
	bsonOpts := &options.BSONOptions{
		NilSliceAsEmpty: true,
	}

	//user := dbConfig.Username
	//pass := dbConfig.Password
	host := cfg.Host
	port := cfg.Port
	//replica := dbConfig.Replica

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	client, err := mongo.Connect(options.Client().ApplyURI(uri).
		SetBSONOptions(bsonOpts))

	if err != nil {
		log.Error(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Error(err)
	}

	return client
}

func initServer(e *echo.Echo) {
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	e.Use(middleware.Logger())
}

func loadConfigs() config.ConfigRoot {
	// load configuration

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

func initRedis(ctx context.Context, cfg config.RedisConfig) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	return client
}
