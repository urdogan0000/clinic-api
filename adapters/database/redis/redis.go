package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"main/pkg/env"
	"os"
)

var (
	RDB *redis.Client
	Ctx context.Context
	err error
)

func init() {
	println("Driver Redis Initialized")
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("pg_db_host"), os.Getenv("rd_db_port")),
		Password: os.Getenv("rd_db_pass"),
		DB:       env.EnvInt("rd_db_default"),
	})

	Ctx = context.Background()

}
