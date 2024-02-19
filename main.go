// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"sample/src/controllers"
	"sample/src/models"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	models.Init()

	controllers.SetupRoutes(r)

	r.Run(":8080")
}
