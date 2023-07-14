package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"xiaomotou-website/routers"
)

func main() {

	r := gin.Default()
	r.Use(Recover)
	routers.RouterInit(r)

	r.Run(":7300")

}

func Recover(c *gin.Context) {
	// 加载defer异常处理
	defer func() {
		if err := recover(); err != nil {
			// 异常日志
			log.Printf("出现异常: %v\n", err)

			// 打印错误堆栈信息
			debug.PrintStack()

			// 返回统一的Json风格
			c.JSON(http.StatusOK, gin.H{
				"code":    "500",
				"msg":     err,
				"success": false,
			})
			//终止后续操作
			c.Abort()
		}
	}()
	//继续操作
	c.Next()
}
