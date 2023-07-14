package routers

import (
	"github.com/gin-gonic/gin"
	"xiaomotou-website/controller"
)

func RouterInit(r *gin.Engine) {

	router := r.Group("/website")

	router.POST("/contactMe", controller.Controller{}.ContactMe)

}
