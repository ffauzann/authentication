// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: authentication.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginMethod int32

const (
	LoginMethod_METHOD_UNSPECIFIED LoginMethod = 0
	LoginMethod_METHOD_LOGIN       LoginMethod = 1
	LoginMethod_METHOD_RECOVERY    LoginMethod = 2
)

// Enum value maps for LoginMethod.
var (
	LoginMethod_name = map[int32]string{
		0: "METHOD_UNSPECIFIED",
		1: "METHOD_LOGIN",
		2: "METHOD_RECOVERY",
	}
	LoginMethod_value = map[string]int32{
		"METHOD_UNSPECIFIED": 0,
		"METHOD_LOGIN":       1,
		"METHOD_RECOVERY":    2,
	}
)

func (x LoginMethod) Enum() *LoginMethod {
	p := new(LoginMethod)
	*p = x
	return p
}

func (x LoginMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_authentication_proto_enumTypes[0].Descriptor()
}

func (LoginMethod) Type() protoreflect.EnumType {
	return &file_authentication_proto_enumTypes[0]
}

func (x LoginMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginMethod.Descriptor instead.
func (LoginMethod) EnumDescriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{0}
}

type RegisterStatusCode int32

const (
	RegisterStatusCode_RSC_UNSPECIFIED RegisterStatusCode = 0
	RegisterStatusCode_RSC_SUCCEED     RegisterStatusCode = 1
	RegisterStatusCode_RSC_FAILED      RegisterStatusCode = 2
)

// Enum value maps for RegisterStatusCode.
var (
	RegisterStatusCode_name = map[int32]string{
		0: "RSC_UNSPECIFIED",
		1: "RSC_SUCCEED",
		2: "RSC_FAILED",
	}
	RegisterStatusCode_value = map[string]int32{
		"RSC_UNSPECIFIED": 0,
		"RSC_SUCCEED":     1,
		"RSC_FAILED":      2,
	}
)

func (x RegisterStatusCode) Enum() *RegisterStatusCode {
	p := new(RegisterStatusCode)
	*p = x
	return p
}

func (x RegisterStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_authentication_proto_enumTypes[1].Descriptor()
}

func (RegisterStatusCode) Type() protoreflect.EnumType {
	return &file_authentication_proto_enumTypes[1]
}

func (x RegisterStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterStatusCode.Descriptor instead.
func (RegisterStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{1}
}

type IsUserExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       *string `protobuf:"bytes,1,opt,name=email,proto3,oneof" json:"email,omitempty"`
	PhoneNumber *string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	Username    *string `protobuf:"bytes,3,opt,name=username,proto3,oneof" json:"username,omitempty"`
}

func (x *IsUserExistRequest) Reset() {
	*x = IsUserExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserExistRequest) ProtoMessage() {}

func (x *IsUserExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserExistRequest.ProtoReflect.Descriptor instead.
func (*IsUserExistRequest) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *IsUserExistRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *IsUserExistRequest) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *IsUserExistRequest) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

type IsUserExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExist bool     `protobuf:"varint,1,opt,name=is_exist,json=isExist,proto3" json:"is_exist,omitempty"`
	Reasons []string `protobuf:"bytes,2,rep,name=reasons,proto3" json:"reasons,omitempty"`
}

func (x *IsUserExistResponse) Reset() {
	*x = IsUserExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserExistResponse) ProtoMessage() {}

func (x *IsUserExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserExistResponse.ProtoReflect.Descriptor instead.
func (*IsUserExistResponse) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *IsUserExistResponse) GetIsExist() bool {
	if x != nil {
		return x.IsExist
	}
	return false
}

