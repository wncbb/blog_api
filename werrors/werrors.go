package werrors

import "fmt"

type Error struct {
	Code int64
}

var code2msgMap map[int64]string = map[int64]string{
	ERR_DB_CONFIG_NIL_CODE:         "DB config is nil",
	ERR_DB_CONN_FAIL_CODE:          "DB connect fail",
	ERR_DB_MAP_KEY_NOT_FIND_CODE:   "DB map key not found",
	ERR_FUNC_PARAMS_POINT_NIL_CODE: "function parameters point is nil",
}

func Code2Msg(code int64) string {
	msg, ok := code2msgMap[code]
	fmt.Println(msg, ok)
	if !ok {
		return "unknown error"
	}
	return msg
}

func (e Error) Error() string {
	return fmt.Sprintf("[err] code:%d, msg:'%s'", e.Code, Code2Msg(e.Code))
}

func New(code int64) Error {
	return Error{
		Code: code,
	}
}
