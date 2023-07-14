package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"time"
	"xiaomotou-website/db"
	"xiaomotou-website/model"
)

type Controller struct {
}

func (u Controller) ContactMe(c *gin.Context) {

	validate := validator.New()

	_ = validate.RegisterValidation("PhoneValidationErrors", PhoneValidationErrors)

	var r model.Contact
	err := c.ShouldBindJSON(&r)

	if err != nil {
		fmt.Println(err)
	}

	err = validate.Struct(r)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	r.CreateTime = time.Now()
	db.Db.Create(&r)

	c.JSON(http.StatusOK, r)

}

// 返回TRUE则不会报错，返回FALSE则会报错
func PhoneValidationErrors(fl validator.FieldLevel) bool {
	return CheckMobile(fl.Field().String())
}

// CheckMobile 检验手机号
func CheckMobile(phone string) bool {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)
}
