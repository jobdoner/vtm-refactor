package main

import (
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
	"github.com/jobdoner/vtm-refactor/config"
	"github.com/jobdoner/vtm-refactor/migrations"
	"github.com/jobdoner/vtm-refactor/restapi"
	"github.com/jobdoner/vtm-refactor/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		panic(err)
	}

	config, err := config.Configure()

	if err != nil {
		fmt.Println(err)
	}

	opt, err := pg.ParseURL(config.DB.Postgres.URI)

	if err != nil {
		fmt.Println(err)
	}

	redisOpts, err := redis.ParseURL(config.DB.Redis.URI)

	if err != nil {
		fmt.Println(err)
	}

	migrations.Migrate(config.DB.Postgres.URI, config.DB.Postgres.Migrations, config.DB.Redis.URI)
	restapi.Init(pg.Connect(opt), redis.NewClient(redisOpts), config.API.Secret)

	api := operations.NewVtmRefactorAPI(swaggerSpec)

	server := restapi.NewServer(api)
	server.ConfigureAPI()
	server.Port = config.API.Port
	if err = server.Serve(); err != nil {
		panic(err)
	}

}
