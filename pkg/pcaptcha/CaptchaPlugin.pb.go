// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.7
// source: CaptchaPlugin.proto

package pcaptcha

import (
	context "context"
	emptypb "github.com/knqyf263/go-plugin/types/known/emptypb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VerifyCaptchaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Captcha string `protobuf:"bytes,1,opt,name=captcha,proto3" json:"captcha,omitempty"`
}

func (x *VerifyCaptchaRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *VerifyCaptchaRequest) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

type VerifyCaptchaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

func (x *VerifyCaptchaResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *VerifyCaptchaResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type GetCustomHtmlInputFieldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCustomHtmlInputFieldRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

type GetCustomHtmlInputFieldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *GetCustomHtmlInputFieldResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetCustomHtmlInputFieldResponse) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

type GetCustomHtmlHeadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCustomHtmlHeadRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

type GetCustomHtmlHeadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *GetCustomHtmlHeadResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetCustomHtmlHeadResponse) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

type GetCustomHtmlBodyEndRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCustomHtmlBodyEndRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

type GetCustomHtmlBodyEndResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *GetCustomHtmlBodyEndResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetCustomHtmlBodyEndResponse) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

// **************************
// *        commons         *
// **************************
type ConfigPluginInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigPluginInfoRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

type PluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string         `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name           string         `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Version        string         `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	Author         string         `protobuf:"bytes,4,opt,name=Author,proto3" json:"Author,omitempty"`
	Description    string         `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Icon           string         `protobuf:"bytes,6,opt,name=Icon,proto3" json:"Icon,omitempty"`
	UserAttributes []*PAttributes `protobuf:"bytes,7,rep,name=UserAttributes,proto3" json:"UserAttributes,omitempty"`
	PluginConstant []*PConstants  `protobuf:"bytes,8,rep,name=PluginConstant,proto3" json:"PluginConstant,omitempty"`
}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *PluginInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PluginInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginInfo) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *PluginInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PluginInfo) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *PluginInfo) GetUserAttributes() []*PAttributes {
	if x != nil {
		return x.UserAttributes
	}
	return nil
}

func (x *PluginInfo) GetPluginConstant() []*PConstants {
	if x != nil {
		return x.PluginConstant
	}
	return nil
}

type PAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	VRule       string `protobuf:"bytes,5,opt,name=VRule,proto3" json:"VRule,omitempty"`
	VMsg        string `protobuf:"bytes,6,opt,name=VMsg,proto3" json:"VMsg,omitempty"`
}

func (x *PAttributes) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *PAttributes) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PAttributes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PAttributes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PAttributes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PAttributes) GetVRule() string {
	if x != nil {
		return x.VRule
	}
	return ""
}

func (x *PAttributes) GetVMsg() string {
	if x != nil {
		return x.VMsg
	}
	return ""
}

type PConstants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *PConstants) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *PConstants) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PConstants) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// go:plugin type=plugin version=1
type Captcha interface {
	ConfigPluginInfo(context.Context, emptypb.Empty) (PluginInfo, error)
	// VerifyCaptcha 验证验证码
	VerifyCaptcha(context.Context, VerifyCaptchaRequest) (VerifyCaptchaResponse, error)
	// GetCustomHtmlInputField 获取自定义的html输入框
	GetCustomHtmlInputField(context.Context, GetCustomHtmlInputFieldRequest) (GetCustomHtmlInputFieldResponse, error)
	// GetCustomHtmlHead 自定义 html 头部中需要加的内容
	GetCustomHtmlHead(context.Context, GetCustomHtmlHeadRequest) (GetCustomHtmlHeadResponse, error)
	// GetCustomHtmlBodyEnd 自定义 html body 中最结尾需要加的内容
	GetCustomHtmlBodyEnd(context.Context, GetCustomHtmlBodyEndRequest) (GetCustomHtmlBodyEndResponse, error)
}
