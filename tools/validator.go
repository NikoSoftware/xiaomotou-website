package tools

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

func ValidateStruct(s interface{}) string {
	validate := validator.New()

	_ = validate.RegisterValidation("PhoneValidationErrors", PhoneValidationErrors)

	err := validate.Struct(s)
	return processErr(s, err)
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

func processErr(u interface{}, err error) string {
	if err == nil { //如果为nil 说明校验通过
		return ""
	}

	invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return "输入参数错误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field()                    //获取是哪个字段不符合格式
		field, ok := reflect.TypeOf(u).FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("reg_error_info") //获取field对应的reg_error_info tag值
			return fieldName + ":" + errorInfo           //返回错误
		} else {
			return "缺失reg_error_info"
		}
	}
	return ""
}
