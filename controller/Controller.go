package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"xiaomotou-website/model"
	"xiaomotou-website/tools"
)

type Controller struct {
}

func (u Controller) ContactMe(c *gin.Context) {

	var r model.Contact
	err := c.ShouldBindJSON(&r)

	if err != nil {
		panic(err)
	}
	errorInfo := tools.ValidateStruct(r)
	if len(errorInfo) != 0 {

		panic(errorInfo)
	}
	r.CreateTime = time.Now()
	tools.Db.Create(&r)

	c.JSON(http.StatusOK, r)

}
