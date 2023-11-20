// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"sample/controllers"
	"sample/models/person"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	person.Db_init()

	controllers.SetupRoutes(r)

	r.Run(":8080")
}

