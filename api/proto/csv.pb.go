// Code generated by protoc-gen-go. DO NOT EDIT.
// source: csv.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetTableRequest struct {
	// Deprecated!
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Rows                 uint64   `protobuf:"varint,2,opt,name=rows,proto3" json:"rows,omitempty"`
	StartRow             uint64   `protobuf:"varint,3,opt,name=start_row,json=startRow,proto3" json:"start_row,omitempty"`
	ClientId             string   `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	FlowId               string   `protobuf:"bytes,5,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Artifact             string   `protobuf:"bytes,6,opt,name=artifact,proto3" json:"artifact,omitempty"`
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	HuntId               string   `protobuf:"bytes,8,opt,name=hunt_id,json=huntId,proto3" json:"hunt_id,omitempty"`
	NotebookId           string   `protobuf:"bytes,9,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	CellId               string   `protobuf:"bytes,10,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	TableId              int64    `protobuf:"varint,11,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTableRequest) Reset()         { *m = GetTableRequest{} }
func (m *GetTableRequest) String() string { return proto.CompactTextString(m) }
func (*GetTableRequest) ProtoMessage()    {}
func (*GetTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e185232356a2266, []int{0}
}

func (m *GetTableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTableRequest.Unmarshal(m, b)
}
func (m *GetTableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTableRequest.Marshal(b, m, deterministic)
}
func (m *GetTableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTableRequest.Merge(m, src)
}
func (m *GetTableRequest) XXX_Size() int {
	return xxx_messageInfo_GetTableRequest.Size(m)
}
func (m *GetTableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTableRequest proto.InternalMessageInfo

func (m *GetTableRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GetTableRequest) GetRows() uint64 {
	if m != nil {
		return m.Rows
	}
	return 0
}

func (m *GetTableRequest) GetStartRow() uint64 {
	if m != nil {
		return m.StartRow
	}
	return 0
}

func (m *GetTableRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetTableRequest) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *GetTableRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

func (m *GetTableRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetTableRequest) GetHuntId() string {
	if m != nil {
		return m.HuntId
	}
	return ""
}

func (m *GetTableRequest) GetNotebookId() string {
	if m != nil {
		return m.NotebookId
	}
	return ""
}

func (m *GetTableRequest) GetCellId() string {
	if m != nil {
		return m.CellId
	}
	return ""
}

func (m *GetTableRequest) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

type Row struct {
	Cell                 []string `protobuf:"bytes,1,rep,name=cell,proto3" json:"cell,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e185232356a2266, []int{1}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetCell() []string {
	if m != nil {
		return m.Cell
	}
	return nil
}

type GetTableResponse struct {
	Columns              []string `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows                 []*Row   `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTableResponse) Reset()         { *m = GetTableResponse{} }
func (m *GetTableResponse) String() string { return proto.CompactTextString(m) }
func (*GetTableResponse) ProtoMessage()    {}
func (*GetTableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e185232356a2266, []int{2}
}

func (m *GetTableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTableResponse.Unmarshal(m, b)
}
func (m *GetTableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTableResponse.Marshal(b, m, deterministic)
}
func (m *GetTableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTableResponse.Merge(m, src)
}
func (m *GetTableResponse) XXX_Size() int {
	return xxx_messageInfo_GetTableResponse.Size(m)
}
func (m *GetTableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTableResponse proto.InternalMessageInfo

func (m *GetTableResponse) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *GetTableResponse) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTableRequest)(nil), "proto.GetTableRequest")
	proto.RegisterType((*Row)(nil), "proto.Row")
	proto.RegisterType((*GetTableResponse)(nil), "proto.GetTableResponse")
}

func init() {
	proto.RegisterFile("csv.proto", fileDescriptor_3e185232356a2266)
}

var fileDescriptor_3e185232356a2266 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x52, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x56, 0x48, 0xda, 0xdd, 0xf5, 0x0a, 0x81, 0x5c, 0x50, 0xb3, 0xad, 0x44, 0xad, 0xbd, 0xb0,
	0x17, 0x12, 0x41, 0x2f, 0x1c, 0x51, 0x90, 0x4a, 0x73, 0x4d, 0xab, 0x1e, 0x10, 0x52, 0xe5, 0x75,
	0xbc, 0x8d, 0x45, 0xd6, 0x13, 0xe2, 0xd9, 0x9a, 0xbe, 0x16, 0xcf, 0xc0, 0x93, 0xc0, 0x6b, 0x70,
	0x40, 0x63, 0xb3, 0xc0, 0x69, 0x3c, 0xdf, 0xcf, 0x64, 0x9c, 0xcf, 0x6c, 0xa6, 0xdc, 0x7d, 0x31,
	0x8c, 0x80, 0xc0, 0x0f, 0x42, 0x39, 0x79, 0x16, 0x4a, 0xe9, 0xf4, 0x56, 0x5a, 0x34, 0x2a, 0x92,
	0xcb, 0x6f, 0x29, 0x7b, 0xf2, 0x41, 0xe3, 0xb5, 0x5c, 0xf7, 0xba, 0xd1, 0x5f, 0x76, 0xda, 0x21,
	0xff, 0xc4, 0xb2, 0x41, 0x62, 0x97, 0x27, 0x22, 0x59, 0xcd, 0xaa, 0xcb, 0x1f, 0xbf, 0x7e, 0x7e,
	0x4f, 0x2a, 0xfe, 0xee, 0xba, 0xd3, 0x82, 0x70, 0xe1, 0x0d, 0x76, 0xc6, 0x0a, 0xec, 0xb4, 0x50,
	0xbd, 0xd1, 0x16, 0x5f, 0x3a, 0x71, 0x73, 0x71, 0x25, 0xb0, 0x93, 0x28, 0x14, 0x58, 0x94, 0xc6,
	0x3a, 0xf1, 0xfe, 0xea, 0x46, 0x6c, 0x4c, 0xaf, 0x05, 0x82, 0x18, 0xb5, 0x6c, 0x8b, 0x26, 0x4c,
	0xe5, 0xe7, 0x2c, 0x1b, 0xc1, 0xbb, 0xfc, 0x91, 0x48, 0x56, 0x59, 0x75, 0x16, 0xa6, 0x2f, 0xf8,
	0xf1, 0x25, 0x78, 0xb1, 0x95, 0xf6, 0x41, 0x10, 0x49, 0x9e, 0x8d, 0x46, 0xd5, 0x15, 0x4d, 0x10,
	0xf3, 0xb7, 0x6c, 0xe6, 0x50, 0x8e, 0x78, 0x3b, 0x82, 0xcf, 0xd3, 0xe0, 0x3c, 0x0d, 0xce, 0xe7,
	0xfc, 0xe8, 0xc2, 0x8c, 0x0e, 0xc9, 0xf6, 0x9f, 0x6b, 0x1a, 0xd4, 0x0d, 0x78, 0x7e, 0xca, 0x66,
	0x71, 0xc9, 0x5b, 0xd3, 0xe6, 0x19, 0xdd, 0xa8, 0x99, 0x46, 0xa0, 0x6e, 0xf9, 0x31, 0x9b, 0x6c,
	0x7a, 0xf0, 0x44, 0x1d, 0x04, 0xea, 0x90, 0xda, 0xba, 0xe5, 0x27, 0x6c, 0x2a, 0x47, 0x34, 0x1b,
	0xa9, 0x30, 0x3f, 0x8c, 0xa6, 0x7d, 0xcf, 0x39, 0xcb, 0xf0, 0x61, 0xd0, 0xf9, 0x24, 0xe0, 0xe1,
	0x4c, 0x83, 0xba, 0x5d, 0xfc, 0xc6, 0x34, 0x0e, 0xa2, 0xb6, 0x6e, 0xf9, 0x19, 0x9b, 0x5b, 0x40,
	0xbd, 0x06, 0xf8, 0x4c, 0xe4, 0x2c, 0x90, 0x6c, 0x0f, 0xc5, 0x15, 0x94, 0xee, 0x7b, 0x22, 0x59,
	0x74, 0x52, 0x5b, 0xb7, 0x7c, 0xc1, 0xa6, 0x48, 0xa9, 0x10, 0x33, 0x17, 0xc9, 0x2a, 0x6d, 0x26,
	0xa1, 0xaf, 0xdb, 0xe5, 0x82, 0xa5, 0x74, 0x35, 0xce, 0x32, 0xd2, 0xe6, 0x89, 0x48, 0x69, 0x11,
	0x3a, 0x2f, 0x25, 0x7b, 0xfa, 0x2f, 0x4e, 0x37, 0x80, 0x75, 0x9a, 0xbf, 0x62, 0x13, 0x05, 0xfd,
	0x6e, 0x6b, 0x5d, 0x94, 0x56, 0x47, 0xe1, 0xd7, 0x3d, 0xe6, 0x73, 0x8a, 0xf4, 0x0f, 0xd5, 0xec,
	0x35, 0xfc, 0xc5, 0xdf, 0x80, 0xd2, 0xd5, 0xfc, 0x0d, 0x8b, 0x0f, 0xa5, 0x68, 0xc0, 0xc7, 0x2c,
	0xaa, 0xd7, 0x1f, 0x4b, 0xef, 0x7d, 0x71, 0xaf, 0x7b, 0x50, 0xa6, 0xd5, 0x5f, 0x0b, 0x05, 0xdb,
	0xf2, 0x0e, 0x7a, 0x69, 0xef, 0xca, 0x08, 0x8e, 0x72, 0x40, 0x18, 0x4b, 0x39, 0x98, 0x32, 0x98,
	0xd7, 0x87, 0xa1, 0x9c, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x13, 0xff, 0x8b, 0xbd, 0x96, 0x02,
	0x00, 0x00,
}
