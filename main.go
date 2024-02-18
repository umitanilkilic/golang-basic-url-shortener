package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/umitanilkilic/golang-basic-url-shortener/handler"
	"github.com/umitanilkilic/golang-basic-url-shortener/store"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	godotenv.Load()

	appAddress := os.Getenv("APP_ADDRESS") + ":" + os.Getenv("APP_PORT")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	/// API Endpoints
	r.POST("/shorten-url", func(c *gin.Context) {
		handler.ShortenUrl(c)
	})

	r.GET("/s/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.ConnectToServer(os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD"), 0)

	err := r.Run(appAddress)
	if err != nil {
		panic("fail")
	}
}
