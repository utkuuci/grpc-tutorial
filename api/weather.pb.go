// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: weather.proto

package api

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

type WeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityCode string `protobuf:"bytes,1,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
}

func (x *WeatherRequest) Reset() {
	*x = WeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherRequest) ProtoMessage() {}

func (x *WeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherRequest.ProtoReflect.Descriptor instead.
func (*WeatherRequest) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0}
}

func (x *WeatherRequest) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

type WeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperatuer float32 `protobuf:"fixed32,1,opt,name=temperatuer,proto3" json:"temperatuer,omitempty"`
}

func (x *WeatherResponse) Reset() {
	*x = WeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherResponse) ProtoMessage() {}

func (x *WeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherResponse.ProtoReflect.Descriptor instead.
func (*WeatherResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{1}
}

func (x *WeatherResponse) GetTemperatuer() float32 {
	if x != nil {
		return x.Temperatuer
	}
	return 0
}

type ListCitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCitiesRequest) Reset() {
	*x = ListCitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCitiesRequest) ProtoMessage() {}

func (x *ListCitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCitiesRequest.ProtoReflect.Descriptor instead.
func (*ListCitiesRequest) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{2}
}

type ListCitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CityEntry `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListCitiesResponse) Reset() {
	*x = ListCitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCitiesResponse) ProtoMessage() {}

func (x *ListCitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCitiesResponse.ProtoReflect.Descriptor instead.
func (*ListCitiesResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{3}
}

func (x *ListCitiesResponse) GetItems() []*CityEntry {
	if x != nil {
		return x.Items
	}
	return nil
}

type CityEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityCode string `protobuf:"bytes,1,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	CityName string `protobuf:"bytes,2,opt,name=city_name,json=cityName,proto3" json:"city_name,omitempty"`
}

func (x *CityEntry) Reset() {
	*x = CityEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityEntry) ProtoMessage() {}

func (x *CityEntry) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityEntry.ProtoReflect.Descriptor instead.
func (*CityEntry) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{4}
}

func (x *CityEntry) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

func (x *CityEntry) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

var File_weather_proto protoreflect.FileDescriptor

var file_weather_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2d, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x33,
	0x0a, 0x0f, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x43, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x45, 0x0a, 0x09, 0x43, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x7c, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x35,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_proto_rawDescOnce sync.Once
	file_weather_proto_rawDescData = file_weather_proto_rawDesc
)

func file_weather_proto_rawDescGZIP() []byte {
	file_weather_proto_rawDescOnce.Do(func() {
		file_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_proto_rawDescData)
	})
	return file_weather_proto_rawDescData
}

var file_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_weather_proto_goTypes = []interface{}{
	(*WeatherRequest)(nil),     // 0: WeatherRequest
	(*WeatherResponse)(nil),    // 1: WeatherResponse
	(*ListCitiesRequest)(nil),  // 2: ListCitiesRequest
	(*ListCitiesResponse)(nil), // 3: ListCitiesResponse
	(*CityEntry)(nil),          // 4: CityEntry
}
var file_weather_proto_depIdxs = []int32{
	4, // 0: ListCitiesResponse.items:type_name -> CityEntry
	0, // 1: WeatherService.QueryWeather:input_type -> WeatherRequest
	2, // 2: WeatherService.ListCities:input_type -> ListCitiesRequest
	1, // 3: WeatherService.QueryWeather:output_type -> WeatherResponse
	3, // 4: WeatherService.ListCities:output_type -> ListCitiesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_weather_proto_init() }
func file_weather_proto_init() {
	if File_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherRequest); i {
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
		file_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherResponse); i {
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
		file_weather_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCitiesRequest); i {
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
		file_weather_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCitiesResponse); i {
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
		file_weather_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityEntry); i {
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
			RawDescriptor: file_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_proto_goTypes,
		DependencyIndexes: file_weather_proto_depIdxs,
		MessageInfos:      file_weather_proto_msgTypes,
	}.Build()
	File_weather_proto = out.File
	file_weather_proto_rawDesc = nil
	file_weather_proto_goTypes = nil
	file_weather_proto_depIdxs = nil
}
