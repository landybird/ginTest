package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProduct(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"v2" : "AddProduct",
	})
}