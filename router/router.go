package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-test/common"
	"go-gin-test/contoller/v1"
	"net/http"
	"net/url"
	"strconv"
)

func InitRouter(r *gin.Engine){
	r .GET("/sign", SignFunc)

	GroupV1 := r.Group("v1")
	{
		GroupV1.Any("product", common.VerifySign, v1.AddProduct)
		GroupV1.Any("member",common.VerifySign, v1.AddMember)
	}
	GroupV2 := r.Group("v2")
	{
		GroupV2.Any("product",common.VerifySign, v1.AddProduct)
		GroupV2.Any("member", common.VerifySign,v1.AddMember)
	}
}


func SignFunc(c *gin.Context){
	timestamp:= strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"timestamp": []string{timestamp},
	}

	res["sign"] = common.CreateSign(params)
	res["timestamp"] = timestamp
	common.ReturnJson(http.StatusOK, "", res, c)
}