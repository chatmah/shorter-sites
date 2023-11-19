package main

import (
	"net/http"
	"shorter-sites/encoder"
	"shorter-sites/handler"
	"shorter-sites/model"

	"github.com/gin-gonic/gin"
)

// Converts the link
func ConvertLinkHandler(c *gin.Context) {
	var shortLinkReq model.ShortLinkRequest
	if err := c.BindJSON(&shortLinkReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}
	linkToShort := shortLinkReq.LinkToShort
	result, err := encoder.EncodeLinkToShort(linkToShort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error processing link creation: " + err.Error()})
		return
	}
	obj := model.UrlData{
		LongURL:   linkToShort,
		ShortCode: result,
	}
	if err := handler.CreateHandler(c.Writer, c.Request, obj); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error processing link creation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"encoded_url": result})
}
