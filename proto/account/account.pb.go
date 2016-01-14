// Code generated by protoc-gen-go.
// source: github.com/micro/auth-srv/proto/account/account.proto
// DO NOT EDIT!

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/auth-srv/proto/account/account.proto

It has these top-level messages:
	Record
	ReadRequest
	ReadResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Record struct {
	// uuid
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// service or user
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// service name, username, etc
	ClientId string `protobuf:"bytes,3,opt,name=client_id" json:"client_id,omitempty"`
	// we leave this blank in responses
	// used for update and create
	ClientSecret string            `protobuf:"bytes,4,opt,name=client_secret" json:"client_secret,omitempty"`
	Metadata     map[string]string `protobuf:"bytes,5,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// unix timestamp
	Created int64 `protobuf:"varint,6,opt,name=created" json:"created,omitempty"`
	// unix timestamp
	Updated int64 `protobuf:"varint,7,opt,name=updated" json:"updated,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ReadRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ReadResponse struct {
	Account *Record `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadResponse) GetAccount() *Record {
	if m != nil {
		return m.Account
	}
	return nil
}

type CreateRequest struct {
	// If id is blank, one will be generated
	Account *Record `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateRequest) GetAccount() *Record {
	if m != nil {
		return m.Account
	}
	return nil
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type UpdateRequest struct {
	// Id and client_id cannot be changed
	Account *Record `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetAccount() *Record {
	if m != nil {
		return m.Account
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type SearchRequest struct {
	ClientId string `protobuf:"bytes,1,opt,name=client_id" json:"client_id,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Limit    int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset   int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SearchResponse struct {
	Accounts []*Record `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SearchResponse) GetAccounts() []*Record {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "Record")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type accountClient struct {
	c           client.Client
	serviceName string
}

func NewAccountClient(serviceName string, c client.Client) AccountClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "account"
	}
	return &accountClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler) {
	s.Handle(s.NewHandler(&Account{hdlr}))
}

type Account struct {
	AccountHandler
}

func (h *Account) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.AccountHandler.Read(ctx, in, out)
}

func (h *Account) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AccountHandler.Create(ctx, in, out)
}

func (h *Account) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.AccountHandler.Update(ctx, in, out)
}

func (h *Account) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.AccountHandler.Delete(ctx, in, out)
}

func (h *Account) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.AccountHandler.Search(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0x87, 0x8d, 0xd1, 0x44, 0x4f, 0x4c, 0x6c, 0x07, 0x84, 0x68, 0x6f, 0x24, 0x50, 0x6a, 0x91,
	0x4e, 0xc0, 0x52, 0x28, 0xbd, 0x2b, 0x6d, 0x2f, 0x4b, 0xc1, 0xd2, 0xeb, 0x32, 0x4e, 0xc6, 0x35,
	0x6c, 0xfe, 0x6d, 0x32, 0x11, 0x7c, 0xc3, 0x7d, 0x90, 0x7d, 0x90, 0x9d, 0xcc, 0x4c, 0xd8, 0x44,
	0x76, 0x61, 0xaf, 0x24, 0xbf, 0xf9, 0xce, 0xe1, 0x9c, 0xef, 0x20, 0x7c, 0xb9, 0x89, 0xf9, 0xa9,
	0x3e, 0x60, 0x9a, 0xa7, 0x61, 0x1a, 0xd3, 0x32, 0x0f, 0x49, 0xcd, 0x4f, 0x9f, 0xaa, 0xf2, 0x1c,
	0x16, 0x65, 0xce, 0xc5, 0x27, 0xa5, 0x79, 0x9d, 0xf1, 0xf6, 0x17, 0xcb, 0x34, 0xb8, 0x37, 0xc0,
	0xda, 0x33, 0x9a, 0x97, 0x11, 0x02, 0x18, 0xc6, 0x91, 0x6f, 0xac, 0x8d, 0xcd, 0x14, 0xcd, 0x60,
	0xc4, 0x2f, 0x05, 0xf3, 0x87, 0xf2, 0xeb, 0x2d, 0x4c, 0x69, 0x12, 0xb3, 0x8c, 0xff, 0x17, 0x80,
	0x29, 0xa3, 0x05, 0xb8, 0x3a, 0xaa, 0x18, 0x2d, 0x19, 0xf7, 0x47, 0x32, 0xfe, 0x00, 0x93, 0x94,
	0x71, 0x12, 0x11, 0x4e, 0xfc, 0xf1, 0xda, 0xdc, 0x38, 0xbb, 0x05, 0x56, 0xed, 0xf1, 0x6f, 0x9d,
	0xff, 0xca, 0x78, 0x79, 0x41, 0x73, 0xb0, 0x45, 0x19, 0xe1, 0x2c, 0xf2, 0x2d, 0x51, 0x69, 0x36,
	0x41, 0x5d, 0x44, 0x32, 0xb0, 0x9b, 0x60, 0x15, 0x82, 0xdb, 0x2f, 0x71, 0xc0, 0xbc, 0x65, 0x17,
	0x3d, 0xa0, 0x0b, 0xe3, 0x33, 0x49, 0x6a, 0x3d, 0xe1, 0xb7, 0xe1, 0x57, 0x23, 0x58, 0x82, 0xb3,
	0x67, 0x24, 0xda, 0xb3, 0xbb, 0x9a, 0x55, 0xbc, 0xbb, 0x4e, 0xb0, 0x81, 0x99, 0x7a, 0xaa, 0x8a,
	0x3c, 0xab, 0x18, 0xf2, 0xc1, 0xd6, 0x1a, 0x24, 0xe0, 0xec, 0x6c, 0x3d, 0x65, 0xf0, 0x11, 0xdc,
	0x1f, 0x72, 0xae, 0xb6, 0xcd, 0xcb, 0xe8, 0x1b, 0xf0, 0x5a, 0x54, 0xb5, 0x6d, 0x8a, 0xff, 0xc9,
	0x1d, 0x5e, 0x55, 0xdc, 0xa2, 0xba, 0xf8, 0x1d, 0xb8, 0x3f, 0x59, 0xc2, 0x9e, 0x8a, 0xbb, 0x0b,
	0x08, 0xbc, 0x7d, 0xd4, 0xf8, 0x1f, 0x70, 0xff, 0x32, 0x52, 0xd2, 0x53, 0x8b, 0xf7, 0x8e, 0xf4,
	0xdc, 0x15, 0x85, 0xb2, 0x24, 0x4e, 0x63, 0x2e, 0x2f, 0x68, 0x22, 0x0f, 0xac, 0xfc, 0x78, 0xac,
	0xf4, 0xe9, 0xcc, 0x60, 0x0b, 0x5e, 0xdb, 0x50, 0x5b, 0x5a, 0xc2, 0x44, 0x4f, 0x5f, 0x89, 0x86,
	0x66, 0x67, 0xfc, 0xdd, 0x83, 0x01, 0xf6, 0x77, 0xf5, 0x86, 0xde, 0xc3, 0xa8, 0x91, 0x8b, 0x66,
	0xb8, 0xa3, 0x7f, 0xe5, 0xe2, 0xae, 0xf1, 0x60, 0x80, 0xb6, 0x60, 0x29, 0x5d, 0xc8, 0xc3, 0x3d,
	0xc5, 0xab, 0x39, 0xbe, 0xf2, 0x28, 0x61, 0xa5, 0x47, 0xc0, 0x3d, 0xa5, 0x02, 0xbe, 0xf2, 0x26,
	0x61, 0x25, 0x47, 0xc0, 0x3d, 0x85, 0x02, 0xbe, 0xb2, 0x26, 0x61, 0xb5, 0xa6, 0x80, 0x7b, 0x02,
	0x05, 0xdc, 0xdf, 0x3f, 0x18, 0x1c, 0x2c, 0xf9, 0x27, 0xf9, 0xfc, 0x18, 0x00, 0x00, 0xff, 0xff,
	0x5c, 0x56, 0xaa, 0x9c, 0x5d, 0x03, 0x00, 0x00,
}
