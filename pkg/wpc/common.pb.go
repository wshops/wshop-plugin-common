// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.7
// source: common.proto

package wpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

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

	Id             string            `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name           string            `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Version        string            `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	Author         string            `protobuf:"bytes,4,opt,name=Author,proto3" json:"Author,omitempty"`
	Description    string            `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Icon           string            `protobuf:"bytes,6,opt,name=Icon,proto3" json:"Icon,omitempty"`
	PluginConstant map[string]string `protobuf:"bytes,7,rep,name=PluginConstant,proto3" json:"PluginConstant,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserAttributes []*PAttributes    `protobuf:"bytes,8,rep,name=UserAttributes,proto3" json:"UserAttributes,omitempty"`
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

func (x *PluginInfo) GetPluginConstant() map[string]string {
	if x != nil {
		return x.PluginConstant
	}
	return nil
}

func (x *PluginInfo) GetUserAttributes() []*PAttributes {
	if x != nil {
		return x.UserAttributes
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
