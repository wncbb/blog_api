package validator

import (
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"
)

func age(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	var age int32
	var ok bool
	if age, ok = field.Interface().(int32); !ok {
		return false
	}
	if age < 0 || age > 100 {
		return false
	}
	return true
}

var phonePattern string = `^1[0-9]{10}$`
var mailPattern string = `^[[:word:]]+(\.[[:word:]]+)?@[[:word:]]+(\.[[:word:]]+)+$`
var passwordPattern string = `^[[:alnum:]]+$`

var phoneReg *regexp.Regexp
var mailReg *regexp.Regexp
var passwordReg *regexp.Regexp

func init() {
	var err error
	phoneReg, err = regexp.Compile(phonePattern)
	if err != nil {
		panic(err)
	}
	mailReg, err = regexp.Compile(mailPattern)
	if err != nil {
		panic(err)
	}
	passwordReg, err = regexp.Compile(passwordPattern)
	if err != nil {
		panic(err)
	}
}

func account(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	var account string
	var ok bool
	if account, ok = field.Interface().(string); !ok {
		return false
	}
	if len(account) <= 0 && len(account) > 255 {
		return false
	}
	if match := phoneReg.MatchString(account); match {
		return true
	}

	if match := mailReg.MatchString(account); match {
		return true
	}

	return false
}

func password(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	var password string
	var ok bool
	if password, ok = field.Interface().(string); !ok {
		return false
	}
	if len(password) < 12 || len(password) > 255 {
		return false
	}
	if match := passwordReg.MatchString(password); match {
		return true
	}
	return false
}

func articleTitle(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	var articleTitle string
	var ok bool
	if articleTitle, ok = field.Interface().(string); !ok {
		return false
	}
	if len(articleTitle) <= 10 || len(articleTitle) > 255 {
		return false
	}
	return false
}

func init() {
	var v *validator.Validate
	var ok bool

	if v, ok = binding.Validator.Engine().(*validator.Validate); !ok {
		panic("init validator failed")
	}
	v.RegisterValidation("age", age)
	v.RegisterValidation("account", account)
	v.RegisterValidation("password", password)
	v.RegisterValidation("articleTitle", articleTitle)

}
