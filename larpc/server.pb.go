// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package larpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type ContractType int32

const (
	ContractType_FUNDED   ContractType = 0
	ContractType_UNFUNDED ContractType = 1
)

var ContractType_name = map[int32]string{
	0: "FUNDED",
	1: "UNFUNDED",
}

var ContractType_value = map[string]int32{
	"FUNDED":   0,
	"UNFUNDED": 1,
}

func (x ContractType) String() string {
	return proto.EnumName(ContractType_name, int32(x))
}

func (ContractType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

// Contract is the type of our contract, used to marshal/unmarshal
// and send between hosts
type ServerContract struct {
	Uuid                 string       `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Asset                string       `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount               float64      `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountSats           int64        `protobuf:"varint,4,opt,name=amount_sats,json=amountSats,proto3" json:"amount_sats,omitempty"`
	ClientHost           string       `protobuf:"bytes,5,opt,name=client_host,json=clientHost,proto3" json:"client_host,omitempty"`
	MarginPayReq         string       `protobuf:"bytes,6,opt,name=margin_pay_req,json=marginPayReq,proto3" json:"margin_pay_req,omitempty"`
	InitiatingPayReq     string       `protobuf:"bytes,7,opt,name=initiating_pay_req,json=initiatingPayReq,proto3" json:"initiating_pay_req,omitempty"`
	MarginPaid           bool         `protobuf:"varint,8,opt,name=margin_paid,json=marginPaid,proto3" json:"margin_paid,omitempty"`
	InitiatingPaid       bool         `protobuf:"varint,9,opt,name=initiating_paid,json=initiatingPaid,proto3" json:"initiating_paid,omitempty"`
	ContractType         ContractType `protobuf:"varint,10,opt,name=contract_type,json=contractType,proto3,enum=ladrpc.ContractType" json:"contract_type,omitempty"`
	NumUpdates           int64        `protobuf:"varint,11,opt,name=num_updates,json=numUpdates,proto3" json:"num_updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ServerContract) Reset()         { *m = ServerContract{} }
func (m *ServerContract) String() string { return proto.CompactTextString(m) }
func (*ServerContract) ProtoMessage()    {}
func (*ServerContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *ServerContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerContract.Unmarshal(m, b)
}
func (m *ServerContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerContract.Marshal(b, m, deterministic)
}
func (m *ServerContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerContract.Merge(m, src)
}
func (m *ServerContract) XXX_Size() int {
	return xxx_messageInfo_ServerContract.Size(m)
}
func (m *ServerContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerContract.DiscardUnknown(m)
}

var xxx_messageInfo_ServerContract proto.InternalMessageInfo

func (m *ServerContract) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ServerContract) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ServerContract) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ServerContract) GetAmountSats() int64 {
	if m != nil {
		return m.AmountSats
	}
	return 0
}

func (m *ServerContract) GetClientHost() string {
	if m != nil {
		return m.ClientHost
	}
	return ""
}

func (m *ServerContract) GetMarginPayReq() string {
	if m != nil {
		return m.MarginPayReq
	}
	return ""
}

func (m *ServerContract) GetInitiatingPayReq() string {
	if m != nil {
		return m.InitiatingPayReq
	}
	return ""
}

func (m *ServerContract) GetMarginPaid() bool {
	if m != nil {
		return m.MarginPaid
	}
	return false
}

func (m *ServerContract) GetInitiatingPaid() bool {
	if m != nil {
		return m.InitiatingPaid
	}
	return false
}

func (m *ServerContract) GetContractType() ContractType {
	if m != nil {
		return m.ContractType
	}
	return ContractType_FUNDED
}

func (m *ServerContract) GetNumUpdates() int64 {
	if m != nil {
		return m.NumUpdates
	}
	return 0
}

