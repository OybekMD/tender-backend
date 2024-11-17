package main

import (
	"tender/api"
	"tender/config"
	"tender/memory"
	"tender/storage"

	"fmt"
	"log"
	"os"

	cashe "tender/memory/redis"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.Load()

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect to postgresql database: %v", err)
	}

	str := storage.NewStoragePg(psqlConn)

	rediClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s%s", cfg.RedisHost, cfg.RedisPort),
		DB:       cast.ToInt(cfg.RedisDatabase),
		Password: cfg.RedisPassword,
	})

	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Fatalf("failed to create casbin enforcer: %v", err)
	}

	memory.Init(cashe.NewRedisInit(rediClient))

	logFile, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	apiServer := api.New(&api.Option{
		Conf:     &cfg,
		Storage:  str,
		Enforcer: casbinEnforcer,
	})
	fmt.Println("Server runned", cfg.HttpPort)
	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run http server: %v", err)
	}
}
