package main

import (
	"amitshekar-clean-arch/bootstrap"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.MySql

	server = gin.Default()
	basePath := server.Group("/v1")

	log.Fatal(server.Run(env.ServerAddress))
}
