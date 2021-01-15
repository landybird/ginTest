package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Recover() gin.HandlerFunc{
	return func(c *gin.Context){
		defer func(){
			if r:= recover(); r!= nil {
				logrus.Info(r)
			}
		}()
		c.Next()
	}
}