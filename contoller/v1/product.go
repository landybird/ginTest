package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProduct(c *gin.Context){
	name := c.Query("name")
	price := c.DefaultQuery("price", "100")

	c.JSON(http.StatusCreated, gin.H{
		"v1"    : "AddProduct",
		"name"  : name,
		"price" : price,
	})
}