package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Date: 2021/8/21 5:21 下午
 * @Desc:
 **/

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})

	engine.Run("0.0.0.0:9000")
}
