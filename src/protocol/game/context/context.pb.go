// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/game/context.proto

package context

import (
	common "go-xp-dofus-unity-proto-builder/src/protocol/game/common"
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

type ContextCreationEvent_GameContext int32

const (
	ContextCreationEvent_ROLE_PLAY ContextCreationEvent_GameContext = 0
	ContextCreationEvent_FIGHT     ContextCreationEvent_GameContext = 1
)

// Enum value maps for ContextCreationEvent_GameContext.
var (
	ContextCreationEvent_GameContext_name = map[int32]string{
		0: "ROLE_PLAY",
		1: "FIGHT",
	}
	ContextCreationEvent_GameContext_value = map[string]int32{
		"ROLE_PLAY": 0,
		"FIGHT":     1,
	}
)

func (x ContextCreationEvent_GameContext) Enum() *ContextCreationEvent_GameContext {
	p := new(ContextCreationEvent_GameContext)
	*p = x
	return p
}

func (x ContextCreationEvent_GameContext) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContextCreationEvent_GameContext) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_game_context_proto_enumTypes[0].Descriptor()
}

func (ContextCreationEvent_GameContext) Type() protoreflect.EnumType {
	return &file_proto_game_context_proto_enumTypes[0]
}

func (x ContextCreationEvent_GameContext) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContextCreationEvent_GameContext.Descriptor instead.
func (ContextCreationEvent_GameContext) EnumDescriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{5, 0}
}

type ContextCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContextCreationRequest) Reset() {
	*x = ContextCreationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextCreationRequest) ProtoMessage() {}

func (x *ContextCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextCreationRequest.ProtoReflect.Descriptor instead.
func (*ContextCreationRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{0}
}

type ContextReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapId int64 `protobuf:"varint,1,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
}

func (x *ContextReadyRequest) Reset() {
	*x = ContextReadyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextReadyRequest) ProtoMessage() {}

func (x *ContextReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextReadyRequest.ProtoReflect.Descriptor instead.
func (*ContextReadyRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{1}
}

func (x *ContextReadyRequest) GetMapId() int64 {
	if x != nil {
		return x.MapId
	}
	return 0
}

type ContextQuitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContextQuitRequest) Reset() {
	*x = ContextQuitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextQuitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextQuitRequest) ProtoMessage() {}

func (x *ContextQuitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextQuitRequest.ProtoReflect.Descriptor instead.
func (*ContextQuitRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{2}
}

type ContextKickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId int64 `protobuf:"varint,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
}

func (x *ContextKickRequest) Reset() {
	*x = ContextKickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextKickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextKickRequest) ProtoMessage() {}

func (x *ContextKickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextKickRequest.ProtoReflect.Descriptor instead.
func (*ContextKickRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{3}
}

func (x *ContextKickRequest) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

type ShowCellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CellId int32 `protobuf:"varint,1,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
}

func (x *ShowCellRequest) Reset() {
	*x = ShowCellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowCellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowCellRequest) ProtoMessage() {}

func (x *ShowCellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowCellRequest.ProtoReflect.Descriptor instead.
func (*ShowCellRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{4}
}

func (x *ShowCellRequest) GetCellId() int32 {
	if x != nil {
		return x.CellId
	}
	return 0
}

type ContextCreationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context ContextCreationEvent_GameContext `protobuf:"varint,1,opt,name=context,proto3,enum=com.ankama.dofus.server.game.protocol.context.ContextCreationEvent_GameContext" json:"context,omitempty"`
}

func (x *ContextCreationEvent) Reset() {
	*x = ContextCreationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextCreationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextCreationEvent) ProtoMessage() {}

func (x *ContextCreationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextCreationEvent.ProtoReflect.Descriptor instead.
func (*ContextCreationEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{5}
}

func (x *ContextCreationEvent) GetContext() ContextCreationEvent_GameContext {
	if x != nil {
		return x.Context
	}
	return ContextCreationEvent_ROLE_PLAY
}

type ContextDestroyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContextDestroyEvent) Reset() {
	*x = ContextDestroyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextDestroyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextDestroyEvent) ProtoMessage() {}

func (x *ContextDestroyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextDestroyEvent.ProtoReflect.Descriptor instead.
func (*ContextDestroyEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{6}
}

type ContextRemoveElementEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElementId int64 `protobuf:"varint,1,opt,name=element_id,json=elementId,proto3" json:"element_id,omitempty"`
}

func (x *ContextRemoveElementEvent) Reset() {
	*x = ContextRemoveElementEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextRemoveElementEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextRemoveElementEvent) ProtoMessage() {}

func (x *ContextRemoveElementEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextRemoveElementEvent.ProtoReflect.Descriptor instead.
func (*ContextRemoveElementEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{7}
}

func (x *ContextRemoveElementEvent) GetElementId() int64 {
	if x != nil {
		return x.ElementId
	}
	return 0
}

type ContextRemoveElementsEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElementId []int64 `protobuf:"varint,1,rep,packed,name=element_id,json=elementId,proto3" json:"element_id,omitempty"`
}

func (x *ContextRemoveElementsEvent) Reset() {
	*x = ContextRemoveElementsEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextRemoveElementsEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextRemoveElementsEvent) ProtoMessage() {}

func (x *ContextRemoveElementsEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextRemoveElementsEvent.ProtoReflect.Descriptor instead.
func (*ContextRemoveElementsEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{8}
}

func (x *ContextRemoveElementsEvent) GetElementId() []int64 {
	if x != nil {
		return x.ElementId
	}
	return nil
}

type ContextRefreshEntityLookEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId int64              `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Look     *common.EntityLook `protobuf:"bytes,2,opt,name=look,proto3" json:"look,omitempty"`
}

