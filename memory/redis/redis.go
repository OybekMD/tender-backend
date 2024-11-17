package memory

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redigo struct {
	rdb *redis.Client
}

func NewRedisInit(client *redis.Client) *Redigo {
	return &Redigo{
		rdb: client,
	}
}

func (r *Redigo) Set(key string, value interface{}) error {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}

	cmd := r.rdb.Set(context.Background(), key, str, 3*time.Minute)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *Redigo) Get(key string) (string, error) {
	str, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return str, nil
}

func (r *Redigo) Delete(key string) error {
	cmd := r.rdb.Del(context.Background(), key)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
