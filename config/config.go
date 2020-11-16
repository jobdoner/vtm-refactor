package config

import (
	"os"
	"strconv"
)

const (
	defaultPostgres = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	defaultPort     = 8080
	defaultSecret   = "secret"
	defaultRedis = "redis://:secret@localhost:6379/1"
)

//nolint:lll,maligned
type (
	Config struct {
		API API
		DB  DB
	}

	DB struct {
		Postgres Postgres
		Redis    Redis
	}

	Postgres struct {
		URI        string `desc:"URI of Postgres Server"`
		Migrations int
	}
	Redis struct {
		URI        string `desc:"URI of Redis Server"`
		Migrations int
	}

	API struct {
		Port            int    `default:"8080" desc:"Port to start HTTP server"`
		Secret          string `default:"" desc:"Secret key to verify sign"`
		ShutdownTimeout int    `split_words:"true" default:"10" desc:"Timeout to shutdown the server"`
	}
)

func Configure() (c Config, err error) {

	c.API.Port = defaultPort
	c.API.Secret = defaultSecret
	c.DB.Postgres.URI = defaultPostgres
	c.DB.Postgres.Migrations = 1
	c.DB.Redis.URI=defaultRedis
	if uri, ok := os.LookupEnv("URI"); ok {
		c.DB.Postgres.URI = uri
	}
	if secret, ok := os.LookupEnv("SECRET"); ok {
		c.API.Secret = secret
	}
	if port, ok := os.LookupEnv("PORT"); ok {
		c.API.Port, err = strconv.Atoi(port)
		if err != nil {
			panic(err)
		}
	}
	if mig, ok := os.LookupEnv("MIG"); ok {
		c.DB.Postgres.Migrations, err = strconv.Atoi(mig)
		if err != nil {
			panic(err)
		}
	}
	if ruri, ok := os.LookupEnv("REDISURI"); ok {
		c.DB.Redis.URI = ruri
	}


	return
}
