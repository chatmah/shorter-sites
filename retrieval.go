package main

import (
	"fmt"
	"net/http"

	"shorter-sites/handler"

	"github.com/gin-gonic/gin"
)

func RetriveLinkHandler(c *gin.Context) {
	fmt.Println("Am i even getting called")
	// aaaaaaaag
	encodedLink := c.Query("encodedUrl")

	res, err := handler.ReadHandler(encodedLink)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error retrieving link: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"decoded_url": res.LongURL})
}
