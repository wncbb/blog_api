package main

import (
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"wncbb.cn/api/dal"
	handler_article "wncbb.cn/api/handler/article"
	handler_login "wncbb.cn/api/handler/login"
	handler_register "wncbb.cn/api/handler/register"
	handler_session "wncbb.cn/api/handler/session"
	handler_user "wncbb.cn/api/handler/user"
	middleware_cache "wncbb.cn/api/middleware/cache.v2"
	middleware_check_login "wncbb.cn/api/middleware/check_login"
	middleware_session "wncbb.cn/api/middleware/session"

	middleware_write_response "wncbb.cn/api/middleware/write_response"
	pb "wncbb.cn/api/protobuf"
	_ "wncbb.cn/api/validator"
	"wncbb.cn/config"
	model_article "wncbb.cn/db/model/article"
	model_user "wncbb.cn/db/model/user"
	"wncbb.cn/define"
	"wncbb.cn/log"
)

func main() {
	// init config
	config.Init()

	// init postgres
	dal.InitPostgres(config.GetPgDBConf())

	// init module
	model_user.Init()
	model_article.Init()

	r := gin.Default()

	// use session middleware
	r.Use(middleware_session.Session(define.SESSION_NAME, config.GetSessionRedisConfig()))

	// add monitor
	pprof.Register(r, "api/dev/pprof")

	// test session
	r.GET("/api/session/setname", handler_session.TestSetName())
	r.GET("/api/session/getname", handler_session.TestGetName())
	r.GET("/api/user/showme", handler_user.ShowMe())

	g := r.Group("/api")
	g.POST(
		"/register",
		middleware_write_response.WriteResponseMw(),
		handler_register.Register(),
	)
	g.POST(
		"/login",
		middleware_write_response.WriteResponseMw(),
		// middleware_cache.CacheMw(10*time.Second, "key", &pb.LoginResponse{}),
		handler_login.Login(),
	)
	g.GET(
		"/logout",
		middleware_write_response.WriteResponseMw(),
		handler_login.Logout(),
	)
	g.GET(
		"/articles/:articleId",
		middleware_write_response.WriteResponseMw(),
		middleware_cache.CacheMw(10*time.Second, "key", &pb.ArticleResponse{}),
		handler_article.GetById(),
	)
	g.POST(
		"/articles",
		middleware_write_response.WriteResponseMw(),
		middleware_check_login.CheckLogin(),
		handler_article.Create(),
	)

	log.DefaultLog().Debugf("Starting")
	r.Run(config.RunAddr())
}
