package handler

import (
	"url-shortener/internal/shortener"
	"url-shortener/internal/store"

	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserID)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserID)

	host := "http://localhost:9808"
	c.JSON(200, gin.H{
		"message":   "Short URL created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	originalUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, originalUrl)
}
