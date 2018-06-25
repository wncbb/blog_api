package article

import (
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
		log.DefaultLogDebug(c.GetString(api_define.CtxLogIdKey), "")
		if err != nil {
			log.DefaultLogError(c.GetString(api_define.CtxLogIdKey), "", err)
			pbResp.Code = pb.ResponseCode_QueryArgumentsError
			pbResp.Msg = err.Error()
			c.Set(define.CtxRespKey, pbResp)
			return
		}
		article := &article.Article{
			Title:   postReq.Title,
			Content: postReq.Content,
			UserId:  0,
		}
		modelArticle := model_article.NewModel(c)
		err = modelArticle.Create(article)
		if err != nil {
			log.DefaultLogError(c.GetString(api_define.CtxLogIdKey), "", err)
			pbResp.Code = pb.ResponseCode_InternalError
			pbResp.Msg = ""
			c.Set(define.CtxRespKey, pbResp)
			return
		}
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
			log.DefaultLogError(c.GetString(api_define.CtxLogIdKey), "", err)
			resp.Code = pb.ResponseCode_QueryArgumentsError
			c.Set(api_define.CtxRespKey, resp)
			return
		}
		modelArticle := model_article.NewModel(c)
		list, err := modelArticle.GetList(qry.Offset, qry.Limit)
		if err != nil {
			log.DefaultLogError(c.GetString(api_define.CtxLogIdKey), "", err)
			resp.Code = pb.ResponseCode_InternalError
			c.Set(api_define.CtxRespKey, resp)
			return
		}

		num, err := modelArticle.GetNum()
		if err != nil {
			log.DefaultLogError(c.GetString(api_define.CtxLogIdKey), "", err)
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
			log.DefaultLogError(c.GetString(api_define.CtxLogIdKey), "", err)
			pbResp.Code = pb.ResponseCode_QueryArgumentsError
			pbResp.Msg = err.Error()
			c.Set(define.CtxRespKey, pbResp)
		}
		modelArticle := model_article.NewModel(c)
		article, err := modelArticle.GetById(articleId)
		if err != nil {
			log.DefaultLogError(c.GetString(api_define.CtxLogIdKey), "", err)
			pbResp.Code = pb.ResponseCode_InternalError
			pbResp.Msg = ""
			c.Set(define.CtxRespKey, pbResp)
			return
		}

		pbResp.Data = &pb.ArticleData{
			Id:      strconv.FormatInt(article.Id, 10),
			Title:   article.Title,
			Content: article.Content,
		}
		c.Set(define.CtxRespKey, pbResp)
	}
}