func (x *ContextRefreshEntityLookEvent) Reset() {
	*x = ContextRefreshEntityLookEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextRefreshEntityLookEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextRefreshEntityLookEvent) ProtoMessage() {}

func (x *ContextRefreshEntityLookEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextRefreshEntityLookEvent.ProtoReflect.Descriptor instead.
func (*ContextRefreshEntityLookEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{9}
}

func (x *ContextRefreshEntityLookEvent) GetEntityId() int64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *ContextRefreshEntityLookEvent) GetLook() *common.EntityLook {
	if x != nil {
		return x.Look
	}
	return nil
}

type EntitiesDispositionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dispositions []*common.EntityDisposition `protobuf:"bytes,1,rep,name=dispositions,proto3" json:"dispositions,omitempty"`
}

func (x *EntitiesDispositionEvent) Reset() {
	*x = EntitiesDispositionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntitiesDispositionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitiesDispositionEvent) ProtoMessage() {}

func (x *EntitiesDispositionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitiesDispositionEvent.ProtoReflect.Descriptor instead.
func (*EntitiesDispositionEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{10}
}

func (x *EntitiesDispositionEvent) GetDispositions() []*common.EntityDisposition {
	if x != nil {
		return x.Dispositions
	}
	return nil
}

type EntityDispositionErrorEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EntityDispositionErrorEvent) Reset() {
	*x = EntityDispositionErrorEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityDispositionErrorEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityDispositionErrorEvent) ProtoMessage() {}

func (x *EntityDispositionErrorEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityDispositionErrorEvent.ProtoReflect.Descriptor instead.
func (*EntityDispositionErrorEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{11}
}

type MonsterBoosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XpBoost   int32 `protobuf:"varint,2,opt,name=xp_boost,json=xpBoost,proto3" json:"xp_boost,omitempty"`
	DropBoost int32 `protobuf:"varint,3,opt,name=drop_boost,json=dropBoost,proto3" json:"drop_boost,omitempty"`
}

func (x *MonsterBoosts) Reset() {
	*x = MonsterBoosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonsterBoosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonsterBoosts) ProtoMessage() {}

func (x *MonsterBoosts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonsterBoosts.ProtoReflect.Descriptor instead.
func (*MonsterBoosts) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{12}
}

func (x *MonsterBoosts) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MonsterBoosts) GetXpBoost() int32 {
	if x != nil {
		return x.XpBoost
	}
	return 0
}

func (x *MonsterBoosts) GetDropBoost() int32 {
	if x != nil {
		return x.DropBoost
	}
	return 0
}

type RefreshMonsterBoostsEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonsterBoosts []*MonsterBoosts `protobuf:"bytes,1,rep,name=monster_boosts,json=monsterBoosts,proto3" json:"monster_boosts,omitempty"`
	FamilyBoosts  []*MonsterBoosts `protobuf:"bytes,2,rep,name=family_boosts,json=familyBoosts,proto3" json:"family_boosts,omitempty"`
}

