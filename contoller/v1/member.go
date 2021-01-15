package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-test/common"
	"net/http"
)

func AddMember(c *gin.Context){
	common.VerifySign(c)
	name := c.Query("name")
	price := c.DefaultQuery("price", "100")
	c.JSON(http.StatusCreated, gin.H{
		"v1"    : "AddMember",
		"name"  : name,
		"price" : price,
	})
}