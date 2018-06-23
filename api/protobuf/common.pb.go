// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	common.proto

It has these top-level messages:
	BaseResponse
	LoginData
	LoginResponse
	ArticleData
	ArticleResponse
	ArticleListData
	ArticleListResponse
	CreateArticleResponse
	RegisterData
	RegisterResponse
	GetCaptchaData
	GetCaptchaResponse
	VerifyCaptchaData
	VerifyCaptchaResponse
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResponseCode int32

const (
	ResponseCode_Success             ResponseCode = 0
	ResponseCode_Empty               ResponseCode = 1
	ResponseCode_QueryArgumentsError ResponseCode = 10001
	ResponseCode_ShouldLogoutFirst   ResponseCode = 10002
	ResponseCode_ShouldLoginFirst    ResponseCode = 10003
	ResponseCode_InternalError       ResponseCode = 40001
	ResponseCode_GenRandomError      ResponseCode = 40002
	ResponseCode_VerifyCaptchaError  ResponseCode = 40003
)

var ResponseCode_name = map[int32]string{
	0:     "Success",
	1:     "Empty",
	10001: "QueryArgumentsError",
	10002: "ShouldLogoutFirst",
	10003: "ShouldLoginFirst",
	40001: "InternalError",
	40002: "GenRandomError",
	40003: "VerifyCaptchaError",
}
var ResponseCode_value = map[string]int32{
	"Success":             0,
	"Empty":               1,
	"QueryArgumentsError": 10001,
	"ShouldLogoutFirst":   10002,
	"ShouldLoginFirst":    10003,
	"InternalError":       40001,
	"GenRandomError":      40002,
	"VerifyCaptchaError":  40003,
}

