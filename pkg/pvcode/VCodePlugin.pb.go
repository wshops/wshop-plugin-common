// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.9
// source: VCodePlugin.proto

package pvcode

import (
	context "context"
	emptypb "github.com/knqyf263/go-plugin/types/known/emptypb"
	wpc "github.com/wshops/wshop-plugin-common/pkg/wpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendType int32

const (
	SendType_EMAIL SendType = 0
	SendType_SMS   SendType = 1
	SendType_OTHER SendType = 5
)

// Enum value maps for SendType.
var (
	SendType_name = map[int32]string{
		0: "EMAIL",
		1: "SMS",
		5: "OTHER",
	}
	SendType_value = map[string]int32{
		"EMAIL": 0,
		"SMS":   1,
		"OTHER": 5,
	}
)

func (x SendType) Enum() *SendType {
	p := new(SendType)
	*p = x
	return p
}

type SendVerificationCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendType SendType `protobuf:"varint,1,opt,name=sendType,proto3,enum=pvcode.SendType" json:"sendType,omitempty"`
	Target   string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Code     string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SendVerificationCodeRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SendVerificationCodeRequest) GetSendType() SendType {
	if x != nil {
		return x.SendType
	}
	return SendType_EMAIL
}

func (x *SendVerificationCodeRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *SendVerificationCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type SendVerificationCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SendVerificationCodeResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SendVerificationCodeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendVerificationCodeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// go:plugin type=plugin version=1
type VCode interface {
	GetPluginInfo(context.Context, emptypb.Empty) (wpc.PluginInfo, error)
	SendVerificationCode(context.Context, SendVerificationCodeRequest) (SendVerificationCodeResponse, error)
}

// go:plugin type=host
type HostFunctions interface {
	LogInfo(context.Context, wpc.HFuncLogRequest) (emptypb.Empty, error)
	LogError(context.Context, wpc.HFuncLogRequest) (emptypb.Empty, error)
	LogDebug(context.Context, wpc.HFuncLogRequest) (emptypb.Empty, error)
	LogWarn(context.Context, wpc.HFuncLogRequest) (emptypb.Empty, error)
	LogFatal(context.Context, wpc.HFuncLogRequest) (emptypb.Empty, error)
	LogPanic(context.Context, wpc.HFuncLogRequest) (emptypb.Empty, error)
	HttpRequest(context.Context, wpc.HFuncHttpRequest) (wpc.HFuncHttpResponse, error)
}
