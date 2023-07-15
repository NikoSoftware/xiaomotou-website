package main

import (
	"github.com/gin-gonic/gin"
	"xiaomotou-website/routers"
)

func main() {

	r := gin.Default()
	r.Use(routers.Recover)
	routers.RouterInit(r)

	r.Run(":7300")

}