func (x ResponseCode) String() string {
	return proto.EnumName(ResponseCode_name, int32(x))
}
func (ResponseCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BaseResponse struct {
	Code ResponseCode `protobuf:"varint,1,opt,name=code,enum=ResponseCode" json:"code,omitempty"`
	Msg  string       `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *BaseResponse) Reset()                    { *m = BaseResponse{} }
func (m *BaseResponse) String() string            { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()               {}
func (*BaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BaseResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_Success
}

func (m *BaseResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LoginData struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *LoginData) Reset()                    { *m = LoginData{} }
func (m *LoginData) String() string            { return proto.CompactTextString(m) }
func (*LoginData) ProtoMessage()               {}
func (*LoginData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginData) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginData) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LoginResponse struct {
	Code ResponseCode `protobuf:"varint,1,opt,name=code,enum=ResponseCode" json:"code,omitempty"`
	Msg  string       `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data *LoginData   `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_Success
}

func (m *LoginResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LoginResponse) GetData() *LoginData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ArticleData struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *ArticleData) Reset()                    { *m = ArticleData{} }
func (m *ArticleData) String() string            { return proto.CompactTextString(m) }
func (*ArticleData) ProtoMessage()               {}
func (*ArticleData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ArticleData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ArticleData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticleData) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ArticleResponse struct {
	Code ResponseCode `protobuf:"varint,1,opt,name=code,enum=ResponseCode" json:"code,omitempty"`
	Msg  string       `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data *ArticleData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *ArticleResponse) Reset()                    { *m = ArticleResponse{} }
func (m *ArticleResponse) String() string            { return proto.CompactTextString(m) }
func (*ArticleResponse) ProtoMessage()               {}
func (*ArticleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ArticleResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_Success
}

func (m *ArticleResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ArticleResponse) GetData() *ArticleData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ArticleListData struct {
	List []*ArticleData `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	Num  string         `protobuf:"bytes,2,opt,name=num" json:"num,omitempty"`
}

func (m *ArticleListData) Reset()                    { *m = ArticleListData{} }
func (m *ArticleListData) String() string            { return proto.CompactTextString(m) }
func (*ArticleListData) ProtoMessage()               {}
func (*ArticleListData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ArticleListData) GetList() []*ArticleData {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ArticleListData) GetNum() string {
	if m != nil {
		return m.Num
	}
	return ""
}

type ArticleListResponse struct {
	Code ResponseCode     `protobuf:"varint,1,opt,name=code,enum=ResponseCode" json:"code,omitempty"`
	Msg  string           `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data *ArticleListData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *ArticleListResponse) Reset()                    { *m = ArticleListResponse{} }
func (m *ArticleListResponse) String() string            { return proto.CompactTextString(m) }
func (*ArticleListResponse) ProtoMessage()               {}
func (*ArticleListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ArticleListResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_Success
}

func (m *ArticleListResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ArticleListResponse) GetData() *ArticleListData {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateArticleResponse struct {
	Code ResponseCode `protobuf:"varint,1,opt,name=code,enum=ResponseCode" json:"code,omitempty"`
	Msg  string       `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data *ArticleData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *CreateArticleResponse) Reset()                    { *m = CreateArticleResponse{} }
func (m *CreateArticleResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateArticleResponse) ProtoMessage()               {}
func (*CreateArticleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateArticleResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_Success
}

func (m *CreateArticleResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CreateArticleResponse) GetData() *ArticleData {
	if m != nil {
		return m.Data
	}
	return nil
}

type RegisterData struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ShowMsg string `protobuf:"bytes,2,opt,name=showMsg" json:"showMsg,omitempty"`
}

func (m *RegisterData) Reset()                    { *m = RegisterData{} }
func (m *RegisterData) String() string            { return proto.CompactTextString(m) }
func (*RegisterData) ProtoMessage()               {}
func (*RegisterData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RegisterData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterData) GetShowMsg() string {
	if m != nil {
		return m.ShowMsg
	}
	return ""
}

type RegisterResponse struct {
	Code ResponseCode  `protobuf:"varint,1,opt,name=code,enum=ResponseCode" json:"code,omitempty"`
	Msg  string        `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data *RegisterData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RegisterResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_Success
}

func (m *RegisterResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RegisterResponse) GetData() *RegisterData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetCaptchaData struct {
	CaptchaData string `protobuf:"bytes,1,opt,name=captchaData" json:"captchaData,omitempty"`
	CaptchaId   string `protobuf:"bytes,2,opt,name=captchaId" json:"captchaId,omitempty"`
}

func (m *GetCaptchaData) Reset()                    { *m = GetCaptchaData{} }
func (m *GetCaptchaData) String() string            { return proto.CompactTextString(m) }
func (*GetCaptchaData) ProtoMessage()               {}
func (*GetCaptchaData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetCaptchaData) GetCaptchaData() string {
	if m != nil {
		return m.CaptchaData
	}
	return ""
}

func (m *GetCaptchaData) GetCaptchaId() string {
	if m != nil {
		return m.CaptchaId
	}
	return ""
}

type GetCaptchaResponse struct {
	Code ResponseCode    `protobuf:"varint,1,opt,name=code,enum=ResponseCode" json:"code,omitempty"`
	Msg  string          `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data *GetCaptchaData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *GetCaptchaResponse) Reset()                    { *m = GetCaptchaResponse{} }
func (m *GetCaptchaResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCaptchaResponse) ProtoMessage()               {}
func (*GetCaptchaResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetCaptchaResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_Success
}

func (m *GetCaptchaResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetCaptchaResponse) GetData() *GetCaptchaData {
	if m != nil {
		return m.Data
	}
	return nil
}

type VerifyCaptchaData struct {
	ShowMsg string `protobuf:"bytes,1,opt,name=showMsg" json:"showMsg,omitempty"`
}

func (m *VerifyCaptchaData) Reset()                    { *m = VerifyCaptchaData{} }
func (m *VerifyCaptchaData) String() string            { return proto.CompactTextString(m) }
func (*VerifyCaptchaData) ProtoMessage()               {}
func (*VerifyCaptchaData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *VerifyCaptchaData) GetShowMsg() string {
	if m != nil {
		return m.ShowMsg
	}
	return ""
}

type VerifyCaptchaResponse struct {
	Code ResponseCode       `protobuf:"varint,1,opt,name=code,enum=ResponseCode" json:"code,omitempty"`
	Msg  string             `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data *VerifyCaptchaData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *VerifyCaptchaResponse) Reset()                    { *m = VerifyCaptchaResponse{} }
func (m *VerifyCaptchaResponse) String() string            { return proto.CompactTextString(m) }
func (*VerifyCaptchaResponse) ProtoMessage()               {}
func (*VerifyCaptchaResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *VerifyCaptchaResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_Success
}

func (m *VerifyCaptchaResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *VerifyCaptchaResponse) GetData() *VerifyCaptchaData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseResponse)(nil), "BaseResponse")
	proto.RegisterType((*LoginData)(nil), "LoginData")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*ArticleData)(nil), "ArticleData")
	proto.RegisterType((*ArticleResponse)(nil), "ArticleResponse")
	proto.RegisterType((*ArticleListData)(nil), "ArticleListData")
	proto.RegisterType((*ArticleListResponse)(nil), "ArticleListResponse")
	proto.RegisterType((*CreateArticleResponse)(nil), "CreateArticleResponse")
	proto.RegisterType((*RegisterData)(nil), "RegisterData")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
	proto.RegisterType((*GetCaptchaData)(nil), "GetCaptchaData")
	proto.RegisterType((*GetCaptchaResponse)(nil), "GetCaptchaResponse")
	proto.RegisterType((*VerifyCaptchaData)(nil), "VerifyCaptchaData")
	proto.RegisterType((*VerifyCaptchaResponse)(nil), "VerifyCaptchaResponse")
	proto.RegisterEnum("ResponseCode", ResponseCode_name, ResponseCode_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0x93, 0x36, 0x78, 0xf2, 0xa7, 0xee, 0xa6, 0xa9, 0x7c, 0x40, 0xc8, 0x35, 0x08,
	0x45, 0x48, 0xf8, 0x50, 0x84, 0xc4, 0xb5, 0x0d, 0x01, 0x45, 0x4a, 0x23, 0x70, 0x25, 0x0e, 0xdc,
	0x5c, 0x7b, 0x9b, 0x2c, 0x8a, 0x77, 0xa3, 0xdd, 0xb1, 0x50, 0x1e, 0x03, 0x78, 0x15, 0x0e, 0xb9,
	0x02, 0x2f, 0x86, 0x6c, 0xaf, 0x93, 0x98, 0xf6, 0x66, 0xa9, 0xb7, 0xfd, 0xbe, 0xc9, 0x7c, 0xf3,
	0xdb, 0x9d, 0xc8, 0xd0, 0x89, 0x44, 0x92, 0x08, 0xee, 0xaf, 0xa4, 0x40, 0xe1, 0x8d, 0xa0, 0x73,
	0x19, 0x2a, 0x1a, 0x50, 0xb5, 0x12, 0x5c, 0x51, 0x72, 0x06, 0xcd, 0x48, 0xc4, 0xd4, 0x31, 0x5c,
	0x63, 0xd8, 0x3b, 0xef, 0xfa, 0x65, 0x61, 0x24, 0x62, 0x1a, 0xe4, 0x25, 0x62, 0x43, 0x23, 0x51,
	0x73, 0xc7, 0x74, 0x8d, 0xa1, 0x15, 0x64, 0x47, 0xef, 0x0d, 0x58, 0x53, 0x31, 0x67, 0xfc, 0x5d,
	0x88, 0x21, 0x39, 0x85, 0xc3, 0x54, 0x51, 0x39, 0x89, 0xf3, 0x0c, 0x2b, 0xd0, 0xea, 0x9e, 0xb6,
	0x18, 0xba, 0x79, 0x5b, 0xad, 0xe1, 0xe4, 0x29, 0x34, 0xe3, 0x10, 0x43, 0xa7, 0xe1, 0x1a, 0xc3,
	0xf6, 0x39, 0xf8, 0x5b, 0x92, 0x20, 0xf7, 0xbd, 0x2b, 0x68, 0x5f, 0x48, 0x64, 0xd1, 0x92, 0xe6,
	0x78, 0x3d, 0x30, 0x59, 0x89, 0x66, 0xb2, 0x98, 0x9c, 0xc0, 0x01, 0x32, 0x5c, 0x52, 0x1d, 0x59,
	0x08, 0xe2, 0x40, 0x2b, 0x12, 0x1c, 0x29, 0xc7, 0x3c, 0xd7, 0x0a, 0x4a, 0xe9, 0x2d, 0xe0, 0x48,
	0xc7, 0xd5, 0xc3, 0x76, 0x2b, 0xd8, 0x1d, 0x7f, 0x8f, 0x51, 0x83, 0x8f, 0xb7, 0x93, 0xa6, 0x4c,
	0x61, 0x0e, 0xef, 0x42, 0x73, 0xc9, 0x14, 0x3a, 0x86, 0xdb, 0xb8, 0xdb, 0x94, 0x55, 0xb2, 0x41,
	0x3c, 0x4d, 0xca, 0x41, 0x3c, 0x4d, 0xbc, 0x15, 0xf4, 0xf7, 0x62, 0xea, 0x41, 0x3f, 0xaf, 0x40,
	0xdb, 0xfe, 0x7f, 0x7c, 0x1a, 0x9c, 0xc3, 0x60, 0x24, 0x69, 0x88, 0xf4, 0x81, 0x1e, 0xea, 0x2d,
	0x74, 0x02, 0x3a, 0x67, 0x0a, 0xa9, 0xbc, 0x77, 0xc5, 0x0e, 0xb4, 0xd4, 0x42, 0x7c, 0xbb, 0xda,
	0xe6, 0x96, 0xd2, 0xfb, 0x0a, 0x76, 0xd9, 0x59, 0x0f, 0xf2, 0xac, 0x02, 0x99, 0x35, 0xed, 0x78,
	0x34, 0xe5, 0x47, 0xe8, 0x7d, 0xa0, 0x38, 0x0a, 0x57, 0x18, 0x2d, 0x42, 0xbd, 0xcd, 0x76, 0xb4,
	0x93, 0x1a, 0x78, 0xdf, 0x22, 0x4f, 0xc0, 0xd2, 0x72, 0x12, 0xeb, 0x71, 0x3b, 0xc3, 0xe3, 0x40,
	0x76, 0x89, 0xf5, 0xf8, 0x9f, 0x55, 0xf8, 0x8f, 0xfc, 0x2a, 0xa9, 0xbe, 0xc1, 0x2b, 0x38, 0xfe,
	0x4c, 0x25, 0xbb, 0x5d, 0xef, 0x5f, 0x62, 0xef, 0x71, 0x8d, 0xea, 0xe3, 0x22, 0x0c, 0x2a, 0x3f,
	0xaf, 0x47, 0xf8, 0xa2, 0x42, 0x48, 0xfc, 0x3b, 0x24, 0x05, 0xe4, 0xcb, 0x5f, 0x46, 0xf6, 0x6f,
	0xd8, 0x05, 0x92, 0x36, 0xb4, 0xae, 0xd3, 0x28, 0xa2, 0x4a, 0xd9, 0x8f, 0x88, 0x05, 0x07, 0xe3,
	0x64, 0x85, 0x6b, 0xdb, 0x20, 0x0e, 0xf4, 0x3f, 0xa5, 0x54, 0xae, 0x2f, 0xe4, 0x3c, 0x4d, 0x28,
	0x47, 0x35, 0x96, 0x52, 0x48, 0xfb, 0xfb, 0x8c, 0x9c, 0xc2, 0xf1, 0xf5, 0x42, 0xa4, 0xcb, 0x78,
	0x2a, 0xe6, 0x22, 0xc5, 0xf7, 0x4c, 0x2a, 0xb4, 0x7f, 0xcc, 0xc8, 0x00, 0xec, 0xad, 0xcf, 0x78,
	0x61, 0xff, 0x9c, 0x91, 0x3e, 0x74, 0x27, 0x1c, 0xa9, 0xe4, 0xe1, 0xb2, 0x88, 0xf8, 0xbd, 0x31,
	0xc9, 0x49, 0xb6, 0x6d, 0x1e, 0x84, 0x3c, 0x16, 0x49, 0xe1, 0xfe, 0xd9, 0x98, 0xc4, 0x01, 0x52,
	0xe1, 0x2e, 0x2a, 0x7f, 0x37, 0xe6, 0x25, 0x7c, 0x79, 0x9c, 0x7f, 0x90, 0x6f, 0xd2, 0xdb, 0x9b,
	0xc3, 0xfc, 0xf4, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xc6, 0xe6, 0xc6, 0xaa, 0x05,
	0x00, 0x00,
}