func (x *RefreshMonsterBoostsEvent) Reset() {
	*x = RefreshMonsterBoostsEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshMonsterBoostsEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshMonsterBoostsEvent) ProtoMessage() {}

func (x *RefreshMonsterBoostsEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshMonsterBoostsEvent.ProtoReflect.Descriptor instead.
func (*RefreshMonsterBoostsEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{13}
}

func (x *RefreshMonsterBoostsEvent) GetMonsterBoosts() []*MonsterBoosts {
	if x != nil {
		return x.MonsterBoosts
	}
	return nil
}

func (x *RefreshMonsterBoostsEvent) GetFamilyBoosts() []*MonsterBoosts {
	if x != nil {
		return x.FamilyBoosts
	}
	return nil
}

type ShowCellEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId int64 `protobuf:"varint,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	CellId   int32 `protobuf:"varint,2,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
}

func (x *ShowCellEvent) Reset() {
	*x = ShowCellEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_context_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowCellEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowCellEvent) ProtoMessage() {}

func (x *ShowCellEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_context_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowCellEvent.ProtoReflect.Descriptor instead.
func (*ShowCellEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_context_proto_rawDescGZIP(), []int{14}
}

func (x *ShowCellEvent) GetSourceId() int64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *ShowCellEvent) GetCellId() int32 {
	if x != nil {
		return x.CellId
	}
	return 0
}

var File_proto_game_context_proto protoreflect.FileDescriptor

var file_proto_game_context_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x6e, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x6f, 0x66, 0x75, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x2c, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x4b, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x77,
	0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x65,
	0x6c, 0x6c, 0x49, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x69, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4f,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x6f, 0x66, 0x75,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x27, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x50, 0x4c, 0x41, 0x59, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x47, 0x48, 0x54, 0x10,
	0x01, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65, 0x73, 0x74,
	0x72, 0x6f, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x8a, 0x01, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x6f, 0x6b, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x4c, 0x0a, 0x04, 0x6c, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x6f, 0x66, 0x75,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x6c, 0x6f, 0x6f, 0x6b, 0x22, 0x7f,
	0x0a, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x6f,
	0x66, 0x75, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x1d, 0x0a, 0x1b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x59,
	0x0a, 0x0d, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x78, 0x70, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x78, 0x70, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x72,
	0x6f, 0x70, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x64, 0x72, 0x6f, 0x70, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x19, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73,
	0x74, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x0e, 0x6d, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x6f, 0x66,
	0x75, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x0d, 0x6d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x61, 0x0a, 0x0d,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6b, 0x61, 0x6d, 0x61,
	0x2e, 0x64, 0x6f, 0x66, 0x75, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x0c, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x73, 0x22,
	0x45, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x77, 0x43, 0x65, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x63, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x6f, 0x2d, 0x78, 0x70, 0x2d,
	0x64, 0x6f, 0x66, 0x75, 0x73, 0x2d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_game_context_proto_rawDescOnce sync.Once
	file_proto_game_context_proto_rawDescData = file_proto_game_context_proto_rawDesc
)

func file_proto_game_context_proto_rawDescGZIP() []byte {
	file_proto_game_context_proto_rawDescOnce.Do(func() {
		file_proto_game_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_game_context_proto_rawDescData)
	})
	return file_proto_game_context_proto_rawDescData
}

