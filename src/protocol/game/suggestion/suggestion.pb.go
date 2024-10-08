// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/game/suggestion.proto

package suggestion

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

type ActivitySuggestionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinLevel           int32 `protobuf:"varint,1,opt,name=min_level,json=minLevel,proto3" json:"min_level,omitempty"`
	MaxLevel           int32 `protobuf:"varint,2,opt,name=max_level,json=maxLevel,proto3" json:"max_level,omitempty"`
	AreaId             int32 `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	ActivityCategoryId int32 `protobuf:"varint,4,opt,name=activity_category_id,json=activityCategoryId,proto3" json:"activity_category_id,omitempty"`
	CardsNumber        int32 `protobuf:"varint,5,opt,name=cards_number,json=cardsNumber,proto3" json:"cards_number,omitempty"`
}

func (x *ActivitySuggestionsRequest) Reset() {
	*x = ActivitySuggestionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_suggestion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivitySuggestionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivitySuggestionsRequest) ProtoMessage() {}

func (x *ActivitySuggestionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_suggestion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivitySuggestionsRequest.ProtoReflect.Descriptor instead.
func (*ActivitySuggestionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_suggestion_proto_rawDescGZIP(), []int{0}
}

func (x *ActivitySuggestionsRequest) GetMinLevel() int32 {
	if x != nil {
		return x.MinLevel
	}
	return 0
}

func (x *ActivitySuggestionsRequest) GetMaxLevel() int32 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

func (x *ActivitySuggestionsRequest) GetAreaId() int32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *ActivitySuggestionsRequest) GetActivityCategoryId() int32 {
	if x != nil {
		return x.ActivityCategoryId
	}
	return 0
}

func (x *ActivitySuggestionsRequest) GetCardsNumber() int32 {
	if x != nil {
		return x.CardsNumber
	}
	return 0
}

type ActivityHideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId int32 `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *ActivityHideRequest) Reset() {
	*x = ActivityHideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_suggestion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityHideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityHideRequest) ProtoMessage() {}

func (x *ActivityHideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_suggestion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityHideRequest.ProtoReflect.Descriptor instead.
func (*ActivityHideRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_suggestion_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityHideRequest) GetActivityId() int32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

type ActivityLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId int32 `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *ActivityLockRequest) Reset() {
	*x = ActivityLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_suggestion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLockRequest) ProtoMessage() {}

func (x *ActivityLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_suggestion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLockRequest.ProtoReflect.Descriptor instead.
func (*ActivityLockRequest) Descriptor() ([]byte, []int) {
	return file_proto_game_suggestion_proto_rawDescGZIP(), []int{2}
}

func (x *ActivityLockRequest) GetActivityId() int32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

type ActivitySuggestionsEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockedActivities   []int32 `protobuf:"varint,1,rep,packed,name=locked_activities,json=lockedActivities,proto3" json:"locked_activities,omitempty"`
	UnlockedActivities []int32 `protobuf:"varint,2,rep,packed,name=unlocked_activities,json=unlockedActivities,proto3" json:"unlocked_activities,omitempty"`
}

func (x *ActivitySuggestionsEvent) Reset() {
	*x = ActivitySuggestionsEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_game_suggestion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivitySuggestionsEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivitySuggestionsEvent) ProtoMessage() {}

func (x *ActivitySuggestionsEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_game_suggestion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivitySuggestionsEvent.ProtoReflect.Descriptor instead.
func (*ActivitySuggestionsEvent) Descriptor() ([]byte, []int) {
	return file_proto_game_suggestion_proto_rawDescGZIP(), []int{3}
}

func (x *ActivitySuggestionsEvent) GetLockedActivities() []int32 {
	if x != nil {
		return x.LockedActivities
	}
	return nil
}

func (x *ActivitySuggestionsEvent) GetUnlockedActivities() []int32 {
	if x != nil {
		return x.UnlockedActivities
	}
	return nil
}

var File_proto_game_suggestion_proto protoreflect.FileDescriptor

var file_proto_game_suggestion_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x6f, 0x66, 0x75, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xc4, 0x01, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x12, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x13, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x48, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x36,
	0x0a, 0x13, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x18, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x10, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x2f, 0x0a, 0x13, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x6f, 0x2d, 0x78, 0x70, 0x2d, 0x64, 0x6f, 0x66, 0x75, 0x73, 0x2d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_game_suggestion_proto_rawDescOnce sync.Once
	file_proto_game_suggestion_proto_rawDescData = file_proto_game_suggestion_proto_rawDesc
)

func file_proto_game_suggestion_proto_rawDescGZIP() []byte {
	file_proto_game_suggestion_proto_rawDescOnce.Do(func() {
		file_proto_game_suggestion_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_game_suggestion_proto_rawDescData)
	})
	return file_proto_game_suggestion_proto_rawDescData
}

var file_proto_game_suggestion_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_game_suggestion_proto_goTypes = []any{
	(*ActivitySuggestionsRequest)(nil), // 0: com.ankama.dofus.server.game.protocol.suggestion.ActivitySuggestionsRequest
	(*ActivityHideRequest)(nil),        // 1: com.ankama.dofus.server.game.protocol.suggestion.ActivityHideRequest
	(*ActivityLockRequest)(nil),        // 2: com.ankama.dofus.server.game.protocol.suggestion.ActivityLockRequest
	(*ActivitySuggestionsEvent)(nil),   // 3: com.ankama.dofus.server.game.protocol.suggestion.ActivitySuggestionsEvent
}
var file_proto_game_suggestion_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_game_suggestion_proto_init() }
func file_proto_game_suggestion_proto_init() {
	if File_proto_game_suggestion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_game_suggestion_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActivitySuggestionsRequest); i {
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
		file_proto_game_suggestion_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityHideRequest); i {
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
		file_proto_game_suggestion_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityLockRequest); i {
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
		file_proto_game_suggestion_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ActivitySuggestionsEvent); i {
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
			RawDescriptor: file_proto_game_suggestion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_game_suggestion_proto_goTypes,
		DependencyIndexes: file_proto_game_suggestion_proto_depIdxs,
		MessageInfos:      file_proto_game_suggestion_proto_msgTypes,
	}.Build()
	File_proto_game_suggestion_proto = out.File
	file_proto_game_suggestion_proto_rawDesc = nil
	file_proto_game_suggestion_proto_goTypes = nil
	file_proto_game_suggestion_proto_depIdxs = nil
}
