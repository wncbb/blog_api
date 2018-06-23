package article

import (
	"fmt"
	"strconv"

	"blog_api/api/define"
	api_define "blog_api/api/define"
	pb "blog_api/api/protobuf"
	"blog_api/db/model/article"
	model_article "blog_api/db/model/article"
	"blog_api/log"

	"github.com/gin-gonic/gin"
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

type GetListQuery struct {
	Offset int64 `form:"offset" json:"offset"`
	Limit  int64 `form:"limit" json:"limit"`
}

func GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := &pb.ArticleListResponse{}

		qry := &GetListQuery{}
		err := c.BindQuery(qry)
		if err != nil {
			resp.Code = pb.ResponseCode_QueryArgumentsError
			c.Set(api_define.CtxRespKey, resp)
		}

		list, err := model_article.GetList(qry.Offset, qry.Limit)
		if err != nil {
			resp.Code = pb.ResponseCode_InternalError
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		num, err := model_article.GetNum()
		if err != nil {
			resp.Code = pb.ResponseCode_InternalError
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		respData := &pb.ArticleListData{}
		respData.Num = strconv.FormatInt(num, 10)
		respData.List = toPbArticleList(list)

		resp.Code = pb.ResponseCode_Success
		resp.Data = respData
		c.Set(api_define.CtxRespKey, resp)
		return
	}
}

func toPbArticleList(list []*model_article.Article) []*pb.ArticleData {
	pbList := make([]*pb.ArticleData, 0, len(list))
	for _, v := range list {
		pbArticle := &pb.ArticleData{
			Title:   v.Title,
			Id:      strconv.FormatInt(v.Id, 10),
			Content: v.Content,
		}
		pbList = append(pbList, pbArticle)
	}
	return pbList
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
