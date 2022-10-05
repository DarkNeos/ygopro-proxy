// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ocgcore.proto

package ygopropb

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

type YgoCtosMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*YgoCtosMsg_CtosPlayerInfo
	//	*YgoCtosMsg_CtosJoinGame
	//	*YgoCtosMsg_CtosUpdateDeck
	//	*YgoCtosMsg_CtosHsReady
	Msg isYgoCtosMsg_Msg `protobuf_oneof:"msg"`
}

func (x *YgoCtosMsg) Reset() {
	*x = YgoCtosMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YgoCtosMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YgoCtosMsg) ProtoMessage() {}

func (x *YgoCtosMsg) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YgoCtosMsg.ProtoReflect.Descriptor instead.
func (*YgoCtosMsg) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{0}
}

func (m *YgoCtosMsg) GetMsg() isYgoCtosMsg_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *YgoCtosMsg) GetCtosPlayerInfo() *CtosPlayerInfo {
	if x, ok := x.GetMsg().(*YgoCtosMsg_CtosPlayerInfo); ok {
		return x.CtosPlayerInfo
	}
	return nil
}

func (x *YgoCtosMsg) GetCtosJoinGame() *CtosJoinGame {
	if x, ok := x.GetMsg().(*YgoCtosMsg_CtosJoinGame); ok {
		return x.CtosJoinGame
	}
	return nil
}

func (x *YgoCtosMsg) GetCtosUpdateDeck() *CtosUpdateDeck {
	if x, ok := x.GetMsg().(*YgoCtosMsg_CtosUpdateDeck); ok {
		return x.CtosUpdateDeck
	}
	return nil
}

func (x *YgoCtosMsg) GetCtosHsReady() *CtosHsReady {
	if x, ok := x.GetMsg().(*YgoCtosMsg_CtosHsReady); ok {
		return x.CtosHsReady
	}
	return nil
}

type isYgoCtosMsg_Msg interface {
	isYgoCtosMsg_Msg()
}

type YgoCtosMsg_CtosPlayerInfo struct {
	CtosPlayerInfo *CtosPlayerInfo `protobuf:"bytes,1,opt,name=ctos_player_info,json=ctosPlayerInfo,proto3,oneof"`
}

type YgoCtosMsg_CtosJoinGame struct {
	CtosJoinGame *CtosJoinGame `protobuf:"bytes,2,opt,name=ctos_join_game,json=ctosJoinGame,proto3,oneof"`
}

type YgoCtosMsg_CtosUpdateDeck struct {
	CtosUpdateDeck *CtosUpdateDeck `protobuf:"bytes,3,opt,name=ctos_update_deck,json=ctosUpdateDeck,proto3,oneof"`
}

type YgoCtosMsg_CtosHsReady struct {
	CtosHsReady *CtosHsReady `protobuf:"bytes,4,opt,name=ctos_hs_ready,json=ctosHsReady,proto3,oneof"`
}

func (*YgoCtosMsg_CtosPlayerInfo) isYgoCtosMsg_Msg() {}

func (*YgoCtosMsg_CtosJoinGame) isYgoCtosMsg_Msg() {}

func (*YgoCtosMsg_CtosUpdateDeck) isYgoCtosMsg_Msg() {}

func (*YgoCtosMsg_CtosHsReady) isYgoCtosMsg_Msg() {}

type YgoStocMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*YgoStocMsg_StocJoinGame
	//	*YgoStocMsg_StocChat
	//	*YgoStocMsg_StocHsPlayerEnter
	//	*YgoStocMsg_StocTypeChange
	Msg isYgoStocMsg_Msg `protobuf_oneof:"msg"`
}

func (x *YgoStocMsg) Reset() {
	*x = YgoStocMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YgoStocMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YgoStocMsg) ProtoMessage() {}

func (x *YgoStocMsg) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YgoStocMsg.ProtoReflect.Descriptor instead.
func (*YgoStocMsg) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{1}
}

