package main

import (
	"log"
	gormdb "project2/gormDB"
	"project2/routers"

	"github.com/gin-gonic/gin"
)

func loadDataBase() {
	gormdb.Connection()
}

func main() {
	loadDataBase()

	router := gin.Default()

	routers.AuthRouter(router)

	log.Fatal(router.Run(":8080"))
}
