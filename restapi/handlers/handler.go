package handlers

import (
	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
)

//TODO: add interface to dbs
type Handler struct {
	DB        *pg.DB
	RedisConn *redis.Client
}