func (x *IsUserExistResponse) GetReasons() []string {
	if x != nil {
		return x.Reasons
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id can be filled by email, phone number, or username
	UserId     string                   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password   string                   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Method     LoginMethod              `protobuf:"varint,3,opt,name=method,proto3,enum=authentication.LoginMethod" json:"method,omitempty"`
	DeviceInfo *LoginRequest_DeviceInfo `protobuf:"bytes,4,opt,name=device_info,json=deviceInfo,proto3" json:"device_info,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetMethod() LoginMethod {
	if x != nil {
		return x.Method
	}
	return LoginMethod_METHOD_UNSPECIFIED
}

func (x *LoginRequest) GetDeviceInfo() *LoginRequest_DeviceInfo {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Username    *string `protobuf:"bytes,2,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Email       string  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber *string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	Password    string  `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    RegisterStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=authentication.RegisterStatusCode" json:"code,omitempty"`
	Reasons []string           `protobuf:"bytes,2,rep,name=reasons,proto3" json:"reasons,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterResponse) GetCode() RegisterStatusCode {
	if x != nil {
		return x.Code
	}
	return RegisterStatusCode_RSC_UNSPECIFIED
}

func (x *RegisterResponse) GetReasons() []string {
	if x != nil {
		return x.Reasons
	}
	return nil
}

type LoginRequest_DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   *string                         `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Model  *string                         `protobuf:"bytes,3,opt,name=model,proto3,oneof" json:"model,omitempty"`
	OsInfo *LoginRequest_DeviceInfo_OSInfo `protobuf:"bytes,4,opt,name=os_info,json=osInfo,proto3,oneof" json:"os_info,omitempty"`
}

func (x *LoginRequest_DeviceInfo) Reset() {
	*x = LoginRequest_DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest_DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest_DeviceInfo) ProtoMessage() {}

func (x *LoginRequest_DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest_DeviceInfo.ProtoReflect.Descriptor instead.
func (*LoginRequest_DeviceInfo) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{2, 0}
}

func (x *LoginRequest_DeviceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginRequest_DeviceInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *LoginRequest_DeviceInfo) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *LoginRequest_DeviceInfo) GetOsInfo() *LoginRequest_DeviceInfo_OSInfo {
	if x != nil {
		return x.OsInfo
	}
	return nil
}

type LoginRequest_DeviceInfo_OSInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *LoginRequest_DeviceInfo_OSInfo) Reset() {
	*x = LoginRequest_DeviceInfo_OSInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest_DeviceInfo_OSInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest_DeviceInfo_OSInfo) ProtoMessage() {}

func (x *LoginRequest_DeviceInfo_OSInfo) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest_DeviceInfo_OSInfo.ProtoReflect.Descriptor instead.
func (*LoginRequest_DeviceInfo_OSInfo) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *LoginRequest_DeviceInfo_OSInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginRequest_DeviceInfo_OSInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_authentication_proto protoreflect.FileDescriptor

var file_authentication_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x49, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x13, 0x49, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0xba, 0x03, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x48, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0xf5, 0x01, 0x0a, 0x0a, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x4c, 0x0a,
	0x07, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x02,
	0x52, 0x06, 0x6f, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x1a, 0x36, 0x0a, 0x06, 0x4f,
	0x53, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x73, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x57, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x10,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x73, 0x2a, 0x4c, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x02,
	0x2a, 0x4a, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x53, 0x43, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52,
	0x53, 0x43, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x52, 0x53, 0x43, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0x80, 0x02, 0x0a,
	0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0b,
	0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_authentication_proto_rawDescOnce sync.Once
	file_authentication_proto_rawDescData = file_authentication_proto_rawDesc
)

func file_authentication_proto_rawDescGZIP() []byte {
	file_authentication_proto_rawDescOnce.Do(func() {
		file_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_authentication_proto_rawDescData)
	})
	return file_authentication_proto_rawDescData
}

var file_authentication_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_authentication_proto_goTypes = []interface{}{
	(LoginMethod)(0),                       // 0: authentication.LoginMethod
	(RegisterStatusCode)(0),                // 1: authentication.RegisterStatusCode
	(*IsUserExistRequest)(nil),             // 2: authentication.IsUserExistRequest
	(*IsUserExistResponse)(nil),            // 3: authentication.IsUserExistResponse
	(*LoginRequest)(nil),                   // 4: authentication.LoginRequest
	(*LoginResponse)(nil),                  // 5: authentication.LoginResponse
	(*RegisterRequest)(nil),                // 6: authentication.RegisterRequest
	(*RegisterResponse)(nil),               // 7: authentication.RegisterResponse
	(*LoginRequest_DeviceInfo)(nil),        // 8: authentication.LoginRequest.DeviceInfo
	(*LoginRequest_DeviceInfo_OSInfo)(nil), // 9: authentication.LoginRequest.DeviceInfo.OSInfo
}
var file_authentication_proto_depIdxs = []int32{
	0, // 0: authentication.LoginRequest.method:type_name -> authentication.LoginMethod
	8, // 1: authentication.LoginRequest.device_info:type_name -> authentication.LoginRequest.DeviceInfo
	1, // 2: authentication.RegisterResponse.code:type_name -> authentication.RegisterStatusCode
	9, // 3: authentication.LoginRequest.DeviceInfo.os_info:type_name -> authentication.LoginRequest.DeviceInfo.OSInfo
	2, // 4: authentication.AuthService.IsUserExist:input_type -> authentication.IsUserExistRequest
	4, // 5: authentication.AuthService.Login:input_type -> authentication.LoginRequest
	6, // 6: authentication.AuthService.Register:input_type -> authentication.RegisterRequest
	3, // 7: authentication.AuthService.IsUserExist:output_type -> authentication.IsUserExistResponse
	5, // 8: authentication.AuthService.Login:output_type -> authentication.LoginResponse
	7, // 9: authentication.AuthService.Register:output_type -> authentication.RegisterResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_authentication_proto_init() }
func file_authentication_proto_init() {
	if File_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserExistRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserExistResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authentication_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authentication_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authentication_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest_DeviceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authentication_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest_DeviceInfo_OSInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_authentication_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_authentication_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_authentication_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authentication_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authentication_proto_goTypes,
		DependencyIndexes: file_authentication_proto_depIdxs,
		EnumInfos:         file_authentication_proto_enumTypes,
		MessageInfos:      file_authentication_proto_msgTypes,
	}.Build()
	File_authentication_proto = out.File
	file_authentication_proto_rawDesc = nil
	file_authentication_proto_goTypes = nil
	file_authentication_proto_depIdxs = nil
}
