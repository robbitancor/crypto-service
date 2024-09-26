package redis

import (
	"context"
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
)

type RedisObject struct {
	ctx context.Context
	rc  *redis.Client
}

func NewRedisRepository(ctx context.Context, client *redis.Client) RedisObject {
	return RedisObject{ctx: ctx, rc: client}
}

func (r RedisObject) Create(eth etherium.Etherium) error {
	// for simplicity's sake

	_ = r.rc.Set(r.ctx, eth.Block, eth, 0)

	return nil
}

func (r RedisObject) Read(eth etherium.Etherium) (etherium.Etherium, error) {
	val, err := r.rc.Get(r.ctx, eth.Block).Result()

	if err != nil {
		log.Error(err)
		return etherium.Etherium{}, err
	}

	var ethBalance etherium.Etherium
	eth.Balance = val
	return ethBalance, nil
}

func (r RedisObject) Update(eth etherium.Etherium) error {
	return nil
}

func (r RedisObject) Delete(eth etherium.Etherium) error {
	return nil
}
