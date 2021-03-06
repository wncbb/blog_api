package login

import (
	"errors"
	"strconv"

	api_define "blog_api/api/define"
	"blog_api/api/handler"
	middleware_session "blog_api/api/middleware/session"
	pb "blog_api/api/protobuf"
	service_captcha "blog_api/api/service/captcha"
	model_user "blog_api/db/model/user"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LoginPost struct {
	Account     string `form:"account" json:"account" binding:"account"`
	Password    string `form:"password" json:"password" binding:"password"`
	CaptchaId   string `form:"captchaId" json:"captchaId"`
	CaptchaCode string `form:"captchaCode" json:"captchaCode"`
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := &pb.BaseResponse{}
		s := middleware_session.Default(c)
		userId := s.Get(api_define.SessionUserIdKey)
		if userId == nil {
			resp.Code = pb.ResponseCode_ShouldLoginFirst
			c.Set(api_define.CtxRespKey, resp)
			handler.LogError("handler login", errors.New("unauthorized user call logout"), nil)
			return
		}
		s.Delete(api_define.SessionUserIdKey)
		s.Save()
		resp.Code = pb.ResponseCode_Success
		c.Set(api_define.CtxRespKey, resp)
		return
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		qry := &LoginPost{}
		resp := &pb.LoginResponse{}
		resp.Code = pb.ResponseCode_Empty
		var err error

		// Bind如果出错会直接处理错误,后面没办法处理了
		err = c.ShouldBindWith(qry, binding.JSON)

		handler.LogDebug("handler.login", err, map[string]interface{}{
			"qry": qry,
		})

		if err != nil {
			resp.Code = pb.ResponseCode_QueryArgumentsError
			c.Set(api_define.CtxRespKey, resp)
			handler.LogError("handler login", err, nil)
			return
		}

		s := middleware_session.Default(c)
		userId := s.Get(api_define.SessionUserIdKey)
		if userId != nil {
			resp.Code = pb.ResponseCode_ShouldLogoutFirst
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		showMsg, err := service_captcha.VerifyCaptcha(qry.CaptchaId, qry.CaptchaCode)
		if err != nil {
			resp.Code = pb.ResponseCode_VerifyCaptchaError
			c.Set(api_define.CtxRespKey, resp)
		}
		if showMsg != api_define.ShowMsgSuccess {
			resp.Code = pb.ResponseCode_Success
			resp.Data = &pb.LoginData{
				Msg: showMsg,
			}
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		userCore, err := model_user.GetByAccount(qry.Account)
		if err != nil {
			// 如果err不是记录未找到
			// resp.Code = pb.ResponseCode_InternalError
			// c.Set(api_define.CtxRespKey, resp)

			// 如果err是记录未找到
			resp.Data = &pb.LoginData{
				Msg: "用户名或密码错误",
			}
			resp.Code = pb.ResponseCode_Success
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		if userCore.Password != qry.Password {
			resp.Data = &pb.LoginData{
				Msg: "用户名或密码错误",
			}
			resp.Code = pb.ResponseCode_Success
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		resp.Code = pb.ResponseCode_Success
		resp.Msg = "success"
		resp.Data = &pb.LoginData{
			UserId: strconv.FormatInt(userCore.Id, 10),
			Msg:    "success",
		}

		// 设置session
		s.Set(api_define.SessionUserIdKey, userCore.Id)
		s.Save()

		// 存储的时候存proto.LoginData的指针
		c.Set(api_define.CtxRespKey, resp)
		// c.Set(api_define.CtxNeedCacheKey, true)
		return
	}
}
