package main

import (
	"github.com/gin-gonic/gin"
	"xiaomotou-website/routers"
)

func main() {

	r := gin.Default()
	routers.RouterInit(r)

	r.Run(":7300")

}
