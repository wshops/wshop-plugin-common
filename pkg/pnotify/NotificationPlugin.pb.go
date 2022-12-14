// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.9
// source: NotificationPlugin.proto

package pnotify

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

type SendNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pctx    *wpc.PluginContext `protobuf:"bytes,1,opt,name=pctx,proto3" json:"pctx,omitempty"`
	Target  string             `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Title   string             `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content string             `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendNotificationRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SendNotificationRequest) GetPctx() *wpc.PluginContext {
	if x != nil {
		return x.Pctx
	}
	return nil
}

func (x *SendNotificationRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *SendNotificationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendNotificationRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SendNotificationResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SendNotificationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendNotificationResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// go:plugin type=plugin version=1
type Notification interface {
	GetPluginInfo(context.Context, emptypb.Empty) (wpc.PluginInfo, error)
	SendNotification(context.Context, SendNotificationRequest) (SendNotificationResponse, error)
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
