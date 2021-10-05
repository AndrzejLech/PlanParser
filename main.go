package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.New()
	router.LoadHTMLGlob("image.html")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", GetImage)
	router.GET("/inf-1-1", GetInfOneOne)
	router.GET("/inf-1-2", GetInfOneTwo)
	router.GET("/nur", GetNur)
	router.GET("/szymin", GetSzymin)
	router.GET("/mamlina", GetMamlina)
	err := router.Run()

	if err != nil {
		return
	}
}
