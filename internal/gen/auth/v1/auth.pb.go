// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: auth/v1/auth.proto

package authv1

import (
	_ "buf.build/gen/go/envoyproxy/protoc-gen-validate/protocolbuffers/go/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Create contains information needed to create a new authUser.
type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegisterAuthUser *RegisterAuthUser `protobuf:"bytes,1,opt,name=register_auth_user,json=registerAuthUser,proto3" json:"register_auth_user,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetRegisterAuthUser() *RegisterAuthUser {
	if x != nil {
		return x.RegisterAuthUser
	}
	return nil
}

// CreateResponse returns the newly created authUser for the Create method.
type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RegisterAuthUser *RegisterAuthUser `protobuf:"bytes,2,opt,name=register_auth_user,json=registerAuthUser,proto3" json:"register_auth_user,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateAccountResponse) GetRegisterAuthUser() *RegisterAuthUser {
	if x != nil {
		return x.RegisterAuthUser
	}
	return nil
}

// logs a user in
type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthUserSignIn *AuthUserSignIn `protobuf:"bytes,1,opt,name=auth_user_sign_in,json=authUserSignIn,proto3" json:"auth_user_sign_in,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SignInRequest) GetAuthUserSignIn() *AuthUserSignIn {
	if x != nil {
		return x.AuthUserSignIn
	}
	return nil
}

// returns the user that was signed in
type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	IdToken     string `protobuf:"bytes,2,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	Scope       string `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
	TokenType   string `protobuf:"bytes,4,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ExpiresIn   int32  `protobuf:"varint,5,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SignInResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignInResponse) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

func (x *SignInResponse) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *SignInResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *SignInResponse) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

// GetAccountRequest contains the information needed to get a user from auth0
type GetAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccountRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

// GetAccountResponse returns the user that was requested
type GetAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *GetAccountResponse) Reset() {
	*x = GetAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountResponse) ProtoMessage() {}

func (x *GetAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountResponse.ProtoReflect.Descriptor instead.
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccountResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

// SignOutRequest contains the information needed to sign a user out
type SignOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutRequest) Reset() {
	*x = SignOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutRequest) ProtoMessage() {}

func (x *SignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutRequest.ProtoReflect.Descriptor instead.
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{6}
}

// SignOutResponse returns the user null value
type SignOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SignOutResponse) Reset() {
	*x = SignOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutResponse) ProtoMessage() {}

func (x *SignOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutResponse.ProtoReflect.Descriptor instead.
func (*SignOutResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *SignOutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Data types for requests and responses
// values needed to sign in a user
type AuthUserSignIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password     string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	GrantType    string `protobuf:"bytes,3,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`
	Audience     string `protobuf:"bytes,4,opt,name=audience,proto3" json:"audience,omitempty"`
	ClientId     string `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,6,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *AuthUserSignIn) Reset() {
	*x = AuthUserSignIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserSignIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserSignIn) ProtoMessage() {}

func (x *AuthUserSignIn) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserSignIn.ProtoReflect.Descriptor instead.
func (*AuthUserSignIn) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *AuthUserSignIn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthUserSignIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthUserSignIn) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *AuthUserSignIn) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *AuthUserSignIn) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AuthUserSignIn) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// conatains the values used for submitting an entry to auth0
type RegisterAuthUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GivenName  string `protobuf:"bytes,1,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName string `protobuf:"bytes,2,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Nickname   string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password   string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Connection string `protobuf:"bytes,6,opt,name=connection,proto3" json:"connection,omitempty"`
	ClientId   string `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *RegisterAuthUser) Reset() {
	*x = RegisterAuthUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAuthUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAuthUser) ProtoMessage() {}

func (x *RegisterAuthUser) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAuthUser.ProtoReflect.Descriptor instead.
func (*RegisterAuthUser) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{9}
}

func (x *RegisterAuthUser) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

func (x *RegisterAuthUser) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

func (x *RegisterAuthUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterAuthUser) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RegisterAuthUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterAuthUser) GetConnection() string {
	if x != nil {
		return x.Connection
	}
	return ""
}

func (x *RegisterAuthUser) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

// user details
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GivenName     string `protobuf:"bytes,1,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName    string `protobuf:"bytes,2,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified string `protobuf:"bytes,4,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Name          string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Nickname      string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Picture       string `protobuf:"bytes,7,opt,name=picture,proto3" json:"picture,omitempty"`
	Sub           string `protobuf:"bytes,8,opt,name=sub,proto3" json:"sub,omitempty"`
	UpdatedAt     string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UserId        string `protobuf:"bytes,10,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{10}
}

func (x *UserInfo) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

func (x *UserInfo) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetEmailVerified() string {
	if x != nil {
		return x.EmailVerified
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserInfo) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *UserInfo) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *UserInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_auth_v1_auth_proto protoreflect.FileDescriptor

var file_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x12,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x10, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x70, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x5d, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4c, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x49, 0x6e, 0x22, 0x36, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x44, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xb7, 0x02, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x76,
	0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a,
	0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0xf9, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22,
	0x0e, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x4d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x5d,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a,
	0x22, 0x0b, 0x2f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x51, 0x0a,
	0x07, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74,
	0x42, 0x8f, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x64, 0x6f, 0x2d, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x07,
	0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x13, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_auth_proto_rawDescOnce sync.Once
	file_auth_v1_auth_proto_rawDescData = file_auth_v1_auth_proto_rawDesc
)

