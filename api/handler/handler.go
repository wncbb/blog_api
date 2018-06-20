package handler

import (
	"fmt"

	"wncbb.cn/log"
)

func LogError(location string, err error, kv map[string]interface{}) {
	errInfo := fmt.Sprintf("location=[[%s]],err=[[%v]]", location, err)
	for k, v := range kv {
		errInfo += fmt.Sprintf(",%s=[[%v]]", k, v)
	}
	log.DefaultLog().Error(errInfo)
}
