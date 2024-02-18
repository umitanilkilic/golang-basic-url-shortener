package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/umitanilkilic/golang-basic-url-shortener/shortener"
	"github.com/umitanilkilic/golang-basic-url-shortener/store"
)

type URLShorteningRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func ShortenUrl(c *gin.Context) {
	var shorteningRequest URLShorteningRequest
	if err := c.ShouldBindJSON(&shorteningRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := fmt.Sprint(shortener.GenerateID())
	store.SaveMapping(shortUrl, shorteningRequest.LongUrl)

	host := "http://" + os.Getenv("APP_ADDRESS") + ":" + os.Getenv("APP_PORT") + "/s/"
	c.JSON(200, gin.H{
		"message":   "url created",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveLongUrl(shortUrl)

	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"LongURL": initialUrl,
	})
}
