// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.3
// source: message.proto

package proto_msg

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

type UserBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"UserID"
	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// @inject_tag: db:"GameID"
	GameID int32 `protobuf:"varint,2,opt,name=gameID,proto3" json:"gameID,omitempty"`
	// @inject_tag: db:"Accounts"
	Accounts string `protobuf:"bytes,3,opt,name=accounts,proto3" json:"accounts,omitempty"`
	// @inject_tag: db:"NickName"
	NickName string `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	// @inject_tag: db:"FaceID"
	FaceID string `protobuf:"bytes,5,opt,name=faceID,proto3" json:"faceID,omitempty"`
	// @inject_tag: db:"Gender"
	Gender string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	// @inject_tag: db:"Nullity"
	Nullity bool `protobuf:"varint,7,opt,name=nullity,proto3" json:"nullity,omitempty"`
	// @inject_tag: db:"LastLogonDate"
	LastLogonDate int64 `protobuf:"varint,8,opt,name=lastLogonDate,proto3" json:"lastLogonDate,omitempty"`
}

func (x *UserBasic) Reset() {
	*x = UserBasic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBasic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBasic) ProtoMessage() {}

func (x *UserBasic) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBasic.ProtoReflect.Descriptor instead.
func (*UserBasic) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *UserBasic) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserBasic) GetGameID() int32 {
	if x != nil {
		return x.GameID
	}
	return 0
}

func (x *UserBasic) GetAccounts() string {
	if x != nil {
		return x.Accounts
	}
	return ""
}

func (x *UserBasic) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserBasic) GetFaceID() string {
	if x != nil {
		return x.FaceID
	}
	return ""
}

func (x *UserBasic) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserBasic) GetNullity() bool {
	if x != nil {
		return x.Nullity
	}
	return false
}

func (x *UserBasic) GetLastLogonDate() int64 {
	if x != nil {
		return x.LastLogonDate
	}
	return 0
}

type UserBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"UserID"
	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// @inject_tag: db:"Score"
	Score int64 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	// @inject_tag: db:"Revenue"
	Revenue int64 `protobuf:"varint,3,opt,name=revenue,proto3" json:"revenue,omitempty"`
	// @inject_tag: db:"InsureScore"
	InsureScore int64 `protobuf:"varint,4,opt,name=insureScore,proto3" json:"insureScore,omitempty"`
	// @inject_tag: db:"AllWinScore"
	AllWinScore int64 `protobuf:"varint,5,opt,name=allWinScore,proto3" json:"allWinScore,omitempty"`
	// @inject_tag: db:"BetTotal"
	BetTotal int64 `protobuf:"varint,6,opt,name=betTotal,proto3" json:"betTotal,omitempty"`
	// @inject_tag: db:"BetCount"
	BetCount int64 `protobuf:"varint,7,opt,name=betCount,proto3" json:"betCount,omitempty"`
	// @inject_tag: db:"EffectiveBetTotal"
	EffectiveBetTotal int64 `protobuf:"varint,8,opt,name=effectiveBetTotal,proto3" json:"effectiveBetTotal,omitempty"`
	// @inject_tag: db:"JackpotScore"
	JackpotScore int64 `protobuf:"varint,9,opt,name=jackpotScore,proto3" json:"jackpotScore,omitempty"`
	// @inject_tag: db:"TMScore"
	TmScore int64 `protobuf:"varint,10,opt,name=tmScore,proto3" json:"tmScore,omitempty"`
	// @inject_tag: db:"JackpotBet"
	JackpotBet int64 `protobuf:"varint,11,opt,name=jackpotBet,proto3" json:"jackpotBet,omitempty"`
}

func (x *UserBalance) Reset() {
	*x = UserBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBalance) ProtoMessage() {}

func (x *UserBalance) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBalance.ProtoReflect.Descriptor instead.
func (*UserBalance) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *UserBalance) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserBalance) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *UserBalance) GetRevenue() int64 {
	if x != nil {
		return x.Revenue
	}
	return 0
}

func (x *UserBalance) GetInsureScore() int64 {
	if x != nil {
		return x.InsureScore
	}
	return 0
}

func (x *UserBalance) GetAllWinScore() int64 {
	if x != nil {
		return x.AllWinScore
	}
	return 0
}

func (x *UserBalance) GetBetTotal() int64 {
	if x != nil {
		return x.BetTotal
	}
	return 0
}

func (x *UserBalance) GetBetCount() int64 {
	if x != nil {
		return x.BetCount
	}
	return 0
}

func (x *UserBalance) GetEffectiveBetTotal() int64 {
	if x != nil {
		return x.EffectiveBetTotal
	}
	return 0
}

func (x *UserBalance) GetJackpotScore() int64 {
	if x != nil {
		return x.JackpotScore
	}
	return 0
}

func (x *UserBalance) GetTmScore() int64 {
	if x != nil {
		return x.TmScore
	}
	return 0
}

func (x *UserBalance) GetJackpotBet() int64 {
	if x != nil {
		return x.JackpotBet
	}
	return 0
}

type UserStartGameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"UserID"
	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// @inject_tag: db:"NickName"
	NickName string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	// @inject_tag: db:"Score"
	Score int64 `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	// @inject_tag: db:"FaceID"
	FaceID string `protobuf:"bytes,4,opt,name=faceID,proto3" json:"faceID,omitempty"`
	// @inject_tag: db:"Gender"
	Gender string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	// player's 0-based seat position
	SeatPosition int32 `protobuf:"varint,6,opt,name=seatPosition,proto3" json:"seatPosition,omitempty"`
}

