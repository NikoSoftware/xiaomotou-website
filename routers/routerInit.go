package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xiaomotou-website/controller"
)

func RouterInit(r *gin.Engine) {

	router := r.Group("/website")

	router.POST("/contactMe", controller.Controller{}.ContactMe)

}

func Recover(c *gin.Context) {
	// 加载defer异常处理
	defer func() {
		if err := recover(); err != nil {
			// 异常日志
			log.Printf("出现异常: %v\n", err)

			// 打印错误堆栈信息
			//	debug.PrintStack()

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