func (m *YgoStocMsg) GetMsg() isYgoStocMsg_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *YgoStocMsg) GetStocJoinGame() *StocJoinGame {
	if x, ok := x.GetMsg().(*YgoStocMsg_StocJoinGame); ok {
		return x.StocJoinGame
	}
	return nil
}

func (x *YgoStocMsg) GetStocChat() *StocChat {
	if x, ok := x.GetMsg().(*YgoStocMsg_StocChat); ok {
		return x.StocChat
	}
	return nil
}

func (x *YgoStocMsg) GetStocHsPlayerEnter() *StocHsPlayerEnter {
	if x, ok := x.GetMsg().(*YgoStocMsg_StocHsPlayerEnter); ok {
		return x.StocHsPlayerEnter
	}
	return nil
}

func (x *YgoStocMsg) GetStocTypeChange() *StocTypeChange {
	if x, ok := x.GetMsg().(*YgoStocMsg_StocTypeChange); ok {
		return x.StocTypeChange
	}
	return nil
}

type isYgoStocMsg_Msg interface {
	isYgoStocMsg_Msg()
}

type YgoStocMsg_StocJoinGame struct {
	StocJoinGame *StocJoinGame `protobuf:"bytes,1,opt,name=stoc_join_game,json=stocJoinGame,proto3,oneof"`
}

type YgoStocMsg_StocChat struct {
	StocChat *StocChat `protobuf:"bytes,2,opt,name=stoc_chat,json=stocChat,proto3,oneof"`
}

type YgoStocMsg_StocHsPlayerEnter struct {
	StocHsPlayerEnter *StocHsPlayerEnter `protobuf:"bytes,3,opt,name=stoc_hs_player_enter,json=stocHsPlayerEnter,proto3,oneof"`
}

type YgoStocMsg_StocTypeChange struct {
	StocTypeChange *StocTypeChange `protobuf:"bytes,4,opt,name=stoc_type_change,json=stocTypeChange,proto3,oneof"`
}

func (*YgoStocMsg_StocJoinGame) isYgoStocMsg_Msg() {}

func (*YgoStocMsg_StocChat) isYgoStocMsg_Msg() {}

func (*YgoStocMsg_StocHsPlayerEnter) isYgoStocMsg_Msg() {}

func (*YgoStocMsg_StocTypeChange) isYgoStocMsg_Msg() {}

type CtosPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CtosPlayerInfo) Reset() {
	*x = CtosPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtosPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtosPlayerInfo) ProtoMessage() {}

func (x *CtosPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtosPlayerInfo.ProtoReflect.Descriptor instead.
func (*CtosPlayerInfo) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{2}
}

func (x *CtosPlayerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CtosJoinGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Gameid  int32  `protobuf:"varint,2,opt,name=gameid,proto3" json:"gameid,omitempty"`
	Passwd  string `protobuf:"bytes,3,opt,name=passwd,proto3" json:"passwd,omitempty"`
}

func (x *CtosJoinGame) Reset() {
	*x = CtosJoinGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtosJoinGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtosJoinGame) ProtoMessage() {}

func (x *CtosJoinGame) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtosJoinGame.ProtoReflect.Descriptor instead.
func (*CtosJoinGame) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{3}
}

func (x *CtosJoinGame) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CtosJoinGame) GetGameid() int32 {
	if x != nil {
		return x.Gameid
	}
	return 0
}

func (x *CtosJoinGame) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

type CtosUpdateDeck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Main  []int32 `protobuf:"varint,1,rep,packed,name=main,proto3" json:"main,omitempty"`
	Extra []int32 `protobuf:"varint,2,rep,packed,name=extra,proto3" json:"extra,omitempty"`
	Side  []int32 `protobuf:"varint,3,rep,packed,name=side,proto3" json:"side,omitempty"`
}

func (x *CtosUpdateDeck) Reset() {
	*x = CtosUpdateDeck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtosUpdateDeck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtosUpdateDeck) ProtoMessage() {}

func (x *CtosUpdateDeck) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtosUpdateDeck.ProtoReflect.Descriptor instead.
func (*CtosUpdateDeck) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{4}
}

