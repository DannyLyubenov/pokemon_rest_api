package main

import (
	"fmt"

	"github.com/DannyLyubenov/pokemon_rest_api/pkg/apis/app"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("API is starting")
	router := gin.Default()
	router.GET("/api/v2/pokemon/:name", app.GetBasicInfo)
	router.GET("/api/v2/pokemon/translated/:name", app.GetFunDescription)
	router.Run(":80")
}
