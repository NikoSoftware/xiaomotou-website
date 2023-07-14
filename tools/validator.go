package tools

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var ValidateNew *validator.Validate

// 初始化请求参数校验
func init() {
	validate := validator.New()

	_ = validate.RegisterValidation("PhoneValidationErrors", PhoneValidationErrors)
	ValidateNew = validate
}

// 返回TRUE则不会报错，返回FALSE则会报错
func PhoneValidationErrors(fl validator.FieldLevel) bool {
	return checkMobile(fl.Field().String())
}

// CheckMobile 检验手机号
func checkMobile(phone string) bool {
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
