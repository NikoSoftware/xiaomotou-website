package controller

import (
	"fmt"
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
		fmt.Println(err)
	}

	err = tools.ValidateNew.Struct(r)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	r.CreateTime = time.Now()
	tools.Db.Create(&r)

	c.JSON(http.StatusOK, r)

}
