package write_response

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"blog_api/api/define"
)

type BaseResponse interface {
	GetCode() int64
	GetMsg() string
}

func WriteResponseMw() func(*gin.Context) {
	return func(c *gin.Context) {
		var err error
		generalQuery := define.NewDefaultGeneralQuery()
		err = c.ShouldBindQuery(generalQuery)

		if err != nil {
			//TODO
			c.String(http.StatusOK, "general query arguments error")
			return
		}
		c.Next()
		resp, exists := c.Get(define.CtxRespKey)
		if !exists {
			//TODO
			c.String(http.StatusOK, "no response")
			return
		}

		// 指针类型断言proto.Message
		pbMsg, ok := resp.(proto.Message)
		if !ok {
			//TODO
			c.String(http.StatusOK, "response pb error")
			return
		}

		// 可以拿到Code, Msg
		// baseResp, ok := resp.(BaseResponse)
		// if ok {
		// 	if baseResp.GetMsg() == "" {
		// 	}
		// }

		switch generalQuery.Format {
		case define.RespFormatTypeJSON:
			var jsonConverter = &jsonpb.Marshaler{
				EnumsAsInts: true,
				// 如果key:val val是默认值,是否显示
				EmitDefaults: true,
				OrigName:     true,
			}
			jsonBytes := new(bytes.Buffer)
			err := jsonConverter.Marshal(jsonBytes, pbMsg)
			if err != nil {
				//TODO
				c.String(http.StatusOK, "transform response to json error")
				return
			}

			c.Data(http.StatusOK, binding.MIMEJSON, jsonBytes.Bytes())
		case define.RespFormatTypePb:
			data, err := proto.Marshal(pbMsg)
			if err != nil {
				//TODO
				c.String(http.StatusOK, "transform response to pb error")
				return
			}
			c.Data(http.StatusOK, binding.MIMEPROTOBUF, data)
		}
	}
}
