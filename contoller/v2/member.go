package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMember(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"v2" : "AddMember",
	})
}