func file_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_auth_proto_rawDescData)
	})
	return file_auth_v1_auth_proto_rawDescData
}

var file_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_auth_v1_auth_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),  // 0: auth.v1.CreateAccountRequest
	(*CreateAccountResponse)(nil), // 1: auth.v1.CreateAccountResponse
	(*SignInRequest)(nil),         // 2: auth.v1.SignInRequest
	(*SignInResponse)(nil),        // 3: auth.v1.SignInResponse
	(*GetAccountRequest)(nil),     // 4: auth.v1.GetAccountRequest
	(*GetAccountResponse)(nil),    // 5: auth.v1.GetAccountResponse
	(*SignOutRequest)(nil),        // 6: auth.v1.SignOutRequest
	(*SignOutResponse)(nil),       // 7: auth.v1.SignOutResponse
	(*AuthUserSignIn)(nil),        // 8: auth.v1.AuthUserSignIn
	(*RegisterAuthUser)(nil),      // 9: auth.v1.RegisterAuthUser
	(*UserInfo)(nil),              // 10: auth.v1.UserInfo
}
var file_auth_v1_auth_proto_depIdxs = []int32{
	9,  // 0: auth.v1.CreateAccountRequest.register_auth_user:type_name -> auth.v1.RegisterAuthUser
	9,  // 1: auth.v1.CreateAccountResponse.register_auth_user:type_name -> auth.v1.RegisterAuthUser
	8,  // 2: auth.v1.SignInRequest.auth_user_sign_in:type_name -> auth.v1.AuthUserSignIn
	10, // 3: auth.v1.GetAccountResponse.user_info:type_name -> auth.v1.UserInfo
	0,  // 4: auth.v1.AuthService.CreateAccount:input_type -> auth.v1.CreateAccountRequest
	2,  // 5: auth.v1.AuthService.SignIn:input_type -> auth.v1.SignInRequest
	4,  // 6: auth.v1.AuthService.GetAccount:input_type -> auth.v1.GetAccountRequest
	6,  // 7: auth.v1.AuthService.SignOut:input_type -> auth.v1.SignOutRequest
	1,  // 8: auth.v1.AuthService.CreateAccount:output_type -> auth.v1.CreateAccountResponse
	3,  // 9: auth.v1.AuthService.SignIn:output_type -> auth.v1.SignInResponse
	5,  // 10: auth.v1.AuthService.GetAccount:output_type -> auth.v1.GetAccountResponse
	7,  // 11: auth.v1.AuthService.SignOut:output_type -> auth.v1.SignOutResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_auth_v1_auth_proto_init() }
func file_auth_v1_auth_proto_init() {
	if File_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_auth_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_auth_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_auth_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
		file_auth_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountRequest); i {
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
		file_auth_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountResponse); i {
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
		file_auth_v1_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutRequest); i {
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
		file_auth_v1_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutResponse); i {
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
		file_auth_v1_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserSignIn); i {
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
		file_auth_v1_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAuthUser); i {
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
		file_auth_v1_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_auth_v1_auth_proto = out.File
	file_auth_v1_auth_proto_rawDesc = nil
	file_auth_v1_auth_proto_goTypes = nil
	file_auth_v1_auth_proto_depIdxs = nil
}
