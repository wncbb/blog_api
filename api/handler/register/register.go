package register

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"wncbb.cn/api/define"
	"wncbb.cn/api/handler"
	middleware_session "wncbb.cn/api/middleware/session"
	pb "wncbb.cn/api/protobuf"
	model_user "wncbb.cn/db/model/user"
)

type RegisterPost struct {
	Account  string `form:"account" json:"account" binding:"account"`
	Password string `form:"password" json:"password" binding:"password"`
}

func Register() gin.HandlerFunc {

	return func(c *gin.Context) {
		qry := &RegisterPost{}
		resp := &pb.RegisterResponse{}
		resp.Code = pb.ResponseCode_Empty
		var err error

		err = c.ShouldBindWith(qry, binding.Form)
		if err != nil {
			resp.Code = pb.ResponseCode_QueryArgumentsError
			resp.Msg = "request arguments illegal"
			handler.LogError("handler register", err, nil)
			c.Set(define.CtxRespKey, resp)
			return
		}

		s := middleware_session.Default(c)
		userId := s.Get(define.SessionUserIdKey)
		if userId != nil {
			resp.Code = pb.ResponseCode_ShouldLogoutFirst
			c.Set(define.CtxRespKey, resp)
			return
		}

		isExist, err := model_user.IsExistAccount(qry.Account)
		if err != nil {
			resp.Code = pb.ResponseCode_InternalError
			resp.Msg = "internal error"
			c.Set(define.CtxRespKey, resp)
			handler.LogError("handler register", err, nil)
			return
		}

		if isExist {
			resp.Code = 0
			resp.Data = &pb.RegisterData{
				ShowMsg: "用户名重复",
			}
			c.Set(define.CtxRespKey, resp)
			return
		}

		userCore := &model_user.UserCore{
			Account:  qry.Account,
			Password: qry.Password,
		}
		err = model_user.Create(userCore)
		if err != nil {
			resp.Code = 10001
			resp.Msg = "internal error"
			c.Set(define.CtxRespKey, resp)
			handler.LogError("handler register", err, nil)
			return
		}

		resp.Code = 0
		resp.Data = &pb.RegisterData{
			ShowMsg: "success",
			Id:      strconv.FormatInt(userCore.Id, 10),
		}
		c.Set(define.CtxRespKey, resp)
		return
	}
}
