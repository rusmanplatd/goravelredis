package facades

import (
	"github.com/rusmanplatd/goravelframework/contracts/cache"
	"github.com/rusmanplatd/goravelframework/contracts/queue"
	"github.com/rusmanplatd/goravelframework/contracts/session"

	redis "github.com/rusmanplatd/goravelredis"
)

func Cache(store string) (cache.Driver, error) {
	if redis.App == nil {
		return nil, redis.ErrRedisServiceProviderNotRegistered
	}
	if store == "" {
		return nil, redis.ErrRedisStoreIsRequired
	}

	instance, err := redis.App.MakeWith(redis.BindingCache, map[string]any{"store": store})
	if err != nil {
		return nil, err
	}

	return instance.(*redis.Cache), nil
}

func Queue(connection string) (queue.Driver, error) {
	if redis.App == nil {
		return nil, redis.ErrRedisServiceProviderNotRegistered
	}
	if connection == "" {
		return nil, redis.ErrRedisConnectionIsRequired
	}

	instance, err := redis.App.MakeWith(redis.BindingQueue, map[string]any{"connection": connection})
	if err != nil {
		return nil, err
	}

	return instance.(*redis.Queue), nil
}

func Session(driver string) (session.Driver, error) {
	if redis.App == nil {
		return nil, redis.ErrRedisServiceProviderNotRegistered
	}
	if driver == "" {
		return nil, redis.ErrRedisConnectionIsRequired
	}

	instance, err := redis.App.MakeWith(redis.BindingSession, map[string]any{"driver": driver})
	if err != nil {
		return nil, err
	}

	return instance.(*redis.Session), nil
}
