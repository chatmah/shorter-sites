package main

import (
	"shorter-sites/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	router.POST("/shortenUrl", ConvertLinkHandler)
	router.GET("/shortenUrl", RetriveLinkHandler)
	router.Run(":8081")
}
