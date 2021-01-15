package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-test/config"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

func Print(i interface{}){
	fmt.Println("======")
	fmt.Println(i)
	fmt.Println("======")
}



func ReturnJson(code int,  msg string,  data interface{}, c *gin.Context){
	c.JSON(code, gin.H{
		"data": data,
		"msg": msg,
	})
	c.Abort()
}


func GetTimeUnix() int64{
	return time.Now().Unix()
}

func MD5(str string) string{
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum([]byte("hello")))
}


func CreateSign(params url.Values) string{
	var key []string
	var str = ""
	for k:= range params{
		if k != "sign" {
			key = append(key, k)
		}

	}
	sort.Strings(key)
	fmt.Println(key)
	for i :=0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("v%=%v", key[i], params.Get(key[i]))
		}else{
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}

	}
	sign := MD5(MD5(str) + MD5(config.APP_NAME + config.APP_SECRET))
	return sign
}

func VerifySign(c *gin.Context){
	var method = c.Request.Method
	var timestamp int64
	var sign string
	var request url.Values

	switch{
		case method == "GET":
			request = c.Request.URL.Query()
			sign = c.Query("sign")
			timestamp, _ = strconv.ParseInt(c.Query("timestamp"), 10, 64)
		case method == "POST":
			err := c.Request.ParseForm()
			if err != nil {
				ReturnJson(http.StatusInternalServerError, "Bad requests", "", c)
			}
			request = c.Request.PostForm
			sign = c.PostForm("sign")
			timestamp, _ = strconv.ParseInt(c.PostForm("timestamp"), 10, 64)
		default:
			ReturnJson(http.StatusInternalServerError, "Bad requests", "", c)
	}
	exp, _ := strconv.ParseInt(config.APP_EXPIRY, 10, 64)

	if timestamp > GetTimeUnix() || GetTimeUnix() - timestamp >= exp {
		ReturnJson(http.StatusInternalServerError, "timestamp error", "", c)
	}else{
		if sign == "" || sign != CreateSign(request){
			ReturnJson(http.StatusInternalServerError, "sign error", "", c)
		}
	}


}

//  1610702619  68656c6c6fe126e338199266af1e58c1f08fa678ab

