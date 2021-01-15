package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-test/config"
	"go-gin-test/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	router.InitRouter(engine)
	_ = engine.Run(fmt.Sprintf(":%d", config.PORT))
}

// 68656c6c6f3f365c576d48240a477233de65040577     1610702851