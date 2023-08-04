package main

import (
	"amitshekar-clean-arch/api/route"
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

	route.Setup(env, db, server)

	log.Fatal(server.Run(env.ServerAddress))
}