func (x *UserStartGameInfo) Reset() {
	*x = UserStartGameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStartGameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStartGameInfo) ProtoMessage() {}

func (x *UserStartGameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStartGameInfo.ProtoReflect.Descriptor instead.
func (*UserStartGameInfo) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *UserStartGameInfo) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserStartGameInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserStartGameInfo) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *UserStartGameInfo) GetFaceID() string {
	if x != nil {
		return x.FaceID
	}
	return ""
}

func (x *UserStartGameInfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserStartGameInfo) GetSeatPosition() int32 {
	if x != nil {
		return x.SeatPosition
	}
	return 0
}

type UserToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"UserID"
	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// @inject_tag: db:"DynamicPass"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserToken) Reset() {
	*x = UserToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserToken) ProtoMessage() {}

func (x *UserToken) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserToken.ProtoReflect.Descriptor instead.
func (*UserToken) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *UserToken) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// MQ message
type UserMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*UserMsg_UserBasic
	//	*UserMsg_UserBalance
	//	*UserMsg_UserStartGameInfo
	//	*UserMsg_UserToken
	Msg isUserMsg_Msg `protobuf_oneof:"msg"`
}

func (x *UserMsg) Reset() {
	*x = UserMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMsg) ProtoMessage() {}

func (x *UserMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMsg.ProtoReflect.Descriptor instead.
func (*UserMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (m *UserMsg) GetMsg() isUserMsg_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *UserMsg) GetUserBasic() *UserBasic {
	if x, ok := x.GetMsg().(*UserMsg_UserBasic); ok {
		return x.UserBasic
	}
	return nil
}

func (x *UserMsg) GetUserBalance() *UserBalance {
	if x, ok := x.GetMsg().(*UserMsg_UserBalance); ok {
		return x.UserBalance
	}
	return nil
}

func (x *UserMsg) GetUserStartGameInfo() *UserStartGameInfo {
	if x, ok := x.GetMsg().(*UserMsg_UserStartGameInfo); ok {
		return x.UserStartGameInfo
	}
	return nil
}

func (x *UserMsg) GetUserToken() *UserToken {
	if x, ok := x.GetMsg().(*UserMsg_UserToken); ok {
		return x.UserToken
	}
	return nil
}

type isUserMsg_Msg interface {
	isUserMsg_Msg()
}

type UserMsg_UserBasic struct {
	UserBasic *UserBasic `protobuf:"bytes,1,opt,name=userBasic,proto3,oneof"`
}

type UserMsg_UserBalance struct {
	UserBalance *UserBalance `protobuf:"bytes,2,opt,name=userBalance,proto3,oneof"`
}

type UserMsg_UserStartGameInfo struct {
	UserStartGameInfo *UserStartGameInfo `protobuf:"bytes,3,opt,name=userStartGameInfo,proto3,oneof"`
}

type UserMsg_UserToken struct {
	UserToken *UserToken `protobuf:"bytes,4,opt,name=userToken,proto3,oneof"`
}

func (*UserMsg_UserBasic) isUserMsg_Msg() {}

func (*UserMsg_UserBalance) isUserMsg_Msg() {}

func (*UserMsg_UserStartGameInfo) isUserMsg_Msg() {}

func (*UserMsg_UserToken) isUserMsg_Msg() {}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0xe3, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x63, 0x65,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75,
	0x6c, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x75, 0x6c,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0xdd, 0x02, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65,
	0x6e, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x6e,
	0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x57, 0x69, 0x6e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x57, 0x69,
	0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x74, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x65, 0x74, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c,
	0x0a, 0x11, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x65, 0x74, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x42, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c,
	0x6a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x6a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x74, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x61,
	0x63, 0x6b, 0x70, 0x6f, 0x74, 0x42, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x6a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x42, 0x65, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x63, 0x65,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65,
	0x61, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x73, 0x65, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x39,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xee, 0x01, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x75,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x11, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x37, 0x5a, 0x35, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74,
	0x6d, 0x71, 0x2f, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x2d, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6d, 0x73, 0x67, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(*UserBasic)(nil),         // 0: msg.UserBasic
	(*UserBalance)(nil),       // 1: msg.UserBalance
	(*UserStartGameInfo)(nil), // 2: msg.UserStartGameInfo
	(*UserToken)(nil),         // 3: msg.UserToken
	(*UserMsg)(nil),           // 4: msg.UserMsg
}
var file_message_proto_depIdxs = []int32{
	0, // 0: msg.UserMsg.userBasic:type_name -> msg.UserBasic
	1, // 1: msg.UserMsg.userBalance:type_name -> msg.UserBalance
	2, // 2: msg.UserMsg.userStartGameInfo:type_name -> msg.UserStartGameInfo
	3, // 3: msg.UserMsg.userToken:type_name -> msg.UserToken
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBasic); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBalance); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStartGameInfo); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserToken); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMsg); i {
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
	file_message_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UserMsg_UserBasic)(nil),
		(*UserMsg_UserBalance)(nil),
		(*UserMsg_UserStartGameInfo)(nil),
		(*UserMsg_UserToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}