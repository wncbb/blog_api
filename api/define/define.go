package define

const CtxRespKey = "resp"
const CtxNeedCacheKey = "needCache"
const CtxUserIdKey = "userId"

const SessionUserIdKey = "uid"

type RespFormatType int

const (
	RespFormatTypeUnset RespFormatType = 0
	RespFormatTypeJSON  RespFormatType = 1
	RespFormatTypePb    RespFormatType = 2
)

const QueryFormat = "format"

type GeneralQuery struct {
	// 必须在tag里加form
	Format RespFormatType `json:"format" form:"format"`
}

func NewDefaultGeneralQuery() *GeneralQuery {
	return &GeneralQuery{
		Format: RespFormatTypeJSON,
	}
}