func (x *CtosUpdateDeck) GetMain() []int32 {
	if x != nil {
		return x.Main
	}
	return nil
}

func (x *CtosUpdateDeck) GetExtra() []int32 {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *CtosUpdateDeck) GetSide() []int32 {
	if x != nil {
		return x.Side
	}
	return nil
}

type CtosHsReady struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CtosHsReady) Reset() {
	*x = CtosHsReady{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtosHsReady) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtosHsReady) ProtoMessage() {}

func (x *CtosHsReady) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtosHsReady.ProtoReflect.Descriptor instead.
func (*CtosHsReady) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{5}
}

type StocJoinGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lflist        int32 `protobuf:"varint,1,opt,name=lflist,proto3" json:"lflist,omitempty"`
	Rule          int32 `protobuf:"varint,2,opt,name=rule,proto3" json:"rule,omitempty"`
	Mode          int32 `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	DuelRule      int32 `protobuf:"varint,4,opt,name=duel_rule,json=duelRule,proto3" json:"duel_rule,omitempty"`
	NoCheckDeck   bool  `protobuf:"varint,5,opt,name=no_check_deck,json=noCheckDeck,proto3" json:"no_check_deck,omitempty"`
	NoShuffleDeck bool  `protobuf:"varint,6,opt,name=no_shuffle_deck,json=noShuffleDeck,proto3" json:"no_shuffle_deck,omitempty"`
	StartLp       int32 `protobuf:"varint,7,opt,name=start_lp,json=startLp,proto3" json:"start_lp,omitempty"`
	StartHand     int32 `protobuf:"varint,8,opt,name=start_hand,json=startHand,proto3" json:"start_hand,omitempty"`
	DrawCount     int32 `protobuf:"varint,9,opt,name=draw_count,json=drawCount,proto3" json:"draw_count,omitempty"`
	TimeLimit     int32 `protobuf:"varint,10,opt,name=time_limit,json=timeLimit,proto3" json:"time_limit,omitempty"`
}

func (x *StocJoinGame) Reset() {
	*x = StocJoinGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocJoinGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocJoinGame) ProtoMessage() {}

func (x *StocJoinGame) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocJoinGame.ProtoReflect.Descriptor instead.
func (*StocJoinGame) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{6}
}

func (x *StocJoinGame) GetLflist() int32 {
	if x != nil {
		return x.Lflist
	}
	return 0
}

func (x *StocJoinGame) GetRule() int32 {
	if x != nil {
		return x.Rule
	}
	return 0
}

func (x *StocJoinGame) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *StocJoinGame) GetDuelRule() int32 {
	if x != nil {
		return x.DuelRule
	}
	return 0
}

func (x *StocJoinGame) GetNoCheckDeck() bool {
	if x != nil {
		return x.NoCheckDeck
	}
	return false
}

func (x *StocJoinGame) GetNoShuffleDeck() bool {
	if x != nil {
		return x.NoShuffleDeck
	}
	return false
}

func (x *StocJoinGame) GetStartLp() int32 {
	if x != nil {
		return x.StartLp
	}
	return 0
}

func (x *StocJoinGame) GetStartHand() int32 {
	if x != nil {
		return x.StartHand
	}
	return 0
}

func (x *StocJoinGame) GetDrawCount() int32 {
	if x != nil {
		return x.DrawCount
	}
	return 0
}

func (x *StocJoinGame) GetTimeLimit() int32 {
	if x != nil {
		return x.TimeLimit
	}
	return 0
}

type StocChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player int32  `protobuf:"varint,1,opt,name=player,proto3" json:"player,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *StocChat) Reset() {
	*x = StocChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocChat) ProtoMessage() {}

func (x *StocChat) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocChat.ProtoReflect.Descriptor instead.
func (*StocChat) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{7}
}

func (x *StocChat) GetPlayer() int32 {
	if x != nil {
		return x.Player
	}
	return 0
}

