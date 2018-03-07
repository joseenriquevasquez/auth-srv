// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/auth-srv/proto/oauth2/oauth2.proto

/*
Package go_micro_srv_auth_oauth2 is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/auth-srv/proto/oauth2/oauth2.proto

It has these top-level messages:
	Token
	AuthorizeRequest
	AuthorizeResponse
	TokenRequest
	TokenResponse
	RevokeRequest
	RevokeResponse
	IntrospectRequest
	IntrospectResponse
*/
package go_micro_srv_auth_oauth2

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

type Token struct {
	AccessToken  string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	TokenType    string   `protobuf:"bytes,2,opt,name=token_type,json=tokenType" json:"token_type,omitempty"`
	RefreshToken string   `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	ExpiresAt    int64    `protobuf:"varint,4,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	Scopes       []string `protobuf:"bytes,5,rep,name=scopes" json:"scopes,omitempty"`
	// metadata associated with the token
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Token) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *Token) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type AuthorizeRequest struct {
	// code, token (not supported)
	ResponseType string   `protobuf:"bytes,1,opt,name=response_type,json=responseType" json:"response_type,omitempty"`
	ClientId     string   `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Scopes       []string `protobuf:"bytes,3,rep,name=scopes" json:"scopes,omitempty"`
	State        string   `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	RedirectUri  string   `protobuf:"bytes,5,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
}

func (m *AuthorizeRequest) Reset()                    { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()               {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthorizeRequest) GetResponseType() string {
	if m != nil {
		return m.ResponseType
	}
	return ""
}

func (m *AuthorizeRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AuthorizeRequest) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *AuthorizeRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *AuthorizeRequest) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

type AuthorizeResponse struct {
	Code  string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	State string `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	// implicit response
	Token *Token `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *AuthorizeResponse) Reset()                    { *m = AuthorizeResponse{} }
func (m *AuthorizeResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthorizeResponse) ProtoMessage()               {}
func (*AuthorizeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthorizeResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *AuthorizeResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *AuthorizeResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type TokenRequest struct {
	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret" json:"client_secret,omitempty"`
	Code         string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	// password (not supported), client_credentials, authorization_code, refresh_token
	GrantType    string `protobuf:"bytes,4,opt,name=grant_type,json=grantType" json:"grant_type,omitempty"`
	RedirectUri  string `protobuf:"bytes,5,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
	RefreshToken string `protobuf:"bytes,6,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	// scopes can be added for client_credentials request
	Scopes []string `protobuf:"bytes,7,rep,name=scopes" json:"scopes,omitempty"`
	// metadata to be stored with a token that's generated
	Metadata map[string]string `protobuf:"bytes,8,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TokenRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *TokenRequest) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *TokenRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *TokenRequest) GetGrantType() string {
	if m != nil {
		return m.GrantType
	}
	return ""
}

func (m *TokenRequest) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *TokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *TokenRequest) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *TokenRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type TokenResponse struct {
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *TokenResponse) Reset()                    { *m = TokenResponse{} }
func (m *TokenResponse) String() string            { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()               {}
func (*TokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type RevokeRequest struct {
	// revoke access token
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	// revoke via refresh token
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *RevokeRequest) Reset()                    { *m = RevokeRequest{} }
func (m *RevokeRequest) String() string            { return proto.CompactTextString(m) }
func (*RevokeRequest) ProtoMessage()               {}
func (*RevokeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RevokeRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RevokeRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RevokeResponse struct {
}

func (m *RevokeResponse) Reset()                    { *m = RevokeResponse{} }
func (m *RevokeResponse) String() string            { return proto.CompactTextString(m) }
func (*RevokeResponse) ProtoMessage()               {}
func (*RevokeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type IntrospectRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *IntrospectRequest) Reset()                    { *m = IntrospectRequest{} }
func (m *IntrospectRequest) String() string            { return proto.CompactTextString(m) }
func (*IntrospectRequest) ProtoMessage()               {}
func (*IntrospectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IntrospectRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type IntrospectResponse struct {
	Token  *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Active bool   `protobuf:"varint,2,opt,name=active" json:"active,omitempty"`
}

func (m *IntrospectResponse) Reset()                    { *m = IntrospectResponse{} }
func (m *IntrospectResponse) String() string            { return proto.CompactTextString(m) }
func (*IntrospectResponse) ProtoMessage()               {}
func (*IntrospectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IntrospectResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *IntrospectResponse) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*Token)(nil), "go.micro.srv.auth.oauth2.Token")
	proto.RegisterType((*AuthorizeRequest)(nil), "go.micro.srv.auth.oauth2.AuthorizeRequest")
	proto.RegisterType((*AuthorizeResponse)(nil), "go.micro.srv.auth.oauth2.AuthorizeResponse")
	proto.RegisterType((*TokenRequest)(nil), "go.micro.srv.auth.oauth2.TokenRequest")
	proto.RegisterType((*TokenResponse)(nil), "go.micro.srv.auth.oauth2.TokenResponse")
	proto.RegisterType((*RevokeRequest)(nil), "go.micro.srv.auth.oauth2.RevokeRequest")
	proto.RegisterType((*RevokeResponse)(nil), "go.micro.srv.auth.oauth2.RevokeResponse")
	proto.RegisterType((*IntrospectRequest)(nil), "go.micro.srv.auth.oauth2.IntrospectRequest")
	proto.RegisterType((*IntrospectResponse)(nil), "go.micro.srv.auth.oauth2.IntrospectResponse")
}

func init() {
	proto.RegisterFile("github.com/micro/auth-srv/proto/oauth2/oauth2.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0x6d, 0xe2, 0x69, 0x82, 0xda, 0x55, 0x85, 0xac, 0xa0, 0x8a, 0xe0, 0x4a, 0x10,
	0x01, 0x75, 0xa4, 0x14, 0x10, 0x82, 0x53, 0x0f, 0x20, 0xf5, 0x80, 0x40, 0xa6, 0x08, 0x09, 0x09,
	0x45, 0xee, 0x66, 0xda, 0x58, 0x6d, 0xbc, 0x66, 0x77, 0x1d, 0x11, 0x3e, 0x84, 0x3f, 0xe0, 0xc8,
	0xcf, 0xf1, 0x05, 0xc8, 0xbb, 0x6b, 0xc7, 0x69, 0x68, 0x13, 0x84, 0x38, 0xd5, 0xf3, 0x76, 0x76,
	0x66, 0xde, 0x7b, 0xd3, 0x0d, 0x1c, 0x9c, 0x25, 0x72, 0x9c, 0x9f, 0x84, 0x94, 0x4d, 0xfa, 0x93,
	0x84, 0x72, 0xd6, 0x8f, 0x73, 0x39, 0xde, 0x17, 0x7c, 0xda, 0xcf, 0x38, 0x93, 0xac, 0xcf, 0x8a,
	0x78, 0x60, 0xfe, 0x84, 0x0a, 0x23, 0xfe, 0x19, 0x0b, 0x55, 0x72, 0x28, 0xf8, 0x34, 0x2c, 0x4e,
	0x42, 0x7d, 0x1e, 0xfc, 0x6c, 0x80, 0x73, 0xcc, 0xce, 0x31, 0x25, 0xf7, 0xa0, 0x15, 0x53, 0x8a,
	0x42, 0x0c, 0x65, 0x11, 0xfb, 0x56, 0xd7, 0xea, 0x79, 0xd1, 0xa6, 0xc6, 0x74, 0xca, 0x2e, 0x80,
	0x3a, 0x1b, 0xca, 0x59, 0x86, 0x7e, 0x43, 0x25, 0x78, 0x0a, 0x39, 0x9e, 0x65, 0x48, 0xf6, 0xa0,
	0xcd, 0xf1, 0x94, 0xa3, 0x18, 0x9b, 0x12, 0xb6, 0xca, 0x68, 0x19, 0xb0, 0xaa, 0x81, 0x5f, 0xb3,
	0x84, 0xa3, 0x18, 0xc6, 0xd2, 0xdf, 0xe8, 0x5a, 0x3d, 0x3b, 0xf2, 0x0c, 0x72, 0x28, 0xc9, 0x6d,
	0x70, 0x05, 0x65, 0x19, 0x0a, 0xdf, 0xe9, 0xda, 0x3d, 0x2f, 0x32, 0x11, 0x39, 0x82, 0xe6, 0x04,
	0x65, 0x3c, 0x8a, 0x65, 0xec, 0xbb, 0x5d, 0xbb, 0xb7, 0x39, 0xd8, 0x0f, 0xaf, 0x22, 0x15, 0xaa,
	0x4e, 0xe1, 0x1b, 0x93, 0xff, 0x2a, 0x95, 0x7c, 0x16, 0x55, 0xd7, 0x3b, 0x2f, 0xa1, 0xbd, 0x70,
	0x44, 0xb6, 0xc0, 0x3e, 0xc7, 0x99, 0x21, 0x5c, 0x7c, 0x92, 0x1d, 0x70, 0xa6, 0xf1, 0x45, 0x5e,
	0x72, 0xd4, 0xc1, 0x8b, 0xc6, 0x73, 0x2b, 0xf8, 0x61, 0xc1, 0xd6, 0x61, 0x2e, 0xc7, 0x8c, 0x27,
	0xdf, 0x30, 0xc2, 0x2f, 0x39, 0x0a, 0xa9, 0x89, 0x8b, 0x8c, 0xa5, 0x02, 0xb5, 0x34, 0x56, 0x49,
	0x5c, 0x83, 0x4a, 0x9d, 0x3b, 0xe0, 0xd1, 0x8b, 0x04, 0x53, 0x39, 0x4c, 0x46, 0xa6, 0x6e, 0x53,
	0x03, 0x47, 0xa3, 0x1a, 0x6d, 0x7b, 0x81, 0xf6, 0x0e, 0x38, 0x42, 0xc6, 0x12, 0x95, 0x50, 0x5e,
	0xa4, 0x83, 0xc2, 0x2a, 0x8e, 0xa3, 0x84, 0x23, 0x95, 0xc3, 0x9c, 0x27, 0xbe, 0xa3, 0xad, 0x2a,
	0xb1, 0x0f, 0x3c, 0x09, 0x24, 0x6c, 0xd7, 0xc6, 0xd4, 0x63, 0x10, 0x02, 0x1b, 0x94, 0x8d, 0xca,
	0xf1, 0xd4, 0xf7, 0xbc, 0x43, 0xa3, 0xde, 0xe1, 0x29, 0x38, 0x73, 0x0b, 0x37, 0x07, 0x77, 0x57,
	0x68, 0x1d, 0xe9, 0xec, 0xe0, 0x57, 0x03, 0x5a, 0x1a, 0x30, 0xca, 0x2c, 0x90, 0xb6, 0x2e, 0x91,
	0xde, 0x83, 0xb6, 0x39, 0x14, 0x48, 0x39, 0x4a, 0x33, 0x42, 0x4b, 0x83, 0xef, 0x15, 0x56, 0xcd,
	0x6c, 0xd7, 0x66, 0xde, 0x05, 0x38, 0xe3, 0x71, 0x2a, 0xb5, 0xd8, 0x5a, 0x1a, 0x4f, 0x21, 0x4a,
	0xe9, 0xd5, 0xf2, 0x2c, 0xaf, 0xaa, 0xfb, 0x87, 0x55, 0x9d, 0x9b, 0x72, 0x73, 0xc1, 0x94, 0x77,
	0xb5, 0x5d, 0x6c, 0xaa, 0x5d, 0x7c, 0xb2, 0x4a, 0x1f, 0x2d, 0xc7, 0xff, 0x59, 0xc9, 0xd7, 0xd0,
	0x36, 0x4d, 0x8c, 0xcd, 0x95, 0x79, 0xd6, 0x5f, 0x99, 0xf7, 0x11, 0xda, 0x11, 0x4e, 0xd9, 0x79,
	0xb5, 0xd6, 0x6b, 0xbc, 0x08, 0x4b, 0x3a, 0x36, 0x96, 0x75, 0x0c, 0xb6, 0xe0, 0x56, 0x59, 0x58,
	0x4f, 0x18, 0x3c, 0x83, 0xed, 0xa3, 0x54, 0x72, 0x26, 0x32, 0xa4, 0x72, 0xfd, 0x76, 0x01, 0x05,
	0x52, 0xbf, 0xf7, 0x4f, 0x7c, 0x0b, 0x7b, 0x63, 0x2a, 0x93, 0xa9, 0x96, 0xb4, 0x19, 0x99, 0x68,
	0xf0, 0xdd, 0x06, 0xf7, 0xad, 0xca, 0x27, 0xa7, 0xe0, 0x55, 0xff, 0x45, 0xe4, 0xe1, 0xd5, 0x75,
	0x2f, 0xbf, 0x08, 0x9d, 0x47, 0x6b, 0xe5, 0x1a, 0x35, 0x6e, 0x90, 0x4f, 0xe5, 0x23, 0x7c, 0x7f,
	0xbd, 0x45, 0xea, 0x3c, 0x58, 0x99, 0x57, 0xd5, 0xfe, 0x0c, 0xae, 0x56, 0x9f, 0x5c, 0x73, 0x69,
	0xc1, 0xf8, 0x4e, 0x6f, 0x75, 0x62, 0x55, 0x3e, 0x01, 0x98, 0x5b, 0x42, 0xae, 0xe1, 0xbd, 0x64,
	0x78, 0xe7, 0xf1, 0x7a, 0xc9, 0x65, 0xab, 0x13, 0x57, 0xfd, 0x98, 0x1d, 0xfc, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0xca, 0xb5, 0xed, 0x03, 0x07, 0x00, 0x00,
}