var file_proto_game_context_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_game_context_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_game_context_proto_goTypes = []any{
	(ContextCreationEvent_GameContext)(0), // 0: com.ankama.dofus.server.game.protocol.context.ContextCreationEvent.GameContext
	(*ContextCreationRequest)(nil),        // 1: com.ankama.dofus.server.game.protocol.context.ContextCreationRequest
	(*ContextReadyRequest)(nil),           // 2: com.ankama.dofus.server.game.protocol.context.ContextReadyRequest
	(*ContextQuitRequest)(nil),            // 3: com.ankama.dofus.server.game.protocol.context.ContextQuitRequest
	(*ContextKickRequest)(nil),            // 4: com.ankama.dofus.server.game.protocol.context.ContextKickRequest
	(*ShowCellRequest)(nil),               // 5: com.ankama.dofus.server.game.protocol.context.ShowCellRequest
	(*ContextCreationEvent)(nil),          // 6: com.ankama.dofus.server.game.protocol.context.ContextCreationEvent
	(*ContextDestroyEvent)(nil),           // 7: com.ankama.dofus.server.game.protocol.context.ContextDestroyEvent
	(*ContextRemoveElementEvent)(nil),     // 8: com.ankama.dofus.server.game.protocol.context.ContextRemoveElementEvent
	(*ContextRemoveElementsEvent)(nil),    // 9: com.ankama.dofus.server.game.protocol.context.ContextRemoveElementsEvent
	(*ContextRefreshEntityLookEvent)(nil), // 10: com.ankama.dofus.server.game.protocol.context.ContextRefreshEntityLookEvent
	(*EntitiesDispositionEvent)(nil),      // 11: com.ankama.dofus.server.game.protocol.context.EntitiesDispositionEvent
	(*EntityDispositionErrorEvent)(nil),   // 12: com.ankama.dofus.server.game.protocol.context.EntityDispositionErrorEvent
	(*MonsterBoosts)(nil),                 // 13: com.ankama.dofus.server.game.protocol.context.MonsterBoosts
	(*RefreshMonsterBoostsEvent)(nil),     // 14: com.ankama.dofus.server.game.protocol.context.RefreshMonsterBoostsEvent
	(*ShowCellEvent)(nil),                 // 15: com.ankama.dofus.server.game.protocol.context.ShowCellEvent
	(*common.EntityLook)(nil),             // 16: com.ankama.dofus.server.game.protocol.common.EntityLook
	(*common.EntityDisposition)(nil),      // 17: com.ankama.dofus.server.game.protocol.common.EntityDisposition
}
var file_proto_game_context_proto_depIdxs = []int32{
	0,  // 0: com.ankama.dofus.server.game.protocol.context.ContextCreationEvent.context:type_name -> com.ankama.dofus.server.game.protocol.context.ContextCreationEvent.GameContext
	16, // 1: com.ankama.dofus.server.game.protocol.context.ContextRefreshEntityLookEvent.look:type_name -> com.ankama.dofus.server.game.protocol.common.EntityLook
	17, // 2: com.ankama.dofus.server.game.protocol.context.EntitiesDispositionEvent.dispositions:type_name -> com.ankama.dofus.server.game.protocol.common.EntityDisposition
	13, // 3: com.ankama.dofus.server.game.protocol.context.RefreshMonsterBoostsEvent.monster_boosts:type_name -> com.ankama.dofus.server.game.protocol.context.MonsterBoosts
	13, // 4: com.ankama.dofus.server.game.protocol.context.RefreshMonsterBoostsEvent.family_boosts:type_name -> com.ankama.dofus.server.game.protocol.context.MonsterBoosts
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_game_context_proto_init() }
func file_proto_game_context_proto_init() {
	if File_proto_game_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_game_context_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ContextCreationRequest); i {
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
		file_proto_game_context_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ContextReadyRequest); i {
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
		file_proto_game_context_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ContextQuitRequest); i {
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
		file_proto_game_context_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ContextKickRequest); i {
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
		file_proto_game_context_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ShowCellRequest); i {
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
		file_proto_game_context_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ContextCreationEvent); i {
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
		file_proto_game_context_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ContextDestroyEvent); i {
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
		file_proto_game_context_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ContextRemoveElementEvent); i {
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
		file_proto_game_context_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ContextRemoveElementsEvent); i {
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
		file_proto_game_context_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ContextRefreshEntityLookEvent); i {
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
		file_proto_game_context_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*EntitiesDispositionEvent); i {
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
		file_proto_game_context_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*EntityDispositionErrorEvent); i {
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
		file_proto_game_context_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*MonsterBoosts); i {
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
		file_proto_game_context_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*RefreshMonsterBoostsEvent); i {
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
		file_proto_game_context_proto_msgTypes[14].Exporter = func(v any, i int) any {
			switch v := v.(*ShowCellEvent); i {
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
			RawDescriptor: file_proto_game_context_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_game_context_proto_goTypes,
		DependencyIndexes: file_proto_game_context_proto_depIdxs,
		EnumInfos:         file_proto_game_context_proto_enumTypes,
		MessageInfos:      file_proto_game_context_proto_msgTypes,
	}.Build()
	File_proto_game_context_proto = out.File
	file_proto_game_context_proto_rawDesc = nil
	file_proto_game_context_proto_goTypes = nil
	file_proto_game_context_proto_depIdxs = nil
}