func (x *StocChat) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type StocHsPlayerEnter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pos  int32  `protobuf:"varint,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *StocHsPlayerEnter) Reset() {
	*x = StocHsPlayerEnter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocHsPlayerEnter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocHsPlayerEnter) ProtoMessage() {}

func (x *StocHsPlayerEnter) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocHsPlayerEnter.ProtoReflect.Descriptor instead.
func (*StocHsPlayerEnter) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{8}
}

func (x *StocHsPlayerEnter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StocHsPlayerEnter) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

type StocTypeChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *StocTypeChange) Reset() {
	*x = StocTypeChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocgcore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocTypeChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocTypeChange) ProtoMessage() {}

func (x *StocTypeChange) ProtoReflect() protoreflect.Message {
	mi := &file_ocgcore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocTypeChange.ProtoReflect.Descriptor instead.
func (*StocTypeChange) Descriptor() ([]byte, []int) {
	return file_ocgcore_proto_rawDescGZIP(), []int{9}
}

func (x *StocTypeChange) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_ocgcore_proto protoreflect.FileDescriptor

var file_ocgcore_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x63, 0x67, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x0a, 0x59, 0x67, 0x6f, 0x43,
	0x74, 0x6f, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x42, 0x0a, 0x10, 0x63, 0x74, 0x6f, 0x73, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x2e, 0x43, 0x74, 0x6f, 0x73, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x74, 0x6f, 0x73,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x0e, 0x63, 0x74,
	0x6f, 0x73, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x2e, 0x43, 0x74, 0x6f, 0x73,
	0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x74, 0x6f, 0x73,
	0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x63, 0x74, 0x6f, 0x73,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x2e, 0x43, 0x74, 0x6f, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x74,
	0x6f, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x39, 0x0a, 0x0d,
	0x63, 0x74, 0x6f, 0x73, 0x5f, 0x68, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x2e, 0x43, 0x74, 0x6f,
	0x73, 0x48, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x74, 0x6f, 0x73,
	0x48, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x94,
	0x02, 0x0a, 0x0a, 0x59, 0x67, 0x6f, 0x53, 0x74, 0x6f, 0x63, 0x4d, 0x73, 0x67, 0x12, 0x3c, 0x0a,
	0x0e, 0x73, 0x74, 0x6f, 0x63, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x73,
	0x74, 0x6f, 0x63, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x73,
	0x74, 0x6f, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x43, 0x68, 0x61, 0x74,
	0x48, 0x00, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x43, 0x68, 0x61, 0x74, 0x12, 0x4c, 0x0a, 0x14,
	0x73, 0x74, 0x6f, 0x63, 0x5f, 0x68, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x79, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x48, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x63, 0x48, 0x73, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x10, 0x73, 0x74,
	0x6f, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x2e, 0x53, 0x74,
	0x6f, 0x63, 0x54, 0x79, 0x70, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0e,
	0x73, 0x74, 0x6f, 0x63, 0x54, 0x79, 0x70, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x05,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x74, 0x6f, 0x73, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x0c, 0x43,
	0x74, 0x6f, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x64, 0x22, 0x4e, 0x0a, 0x0e, 0x43, 0x74, 0x6f, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x64, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x43, 0x74, 0x6f, 0x73, 0x48, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x22, 0xaf, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x4a, 0x6f, 0x69,
	0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x66, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x66, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x75, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x75, 0x65, 0x6c, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x65, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x6f, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x64,
	0x65, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6e, 0x6f, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x6f, 0x5f, 0x73, 0x68, 0x75,
	0x66, 0x66, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x6e, 0x6f, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x72, 0x61, 0x77,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x72,
	0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x34, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x63, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x39, 0x0a, 0x11,
	0x53, 0x74, 0x6f, 0x63, 0x48, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x13, 0x5a,
	0x11, 0x44, 0x61, 0x72, 0x6b, 0x4e, 0x65, 0x6f, 0x73, 0x2f, 0x79, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocgcore_proto_rawDescOnce sync.Once
	file_ocgcore_proto_rawDescData = file_ocgcore_proto_rawDesc
)

func file_ocgcore_proto_rawDescGZIP() []byte {
	file_ocgcore_proto_rawDescOnce.Do(func() {
		file_ocgcore_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocgcore_proto_rawDescData)
	})
	return file_ocgcore_proto_rawDescData
}

var file_ocgcore_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ocgcore_proto_goTypes = []interface{}{
	(*YgoCtosMsg)(nil),        // 0: ygopro.YgoCtosMsg
	(*YgoStocMsg)(nil),        // 1: ygopro.YgoStocMsg
	(*CtosPlayerInfo)(nil),    // 2: ygopro.CtosPlayerInfo
	(*CtosJoinGame)(nil),      // 3: ygopro.CtosJoinGame
	(*CtosUpdateDeck)(nil),    // 4: ygopro.CtosUpdateDeck
	(*CtosHsReady)(nil),       // 5: ygopro.CtosHsReady
	(*StocJoinGame)(nil),      // 6: ygopro.StocJoinGame
	(*StocChat)(nil),          // 7: ygopro.StocChat
	(*StocHsPlayerEnter)(nil), // 8: ygopro.StocHsPlayerEnter
	(*StocTypeChange)(nil),    // 9: ygopro.StocTypeChange
}
var file_ocgcore_proto_depIdxs = []int32{
	2, // 0: ygopro.YgoCtosMsg.ctos_player_info:type_name -> ygopro.CtosPlayerInfo
	3, // 1: ygopro.YgoCtosMsg.ctos_join_game:type_name -> ygopro.CtosJoinGame
	4, // 2: ygopro.YgoCtosMsg.ctos_update_deck:type_name -> ygopro.CtosUpdateDeck
	5, // 3: ygopro.YgoCtosMsg.ctos_hs_ready:type_name -> ygopro.CtosHsReady
	6, // 4: ygopro.YgoStocMsg.stoc_join_game:type_name -> ygopro.StocJoinGame
	7, // 5: ygopro.YgoStocMsg.stoc_chat:type_name -> ygopro.StocChat
	8, // 6: ygopro.YgoStocMsg.stoc_hs_player_enter:type_name -> ygopro.StocHsPlayerEnter
	9, // 7: ygopro.YgoStocMsg.stoc_type_change:type_name -> ygopro.StocTypeChange
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_ocgcore_proto_init() }
func file_ocgcore_proto_init() {
	if File_ocgcore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocgcore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YgoCtosMsg); i {
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
		file_ocgcore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YgoStocMsg); i {
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
		file_ocgcore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtosPlayerInfo); i {
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
		file_ocgcore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtosJoinGame); i {
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
		file_ocgcore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtosUpdateDeck); i {
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
		file_ocgcore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtosHsReady); i {
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
		file_ocgcore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocJoinGame); i {
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
		file_ocgcore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocChat); i {
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
		file_ocgcore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocHsPlayerEnter); i {
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
		file_ocgcore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocTypeChange); i {
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
	file_ocgcore_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*YgoCtosMsg_CtosPlayerInfo)(nil),
		(*YgoCtosMsg_CtosJoinGame)(nil),
		(*YgoCtosMsg_CtosUpdateDeck)(nil),
		(*YgoCtosMsg_CtosHsReady)(nil),
	}
	file_ocgcore_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*YgoStocMsg_StocJoinGame)(nil),
		(*YgoStocMsg_StocChat)(nil),
		(*YgoStocMsg_StocHsPlayerEnter)(nil),
		(*YgoStocMsg_StocTypeChange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ocgcore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ocgcore_proto_goTypes,
		DependencyIndexes: file_ocgcore_proto_depIdxs,
		MessageInfos:      file_ocgcore_proto_msgTypes,
	}.Build()
	File_ocgcore_proto = out.File
	file_ocgcore_proto_rawDesc = nil
	file_ocgcore_proto_goTypes = nil
	file_ocgcore_proto_depIdxs = nil
}
