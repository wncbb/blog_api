package captcha

import (
	api_define "blog_api/api/define"

	pb "blog_api/api/protobuf"
	service_captcha "blog_api/api/service/captcha"

	"github.com/gin-gonic/gin"
)

func GetCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := &pb.GetCaptchaResponse{}

		captchaData, err := service_captcha.GetCaptcha()
		if err != nil {
			resp.Code = pb.ResponseCode_GenRandomError
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		respData := &pb.GetCaptchaData{
			CaptchaData: captchaData.CaptchaData,
			CaptchaId:   captchaData.CaptchaId,
		}
		resp.Code = pb.ResponseCode_Success
		resp.Data = respData
		c.Set(api_define.CtxRespKey, resp)
	}
}

type VerifyCaptchaRequest struct {
	CaptchaCode string `form:"captcha_code" json:"captcha_code"`
	CaptchaId   string `form:"captcha_id" json:"captcha_id"`
}

func VerifyCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := &pb.VerifyCaptchaResponse{}
		resp.Code = pb.ResponseCode_Empty
		qry := &VerifyCaptchaRequest{}

		err := c.BindQuery(qry)
		if err != nil {
			resp.Code = pb.ResponseCode_QueryArgumentsError
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		verifyCaptchaData := &pb.VerifyCaptchaData{}

		showMsg, err := service_captcha.VerifyCaptcha(qry.CaptchaId, qry.CaptchaCode)
		if err != nil || showMsg != api_define.ShowMsgSuccess {
			verifyCaptchaData.ShowMsg = "验证码错误"
			resp.Data = verifyCaptchaData
			resp.Code = pb.ResponseCode_Success
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		verifyCaptchaData.ShowMsg = "success"
		resp.Data = verifyCaptchaData
		resp.Code = pb.ResponseCode_Success
		c.Set(api_define.CtxRespKey, resp)
		return
	}
}
