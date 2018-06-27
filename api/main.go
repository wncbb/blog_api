package main

import (
	"blog_api/api/dal"
	api_define "blog_api/api/define"
	handler_article "blog_api/api/handler/article"
	handler_captcha "blog_api/api/handler/captcha"
	handler_login "blog_api/api/handler/login"
	handler_session "blog_api/api/handler/session"
	handler_user "blog_api/api/handler/user"
	middleware_check_login "blog_api/api/middleware/check_login"
	middleware_cors "blog_api/api/middleware/cors"
	middleware_session "blog_api/api/middleware/session"
	middleware_write_response "blog_api/api/middleware/write_response"
	pb "blog_api/api/protobuf"
	service_captcha "blog_api/api/service/captcha"
	_ "blog_api/api/validator"
	"blog_api/config"
	model_article "blog_api/db/model/article"
	model_user "blog_api/db/model/user"
	"blog_api/define"
	"blog_api/log"
	"blog_api/util/gostrgen"
	"net/http"
	"time"

	middleware_cache "blog_api/api/middleware/cache.v2"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	// init config
	config.Init()

	// init postgres
	dal.InitPostgres(config.GetPgDBConf())

	// init module
	model_user.Init()
	model_article.Init()
	service_captcha.Init()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		// 每个请求有个logId
		logId, err := gostrgen.RandGen(32, gostrgen.LowerUpperDigit, "", "")
		if err != nil {
			log.DefaultLogError("0", "create log id failed", err)
			logId = "0"
		}
		c.Set(api_define.CtxLogIdKey, logId)
	})
	r.Use(middleware_cors.Cors(config.GetCorsConfig()))

	// use session middleware
	r.Use(middleware_session.Session(define.SESSION_NAME, config.GetSessionRedisConfig()))

	// add monitor
	pprof.Register(r, "api/dev/pprof")

	// test session
	r.GET("/api/session/setname", handler_session.TestSetName())
	r.GET("/api/session/getname", handler_session.TestGetName())
	r.GET("/api/user/showme", handler_user.ShowMe())

	// if true {
	// 	model_article.Migrate()
	// 	model_user.Migrate()
	// }

	g := r.Group("/api")
	// g.POST(
	// 	"/register",
	// 	middleware_write_response.WriteResponseMw(),
	// 	handler_register.Register(),
	// )
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

	g.GET("/articles",
		middleware_write_response.WriteResponseMw(),
		middleware_cache.CacheMw(10*time.Second, "key", &pb.ArticleListResponse{}),
		handler_article.GetList(),
	)

	g.GET(
		"/articles/:articleId",
		middleware_write_response.WriteResponseMw(),
		middleware_cache.CacheMw(10*time.Second, "key", &pb.ArticleResponse{}),
		handler_article.GetById(),
	)

	// g.GET(
	// 	"/articles",
	// 	middleware_write_response.WriteResponseMw(),
	// 	middleware_cache.CacheMw(10*time.Second, "key", &pb.ArticleResponse{}),
	// 	handler_article.GetArticlesByU
	// )
	g.POST(
		"/articles",
		middleware_write_response.WriteResponseMw(),
		middleware_check_login.CheckLogin(),
		handler_article.Create(),
	)

	g.GET(
		"/captcha",
		middleware_write_response.WriteResponseMw(),
		handler_captcha.GetCaptcha(),
	)
	g.GET(
		"/captcha/verify",
		middleware_write_response.WriteResponseMw(),
		handler_captcha.VerifyCaptcha(),
	)
	// g.POST(
	// 	"/captcha",
	// )

	log.DefaultLog().Debugf("Starting")
	go accessDatabase()
	r.Run(config.RunAddr())

}

func accessDatabase() {
	tm := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-tm.C:
			func() {
				resp, err := http.Get("http://wncbb.cn/api/articles?offset=0&limit=1")
				log.DefaultLog().Info("%v %v", resp, err)
			}()
		}
	}
}