// Payment is a payment type, used to marshal/unmarshal from the db
type Payment struct {
	ContractUuid   string `protobuf:"bytes,1,opt,name=contract_uuid,json=contractUuid,proto3" json:"contract_uuid,omitempty"`
	AmountSat      int64  `protobuf:"varint,2,opt,name=amount_sat,json=amountSat,proto3" json:"amount_sat,omitempty"`
	PaymentRequest string `protobuf:"bytes,3,opt,name=payment_request,json=paymentRequest,proto3" json:"payment_request,omitempty"`
	// if true, this payment was outbound, ie paid by us
	Outbound             bool     `protobuf:"varint,4,opt,name=outbound,proto3" json:"outbound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetContractUuid() string {
	if m != nil {
		return m.ContractUuid
	}
	return ""
}

func (m *Payment) GetAmountSat() int64 {
	if m != nil {
		return m.AmountSat
	}
	return 0
}

func (m *Payment) GetPaymentRequest() string {
	if m != nil {
		return m.PaymentRequest
	}
	return ""
}

func (m *Payment) GetOutbound() bool {
	if m != nil {
		return m.Outbound
	}
	return false
}

type Quote struct {
	PercentMargin        float64  `protobuf:"fixed64,1,opt,name=percent_margin,json=percentMargin,proto3" json:"percent_margin,omitempty"`
	AmountSats           int64    `protobuf:"varint,2,opt,name=amount_sats,json=amountSats,proto3" json:"amount_sats,omitempty"`
	AssetPrice           float64  `protobuf:"fixed64,3,opt,name=asset_price,json=assetPrice,proto3" json:"asset_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quote) Reset()         { *m = Quote{} }
func (m *Quote) String() string { return proto.CompactTextString(m) }
func (*Quote) ProtoMessage()    {}
func (*Quote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *Quote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quote.Unmarshal(m, b)
}
func (m *Quote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quote.Marshal(b, m, deterministic)
}
func (m *Quote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quote.Merge(m, src)
}
func (m *Quote) XXX_Size() int {
	return xxx_messageInfo_Quote.Size(m)
}
func (m *Quote) XXX_DiscardUnknown() {
	xxx_messageInfo_Quote.DiscardUnknown(m)
}

var xxx_messageInfo_Quote proto.InternalMessageInfo

func (m *Quote) GetPercentMargin() float64 {
	if m != nil {
		return m.PercentMargin
	}
	return 0
}

func (m *Quote) GetAmountSats() int64 {
	if m != nil {
		return m.AmountSats
	}
	return 0
}

func (m *Quote) GetAssetPrice() float64 {
	if m != nil {
		return m.AssetPrice
	}
	return 0
}

type Price struct {
	Asset                string   `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *Price) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// ServerNewContractRequest is used to initiate a new contract
// with another host
type ServerNewContractRequest struct {
	Asset                string       `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount               float64      `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Host                 string       `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	ContractType         ContractType `protobuf:"varint,4,opt,name=contract_type,json=contractType,proto3,enum=ladrpc.ContractType" json:"contract_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ServerNewContractRequest) Reset()         { *m = ServerNewContractRequest{} }
func (m *ServerNewContractRequest) String() string { return proto.CompactTextString(m) }
func (*ServerNewContractRequest) ProtoMessage()    {}
func (*ServerNewContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *ServerNewContractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerNewContractRequest.Unmarshal(m, b)
}
func (m *ServerNewContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerNewContractRequest.Marshal(b, m, deterministic)
}
func (m *ServerNewContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerNewContractRequest.Merge(m, src)
}
func (m *ServerNewContractRequest) XXX_Size() int {
	return xxx_messageInfo_ServerNewContractRequest.Size(m)
}
func (m *ServerNewContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerNewContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerNewContractRequest proto.InternalMessageInfo

func (m *ServerNewContractRequest) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ServerNewContractRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ServerNewContractRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ServerNewContractRequest) GetContractType() ContractType {
	if m != nil {
		return m.ContractType
	}
	return ContractType_FUNDED
}

// If successful, the ServerNewContractResponse returns the created contract
type ServerNewContractResponse struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	MarginPayReq         string   `protobuf:"bytes,2,opt,name=margin_pay_req,json=marginPayReq,proto3" json:"margin_pay_req,omitempty"`
	InitiatingPayReq     string   `protobuf:"bytes,3,opt,name=initiating_pay_req,json=initiatingPayReq,proto3" json:"initiating_pay_req,omitempty"`
	PercentMargin        float64  `protobuf:"fixed64,4,opt,name=percent_margin,json=percentMargin,proto3" json:"percent_margin,omitempty"`
	AssetPrice           float64  `protobuf:"fixed64,5,opt,name=asset_price,json=assetPrice,proto3" json:"asset_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerNewContractResponse) Reset()         { *m = ServerNewContractResponse{} }
func (m *ServerNewContractResponse) String() string { return proto.CompactTextString(m) }
func (*ServerNewContractResponse) ProtoMessage()    {}
func (*ServerNewContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *ServerNewContractResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerNewContractResponse.Unmarshal(m, b)
}
func (m *ServerNewContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerNewContractResponse.Marshal(b, m, deterministic)
}
func (m *ServerNewContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerNewContractResponse.Merge(m, src)
}
func (m *ServerNewContractResponse) XXX_Size() int {
	return xxx_messageInfo_ServerNewContractResponse.Size(m)
}
func (m *ServerNewContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerNewContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerNewContractResponse proto.InternalMessageInfo

func (m *ServerNewContractResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ServerNewContractResponse) GetMarginPayReq() string {
	if m != nil {
		return m.MarginPayReq
	}
	return ""
}

func (m *ServerNewContractResponse) GetInitiatingPayReq() string {
	if m != nil {
		return m.InitiatingPayReq
	}
	return ""
}

func (m *ServerNewContractResponse) GetPercentMargin() float64 {
	if m != nil {
		return m.PercentMargin
	}
	return 0
}

func (m *ServerNewContractResponse) GetAssetPrice() float64 {
	if m != nil {
		return m.AssetPrice
	}
	return 0
}

type ServerCloseContractRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerCloseContractRequest) Reset()         { *m = ServerCloseContractRequest{} }
func (m *ServerCloseContractRequest) String() string { return proto.CompactTextString(m) }
func (*ServerCloseContractRequest) ProtoMessage()    {}
func (*ServerCloseContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{6}
}

func (m *ServerCloseContractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerCloseContractRequest.Unmarshal(m, b)
}
func (m *ServerCloseContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerCloseContractRequest.Marshal(b, m, deterministic)
}
func (m *ServerCloseContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerCloseContractRequest.Merge(m, src)
}
func (m *ServerCloseContractRequest) XXX_Size() int {
	return xxx_messageInfo_ServerCloseContractRequest.Size(m)
}
func (m *ServerCloseContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerCloseContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerCloseContractRequest proto.InternalMessageInfo

func (m *ServerCloseContractRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type ServerCloseContractResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerCloseContractResponse) Reset()         { *m = ServerCloseContractResponse{} }
func (m *ServerCloseContractResponse) String() string { return proto.CompactTextString(m) }
func (*ServerCloseContractResponse) ProtoMessage()    {}
func (*ServerCloseContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{7}
}

func (m *ServerCloseContractResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerCloseContractResponse.Unmarshal(m, b)
}
func (m *ServerCloseContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerCloseContractResponse.Marshal(b, m, deterministic)
}
func (m *ServerCloseContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerCloseContractResponse.Merge(m, src)
}
func (m *ServerCloseContractResponse) XXX_Size() int {
	return xxx_messageInfo_ServerCloseContractResponse.Size(m)
}
func (m *ServerCloseContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerCloseContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerCloseContractResponse proto.InternalMessageInfo

type ServerListAssetsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerListAssetsRequest) Reset()         { *m = ServerListAssetsRequest{} }
func (m *ServerListAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*ServerListAssetsRequest) ProtoMessage()    {}
func (*ServerListAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{8}
}

func (m *ServerListAssetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerListAssetsRequest.Unmarshal(m, b)
}
func (m *ServerListAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerListAssetsRequest.Marshal(b, m, deterministic)
}
func (m *ServerListAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerListAssetsRequest.Merge(m, src)
}
func (m *ServerListAssetsRequest) XXX_Size() int {
	return xxx_messageInfo_ServerListAssetsRequest.Size(m)
}
func (m *ServerListAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerListAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerListAssetsRequest proto.InternalMessageInfo

type ServerListAssetsResponse struct {
	SupportedAssets      []string `protobuf:"bytes,1,rep,name=supported_assets,json=supportedAssets,proto3" json:"supported_assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerListAssetsResponse) Reset()         { *m = ServerListAssetsResponse{} }
func (m *ServerListAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*ServerListAssetsResponse) ProtoMessage()    {}
func (*ServerListAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{9}
}

func (m *ServerListAssetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerListAssetsResponse.Unmarshal(m, b)
}
func (m *ServerListAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerListAssetsResponse.Marshal(b, m, deterministic)
}
func (m *ServerListAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerListAssetsResponse.Merge(m, src)
}
func (m *ServerListAssetsResponse) XXX_Size() int {
	return xxx_messageInfo_ServerListAssetsResponse.Size(m)
}
func (m *ServerListAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerListAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerListAssetsResponse proto.InternalMessageInfo

func (m *ServerListAssetsResponse) GetSupportedAssets() []string {
	if m != nil {
		return m.SupportedAssets
	}
	return nil
}

func init() {
	proto.RegisterEnum("ladrpc.ContractType", ContractType_name, ContractType_value)
	proto.RegisterType((*ServerContract)(nil), "ladrpc.ServerContract")
	proto.RegisterType((*Payment)(nil), "ladrpc.Payment")
	proto.RegisterType((*Quote)(nil), "ladrpc.Quote")
	proto.RegisterType((*Price)(nil), "ladrpc.Price")
	proto.RegisterType((*ServerNewContractRequest)(nil), "ladrpc.ServerNewContractRequest")
	proto.RegisterType((*ServerNewContractResponse)(nil), "ladrpc.ServerNewContractResponse")
	proto.RegisterType((*ServerCloseContractRequest)(nil), "ladrpc.ServerCloseContractRequest")
	proto.RegisterType((*ServerCloseContractResponse)(nil), "ladrpc.ServerCloseContractResponse")
	proto.RegisterType((*ServerListAssetsRequest)(nil), "ladrpc.ServerListAssetsRequest")
	proto.RegisterType((*ServerListAssetsResponse)(nil), "ladrpc.ServerListAssetsResponse")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4f, 0x53, 0xe3, 0x36,
	0x1c, 0xad, 0xf2, 0x8f, 0xe4, 0x97, 0x10, 0x52, 0x95, 0x01, 0x93, 0x96, 0xc1, 0x35, 0x74, 0x9a,
	0x32, 0x2d, 0xe9, 0xc0, 0xa9, 0xdc, 0xda, 0x42, 0xa7, 0x87, 0x96, 0xa1, 0xa6, 0xb9, 0xec, 0xc5,
	0x23, 0x6c, 0x4d, 0x56, 0xb3, 0x8e, 0x24, 0x2c, 0x19, 0x26, 0xd7, 0xfd, 0x00, 0x7b, 0xd8, 0x3d,
	0xec, 0x87, 0xda, 0xe3, 0x5e, 0xf6, 0x03, 0xec, 0x07, 0xd9, 0xb1, 0x64, 0x27, 0x0e, 0x49, 0x18,
	0x6e, 0xd2, 0xd3, 0x4f, 0xef, 0x27, 0xbd, 0xf7, 0x64, 0x43, 0x47, 0xd1, 0xe4, 0x9e, 0x26, 0x27,
	0x32, 0x11, 0x5a, 0xe0, 0x46, 0x4c, 0xa2, 0x44, 0x86, 0xfd, 0xef, 0xc6, 0x42, 0x8c, 0x63, 0x3a,
	0x24, 0x92, 0x0d, 0x09, 0xe7, 0x42, 0x13, 0xcd, 0x04, 0x57, 0xb6, 0xca, 0x7b, 0x53, 0x85, 0xee,
	0x8d, 0xd9, 0xf6, 0xa7, 0xe0, 0x3a, 0x21, 0xa1, 0xc6, 0x18, 0x6a, 0x69, 0xca, 0x22, 0x07, 0xb9,
	0x68, 0xd0, 0xf2, 0xcd, 0x18, 0x6f, 0x43, 0x9d, 0x28, 0x45, 0xb5, 0x53, 0x31, 0xa0, 0x9d, 0xe0,
	0x1d, 0x68, 0x90, 0x89, 0x48, 0xb9, 0x76, 0xaa, 0x2e, 0x1a, 0x20, 0x3f, 0x9f, 0xe1, 0x03, 0x68,
	0xdb, 0x51, 0xa0, 0x88, 0x56, 0x4e, 0xcd, 0x45, 0x83, 0xaa, 0x0f, 0x16, 0xba, 0x21, 0x5a, 0x65,
	0x05, 0x61, 0xcc, 0x28, 0xd7, 0xc1, 0x4b, 0xa1, 0xb4, 0x53, 0x37, 0xa4, 0x60, 0xa1, 0xbf, 0x85,
	0xd2, 0xf8, 0x08, 0xba, 0x13, 0x92, 0x8c, 0x19, 0x0f, 0x24, 0x99, 0x06, 0x09, 0xbd, 0x73, 0x1a,
	0xa6, 0xa6, 0x63, 0xd1, 0x6b, 0x32, 0xf5, 0xe9, 0x1d, 0xfe, 0x19, 0x30, 0xe3, 0x4c, 0x33, 0xa2,
	0x19, 0x1f, 0xcf, 0x2a, 0x37, 0x4c, 0x65, 0x6f, 0xbe, 0x92, 0x57, 0x1f, 0x40, 0x7b, 0xc6, 0xc9,
	0x22, 0xa7, 0xe9, 0xa2, 0x41, 0xd3, 0x87, 0x82, 0x90, 0x45, 0xf8, 0x47, 0xd8, 0x5a, 0xa0, 0x63,
	0x91, 0xd3, 0x32, 0x45, 0xdd, 0x32, 0x17, 0x8b, 0xf0, 0x6f, 0xb0, 0x19, 0xe6, 0x6a, 0x05, 0x7a,
	0x2a, 0xa9, 0x03, 0x2e, 0x1a, 0x74, 0x4f, 0xb7, 0x4f, 0xac, 0xe4, 0x27, 0x85, 0x94, 0xff, 0x4f,
	0x25, 0xf5, 0x3b, 0x61, 0x69, 0x96, 0x1d, 0x82, 0xa7, 0x93, 0x20, 0x95, 0x11, 0xd1, 0x54, 0x39,
	0x6d, 0x2b, 0x0d, 0x4f, 0x27, 0x23, 0x8b, 0x78, 0x6f, 0x11, 0x6c, 0x5c, 0x93, 0xe9, 0x84, 0x72,
	0x8d, 0x0f, 0x4b, 0x7d, 0x4a, 0x96, 0xcc, 0x18, 0x47, 0x99, 0x35, 0xfb, 0x00, 0x73, 0xb1, 0x8d,
	0x3f, 0x55, 0xbf, 0x35, 0xd3, 0x3a, 0xbb, 0x94, 0xb4, 0x74, 0x99, 0x38, 0x29, 0x55, 0xd6, 0xac,
	0x96, 0xdf, 0xcd, 0x61, 0xdf, 0xa2, 0xb8, 0x0f, 0x4d, 0x91, 0xea, 0x5b, 0x91, 0xf2, 0xc8, 0x38,
	0xd6, 0xf4, 0x67, 0x73, 0x4f, 0x42, 0xfd, 0xbf, 0x54, 0x68, 0x8a, 0x7f, 0x80, 0xae, 0xa4, 0x49,
	0x98, 0xb1, 0x59, 0xe1, 0xcc, 0x91, 0x90, 0xbf, 0x99, 0xa3, 0xff, 0x1a, 0xf0, 0x71, 0x00, 0x2a,
	0xab, 0x02, 0x60, 0x22, 0x14, 0xc8, 0x84, 0x85, 0x34, 0x8f, 0x0f, 0x18, 0xe8, 0x3a, 0x43, 0xbc,
	0x33, 0xa8, 0x9b, 0xc1, 0x3c, 0x79, 0xa8, 0x9c, 0xbc, 0x6d, 0xa8, 0xdf, 0x93, 0x38, 0xa5, 0x86,
	0x1a, 0xf9, 0x76, 0xe2, 0xbd, 0x47, 0xe0, 0xd8, 0x30, 0x5f, 0xd1, 0x87, 0xc2, 0x84, 0xe2, 0x7e,
	0xab, 0x89, 0xe6, 0x11, 0xae, 0x2c, 0x44, 0x18, 0x43, 0xcd, 0x44, 0xd3, 0x6a, 0x65, 0xc6, 0xcb,
	0xb6, 0xd7, 0x9e, 0x6b, 0xbb, 0xf7, 0x01, 0xc1, 0xde, 0x8a, 0x93, 0x29, 0x29, 0xb8, 0xa2, 0x2b,
	0x5f, 0xdc, 0xf2, 0x0b, 0xa8, 0x3c, 0xfb, 0x05, 0x54, 0xd7, 0xbc, 0x80, 0x65, 0xf7, 0x6a, 0xeb,
	0xdc, 0x2b, 0x99, 0x53, 0x5f, 0x32, 0xe7, 0x57, 0xe8, 0xe7, 0xdf, 0x8c, 0x58, 0x28, 0xfa, 0x58,
	0xe8, 0x15, 0xb7, 0xf1, 0xf6, 0xe1, 0xdb, 0x95, 0x3b, 0xac, 0x00, 0xde, 0x1e, 0xec, 0xda, 0xe5,
	0x7f, 0x98, 0xd2, 0xbf, 0x67, 0x8d, 0x54, 0xce, 0xe6, 0x5d, 0x16, 0x96, 0x96, 0x97, 0x72, 0xdd,
	0x7e, 0x82, 0x9e, 0x4a, 0xa5, 0x14, 0x89, 0xa6, 0x51, 0x60, 0xce, 0xa7, 0x1c, 0xe4, 0x56, 0x07,
	0x2d, 0x7f, 0x6b, 0x86, 0xdb, 0x2d, 0xc7, 0x03, 0xe8, 0x94, 0xed, 0xc1, 0x00, 0x8d, 0xbf, 0x46,
	0x57, 0x17, 0x97, 0x17, 0xbd, 0xaf, 0x70, 0x07, 0x9a, 0xa3, 0xab, 0x7c, 0x86, 0x4e, 0x3f, 0x55,
	0xa0, 0x6d, 0x36, 0xd9, 0xb6, 0xf8, 0x15, 0xb4, 0x4b, 0x9e, 0x61, 0xb7, 0x70, 0x7b, 0x5d, 0xd0,
	0xfa, 0xdf, 0x3f, 0x51, 0x91, 0xdf, 0x77, 0xf7, 0xf5, 0xc7, 0xcf, 0xef, 0x2a, 0x5f, 0x7b, 0x9d,
	0x21, 0xa7, 0x0f, 0x45, 0x50, 0xce, 0xd1, 0x31, 0x56, 0xb0, 0xb9, 0xa0, 0x10, 0xf6, 0x16, 0xc9,
	0x56, 0x09, 0xde, 0x3f, 0x7c, 0xb2, 0xa6, 0x90, 0xd8, 0xb4, 0xfc, 0xc6, 0xeb, 0x0e, 0xc3, 0x6c,
	0xbd, 0xdc, 0x74, 0x0c, 0x30, 0x17, 0x17, 0x1f, 0x2c, 0xb2, 0x2d, 0x39, 0xd2, 0x77, 0xd7, 0x17,
	0xe4, 0xbd, 0x76, 0x4c, 0xaf, 0x9e, 0xd7, 0x1e, 0xc6, 0x4c, 0x69, 0x6b, 0xcc, 0x39, 0x3a, 0xfe,
	0xe3, 0xe8, 0x85, 0x47, 0x92, 0x90, 0x70, 0x1a, 0x26, 0x53, 0xa9, 0xc5, 0x30, 0xe6, 0x76, 0xed,
	0x17, 0xfb, 0xe5, 0x1f, 0xc6, 0x24, 0x91, 0xe1, 0x6d, 0xc3, 0xfc, 0x99, 0xce, 0xbe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x80, 0xdb, 0x31, 0x3d, 0xcf, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AssetServerClient is the client API for AssetServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetServerClient interface {
	NewContract(ctx context.Context, in *ServerNewContractRequest, opts ...grpc.CallOption) (*ServerNewContractResponse, error)
	// CloseContract is used to close a contract with a specific uuid
	CloseContract(ctx context.Context, in *ServerCloseContractRequest, opts ...grpc.CallOption) (*ServerCloseContractResponse, error)
	// ListAssets lists all supported assets
	ListAssets(ctx context.Context, in *ServerListAssetsRequest, opts ...grpc.CallOption) (*ServerListAssetsResponse, error)
}

type assetServerClient struct {
	cc *grpc.ClientConn
}

func NewAssetServerClient(cc *grpc.ClientConn) AssetServerClient {
	return &assetServerClient{cc}
}

func (c *assetServerClient) NewContract(ctx context.Context, in *ServerNewContractRequest, opts ...grpc.CallOption) (*ServerNewContractResponse, error) {
	out := new(ServerNewContractResponse)
	err := c.cc.Invoke(ctx, "/ladrpc.AssetServer/NewContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServerClient) CloseContract(ctx context.Context, in *ServerCloseContractRequest, opts ...grpc.CallOption) (*ServerCloseContractResponse, error) {
	out := new(ServerCloseContractResponse)
	err := c.cc.Invoke(ctx, "/ladrpc.AssetServer/CloseContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServerClient) ListAssets(ctx context.Context, in *ServerListAssetsRequest, opts ...grpc.CallOption) (*ServerListAssetsResponse, error) {
	out := new(ServerListAssetsResponse)
	err := c.cc.Invoke(ctx, "/ladrpc.AssetServer/ListAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServerServer is the server API for AssetServer service.
type AssetServerServer interface {
	NewContract(context.Context, *ServerNewContractRequest) (*ServerNewContractResponse, error)
	// CloseContract is used to close a contract with a specific uuid
	CloseContract(context.Context, *ServerCloseContractRequest) (*ServerCloseContractResponse, error)
	// ListAssets lists all supported assets
	ListAssets(context.Context, *ServerListAssetsRequest) (*ServerListAssetsResponse, error)
}

// UnimplementedAssetServerServer can be embedded to have forward compatible implementations.
type UnimplementedAssetServerServer struct {
}

func (*UnimplementedAssetServerServer) NewContract(ctx context.Context, req *ServerNewContractRequest) (*ServerNewContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewContract not implemented")
}
func (*UnimplementedAssetServerServer) CloseContract(ctx context.Context, req *ServerCloseContractRequest) (*ServerCloseContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseContract not implemented")
}
func (*UnimplementedAssetServerServer) ListAssets(ctx context.Context, req *ServerListAssetsRequest) (*ServerListAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssets not implemented")
}

func RegisterAssetServerServer(s *grpc.Server, srv AssetServerServer) {
	s.RegisterService(&_AssetServer_serviceDesc, srv)
}

func _AssetServer_NewContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerNewContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServerServer).NewContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ladrpc.AssetServer/NewContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServerServer).NewContract(ctx, req.(*ServerNewContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetServer_CloseContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerCloseContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServerServer).CloseContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ladrpc.AssetServer/CloseContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServerServer).CloseContract(ctx, req.(*ServerCloseContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetServer_ListAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerListAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServerServer).ListAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ladrpc.AssetServer/ListAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServerServer).ListAssets(ctx, req.(*ServerListAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssetServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ladrpc.AssetServer",
	HandlerType: (*AssetServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewContract",
			Handler:    _AssetServer_NewContract_Handler,
		},
		{
			MethodName: "CloseContract",
			Handler:    _AssetServer_CloseContract_Handler,
		},
		{
			MethodName: "ListAssets",
			Handler:    _AssetServer_ListAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}