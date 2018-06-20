package article

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"wncbb.cn/api/define"
	pb "wncbb.cn/api/protobuf"
	"wncbb.cn/db/model/article"
	model_article "wncbb.cn/db/model/article"
	"wncbb.cn/log"
)

type CreateArticleRequest struct {
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		pbResp := &pb.CreateArticleResponse{}
		postReq := &CreateArticleRequest{}
		err = c.ShouldBind(postReq)
		log.DefaultLog().Debugf("req:%v err:%v", postReq, err)
		if err != nil {
			pbResp.Code = -1
			pbResp.Msg = err.Error()
			c.Set(define.CtxRespKey, pbResp)
			return
		}
		article := &article.Article{
			Title:   postReq.Title,
			Content: postReq.Content,
			UserId:  0,
		}
		err = model_article.Create(article)
		if err != nil {
			pbResp.Code = -1
			pbResp.Msg = err.Error()
			c.Set(define.CtxRespKey, pbResp)
		}
		log.DefaultLog().Debugf("create rst: %v", article)
		pbResp.Data = &pb.ArticleData{
			Id:      strconv.FormatInt(article.Id, 10),
			Title:   article.Title,
			Content: article.Content,
		}

		pbResp.Code = 0
		pbResp.Msg = ""
		c.Set(define.CtxRespKey, pbResp)
		return

	}
}

func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		pbResp := &pb.ArticleResponse{}

		articleIdStr := c.Param("articleId")
		articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
		if err != nil {
			pbResp.Code = 100
			pbResp.Msg = err.Error()
			c.Set(define.CtxRespKey, pbResp)
		}

		article, err := model_article.GetById(articleId)
		if err != nil {
			pbResp.Code = 100
			pbResp.Msg = err.Error()
			c.Set(define.CtxRespKey, pbResp)
		}

		pbResp.Data = &pb.ArticleData{
			Id:      strconv.FormatInt(article.Id, 10),
			Title:   article.Title,
			Content: article.Content,
		}
		c.Set(define.CtxRespKey, pbResp)
		fmt.Printf("cookie:%v\n", c.Request.Cookies())
	}
}
