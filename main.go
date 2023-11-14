package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortLinkRequest struct {
	LinkToShort string `json:"linkToShort"`
}

type responseLink struct {
	URL string `json:"short_url"`
}

var foo = []responseLink{
	{
		URL: "test-response",
	},
}

var counter = 0

func generateShortURL() string {
	counter++
	return fmt.Sprintf("short%d", counter)
}

func convertLink(c *gin.Context) {
	var shortLinkReq ShortLinkRequest
	if err := c.BindJSON(&shortLinkReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}

	linkToShort := shortLinkReq.LinkToShort

	fmt.Printf("Received linkToShort: %s\n", linkToShort)

	c.JSON(http.StatusOK, gin.H{"message": "Received the link to short val successfully"})
}

func main() {
	router := gin.Default()
	// router.GET("/shorten/:slug", convertLink)
	router.POST("/shortenUrl", convertLink)
	fmt.Println("Hello World")
	router.Run(":8081")
}
