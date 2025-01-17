// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gogoproto/gogo.proto

package gogoproto

import (
	fmt "fmt"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	proto "github.com/golang/protobuf/proto"
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

var E_GoprotoEnumPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62001,
	Name:          "gogoproto.goproto_enum_prefix",
	Tag:           "varint,62001,opt,name=goproto_enum_prefix",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoEnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62021,
	Name:          "gogoproto.goproto_enum_stringer",
	Tag:           "varint,62021,opt,name=goproto_enum_stringer",
	Filename:      "gogoproto/gogo.proto",
}

var E_EnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62022,
	Name:          "gogoproto.enum_stringer",
	Tag:           "varint,62022,opt,name=enum_stringer",
	Filename:      "gogoproto/gogo.proto",
}

var E_EnumCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         62023,
	Name:          "gogoproto.enum_customname",
	Tag:           "bytes,62023,opt,name=enum_customname",
	Filename:      "gogoproto/gogo.proto",
}

var E_Enumdecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62024,
	Name:          "gogoproto.enumdecl",
	Tag:           "varint,62024,opt,name=enumdecl",
	Filename:      "gogoproto/gogo.proto",
}

var E_EnumvalueCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         66001,
	Name:          "gogoproto.enumvalue_customname",
	Tag:           "bytes,66001,opt,name=enumvalue_customname",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoGettersAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63001,
	Name:          "gogoproto.goproto_getters_all",
	Tag:           "varint,63001,opt,name=goproto_getters_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoEnumPrefixAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63002,
	Name:          "gogoproto.goproto_enum_prefix_all",
	Tag:           "varint,63002,opt,name=goproto_enum_prefix_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63003,
	Name:          "gogoproto.goproto_stringer_all",
	Tag:           "varint,63003,opt,name=goproto_stringer_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_VerboseEqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63004,
	Name:          "gogoproto.verbose_equal_all",
	Tag:           "varint,63004,opt,name=verbose_equal_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_FaceAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63005,
	Name:          "gogoproto.face_all",
	Tag:           "varint,63005,opt,name=face_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GostringAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63006,
	Name:          "gogoproto.gostring_all",
	Tag:           "varint,63006,opt,name=gostring_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_PopulateAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63007,
	Name:          "gogoproto.populate_all",
	Tag:           "varint,63007,opt,name=populate_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_StringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63008,
	Name:          "gogoproto.stringer_all",
	Tag:           "varint,63008,opt,name=stringer_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_OnlyoneAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63009,
	Name:          "gogoproto.onlyone_all",
	Tag:           "varint,63009,opt,name=onlyone_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_EqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63013,
	Name:          "gogoproto.equal_all",
	Tag:           "varint,63013,opt,name=equal_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_DescriptionAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63014,
	Name:          "gogoproto.description_all",
	Tag:           "varint,63014,opt,name=description_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_TestgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63015,
	Name:          "gogoproto.testgen_all",
	Tag:           "varint,63015,opt,name=testgen_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_BenchgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63016,
	Name:          "gogoproto.benchgen_all",
	Tag:           "varint,63016,opt,name=benchgen_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_MarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63017,
	Name:          "gogoproto.marshaler_all",
	Tag:           "varint,63017,opt,name=marshaler_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_UnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63018,
	Name:          "gogoproto.unmarshaler_all",
	Tag:           "varint,63018,opt,name=unmarshaler_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_StableMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63019,
	Name:          "gogoproto.stable_marshaler_all",
	Tag:           "varint,63019,opt,name=stable_marshaler_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_SizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63020,
	Name:          "gogoproto.sizer_all",
	Tag:           "varint,63020,opt,name=sizer_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoEnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63021,
	Name:          "gogoproto.goproto_enum_stringer_all",
	Tag:           "varint,63021,opt,name=goproto_enum_stringer_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_EnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63022,
	Name:          "gogoproto.enum_stringer_all",
	Tag:           "varint,63022,opt,name=enum_stringer_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_UnsafeMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63023,
	Name:          "gogoproto.unsafe_marshaler_all",
	Tag:           "varint,63023,opt,name=unsafe_marshaler_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_UnsafeUnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63024,
	Name:          "gogoproto.unsafe_unmarshaler_all",
	Tag:           "varint,63024,opt,name=unsafe_unmarshaler_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoExtensionsMapAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63025,
	Name:          "gogoproto.goproto_extensions_map_all",
	Tag:           "varint,63025,opt,name=goproto_extensions_map_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoUnrecognizedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63026,
	Name:          "gogoproto.goproto_unrecognized_all",
	Tag:           "varint,63026,opt,name=goproto_unrecognized_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GogoprotoImport = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63027,
	Name:          "gogoproto.gogoproto_import",
	Tag:           "varint,63027,opt,name=gogoproto_import",
	Filename:      "gogoproto/gogo.proto",
}

var E_ProtosizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63028,
	Name:          "gogoproto.protosizer_all",
	Tag:           "varint,63028,opt,name=protosizer_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_CompareAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63029,
	Name:          "gogoproto.compare_all",
	Tag:           "varint,63029,opt,name=compare_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_TypedeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63030,
	Name:          "gogoproto.typedecl_all",
	Tag:           "varint,63030,opt,name=typedecl_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_EnumdeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63031,
	Name:          "gogoproto.enumdecl_all",
	Tag:           "varint,63031,opt,name=enumdecl_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoRegistration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63032,
	Name:          "gogoproto.goproto_registration",
	Tag:           "varint,63032,opt,name=goproto_registration",
	Filename:      "gogoproto/gogo.proto",
}

var E_MessagenameAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63033,
	Name:          "gogoproto.messagename_all",
	Tag:           "varint,63033,opt,name=messagename_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoSizecacheAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63034,
	Name:          "gogoproto.goproto_sizecache_all",
	Tag:           "varint,63034,opt,name=goproto_sizecache_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoUnkeyedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63035,
	Name:          "gogoproto.goproto_unkeyed_all",
	Tag:           "varint,63035,opt,name=goproto_unkeyed_all",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoGetters = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64001,
	Name:          "gogoproto.goproto_getters",
	Tag:           "varint,64001,opt,name=goproto_getters",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64003,
	Name:          "gogoproto.goproto_stringer",
	Tag:           "varint,64003,opt,name=goproto_stringer",
	Filename:      "gogoproto/gogo.proto",
}

var E_VerboseEqual = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64004,
	Name:          "gogoproto.verbose_equal",
	Tag:           "varint,64004,opt,name=verbose_equal",
	Filename:      "gogoproto/gogo.proto",
}

var E_Face = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64005,
	Name:          "gogoproto.face",
	Tag:           "varint,64005,opt,name=face",
	Filename:      "gogoproto/gogo.proto",
}

var E_Gostring = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64006,
	Name:          "gogoproto.gostring",
	Tag:           "varint,64006,opt,name=gostring",
	Filename:      "gogoproto/gogo.proto",
}

var E_Populate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64007,
	Name:          "gogoproto.populate",
	Tag:           "varint,64007,opt,name=populate",
	Filename:      "gogoproto/gogo.proto",
}

var E_Stringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         67008,
	Name:          "gogoproto.stringer",
	Tag:           "varint,67008,opt,name=stringer",
	Filename:      "gogoproto/gogo.proto",
}

var E_Onlyone = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64009,
	Name:          "gogoproto.onlyone",
	Tag:           "varint,64009,opt,name=onlyone",
	Filename:      "gogoproto/gogo.proto",
}

var E_Equal = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64013,
	Name:          "gogoproto.equal",
	Tag:           "varint,64013,opt,name=equal",
	Filename:      "gogoproto/gogo.proto",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64014,
	Name:          "gogoproto.description",
	Tag:           "varint,64014,opt,name=description",
	Filename:      "gogoproto/gogo.proto",
}

var E_Testgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64015,
	Name:          "gogoproto.testgen",
	Tag:           "varint,64015,opt,name=testgen",
	Filename:      "gogoproto/gogo.proto",
}

var E_Benchgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64016,
	Name:          "gogoproto.benchgen",
	Tag:           "varint,64016,opt,name=benchgen",
	Filename:      "gogoproto/gogo.proto",
}

var E_Marshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64017,
	Name:          "gogoproto.marshaler",
	Tag:           "varint,64017,opt,name=marshaler",
	Filename:      "gogoproto/gogo.proto",
}

var E_Unmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64018,
	Name:          "gogoproto.unmarshaler",
	Tag:           "varint,64018,opt,name=unmarshaler",
	Filename:      "gogoproto/gogo.proto",
}

var E_StableMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64019,
	Name:          "gogoproto.stable_marshaler",
	Tag:           "varint,64019,opt,name=stable_marshaler",
	Filename:      "gogoproto/gogo.proto",
}

var E_Sizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64020,
	Name:          "gogoproto.sizer",
	Tag:           "varint,64020,opt,name=sizer",
	Filename:      "gogoproto/gogo.proto",
}

var E_UnsafeMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64023,
	Name:          "gogoproto.unsafe_marshaler",
	Tag:           "varint,64023,opt,name=unsafe_marshaler",
	Filename:      "gogoproto/gogo.proto",
}

var E_UnsafeUnmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64024,
	Name:          "gogoproto.unsafe_unmarshaler",
	Tag:           "varint,64024,opt,name=unsafe_unmarshaler",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoExtensionsMap = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64025,
	Name:          "gogoproto.goproto_extensions_map",
	Tag:           "varint,64025,opt,name=goproto_extensions_map",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoUnrecognized = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64026,
	Name:          "gogoproto.goproto_unrecognized",
	Tag:           "varint,64026,opt,name=goproto_unrecognized",
	Filename:      "gogoproto/gogo.proto",
}

var E_Protosizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64028,
	Name:          "gogoproto.protosizer",
	Tag:           "varint,64028,opt,name=protosizer",
	Filename:      "gogoproto/gogo.proto",
}

var E_Compare = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64029,
	Name:          "gogoproto.compare",
	Tag:           "varint,64029,opt,name=compare",
	Filename:      "gogoproto/gogo.proto",
}

var E_Typedecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64030,
	Name:          "gogoproto.typedecl",
	Tag:           "varint,64030,opt,name=typedecl",
	Filename:      "gogoproto/gogo.proto",
}

var E_Messagename = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64033,
	Name:          "gogoproto.messagename",
	Tag:           "varint,64033,opt,name=messagename",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoSizecache = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64034,
	Name:          "gogoproto.goproto_sizecache",
	Tag:           "varint,64034,opt,name=goproto_sizecache",
	Filename:      "gogoproto/gogo.proto",
}

var E_GoprotoUnkeyed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64035,
	Name:          "gogoproto.goproto_unkeyed",
	Tag:           "varint,64035,opt,name=goproto_unkeyed",
	Filename:      "gogoproto/gogo.proto",
}

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
	Filename:      "gogoproto/gogo.proto",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65002,
	Name:          "gogoproto.embed",
	Tag:           "varint,65002,opt,name=embed",
	Filename:      "gogoproto/gogo.proto",
}

var E_Customtype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "gogoproto.customtype",
	Tag:           "bytes,65003,opt,name=customtype",
	Filename:      "gogoproto/gogo.proto",
}

var E_Customname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65004,
	Name:          "gogoproto.customname",
	Tag:           "bytes,65004,opt,name=customname",
	Filename:      "gogoproto/gogo.proto",
}

var E_Jsontag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65005,
	Name:          "gogoproto.jsontag",
	Tag:           "bytes,65005,opt,name=jsontag",
	Filename:      "gogoproto/gogo.proto",
}

var E_Moretags = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65006,
	Name:          "gogoproto.moretags",
	Tag:           "bytes,65006,opt,name=moretags",
	Filename:      "gogoproto/gogo.proto",
}

var E_Casttype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65007,
	Name:          "gogoproto.casttype",
	Tag:           "bytes,65007,opt,name=casttype",
	Filename:      "gogoproto/gogo.proto",
}

var E_Castkey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65008,
	Name:          "gogoproto.castkey",
	Tag:           "bytes,65008,opt,name=castkey",
	Filename:      "gogoproto/gogo.proto",
}

var E_Castvalue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65009,
	Name:          "gogoproto.castvalue",
	Tag:           "bytes,65009,opt,name=castvalue",
	Filename:      "gogoproto/gogo.proto",
}

var E_Stdtime = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65010,
	Name:          "gogoproto.stdtime",
	Tag:           "varint,65010,opt,name=stdtime",
	Filename:      "gogoproto/gogo.proto",
}

var E_Stdduration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65011,
	Name:          "gogoproto.stdduration",
	Tag:           "varint,65011,opt,name=stdduration",
	Filename:      "gogoproto/gogo.proto",
}

var E_Wktpointer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65012,
	Name:          "gogoproto.wktpointer",
	Tag:           "varint,65012,opt,name=wktpointer",
	Filename:      "gogoproto/gogo.proto",
}

func init() {
	proto.RegisterExtension(E_GoprotoEnumPrefix)
	proto.RegisterExtension(E_GoprotoEnumStringer)
	proto.RegisterExtension(E_EnumStringer)
	proto.RegisterExtension(E_EnumCustomname)
	proto.RegisterExtension(E_Enumdecl)
	proto.RegisterExtension(E_EnumvalueCustomname)
	proto.RegisterExtension(E_GoprotoGettersAll)
	proto.RegisterExtension(E_GoprotoEnumPrefixAll)
	proto.RegisterExtension(E_GoprotoStringerAll)
	proto.RegisterExtension(E_VerboseEqualAll)
	proto.RegisterExtension(E_FaceAll)
	proto.RegisterExtension(E_GostringAll)
	proto.RegisterExtension(E_PopulateAll)
	proto.RegisterExtension(E_StringerAll)
	proto.RegisterExtension(E_OnlyoneAll)
	proto.RegisterExtension(E_EqualAll)
	proto.RegisterExtension(E_DescriptionAll)
	proto.RegisterExtension(E_TestgenAll)
	proto.RegisterExtension(E_BenchgenAll)
	proto.RegisterExtension(E_MarshalerAll)
	proto.RegisterExtension(E_UnmarshalerAll)
	proto.RegisterExtension(E_StableMarshalerAll)
	proto.RegisterExtension(E_SizerAll)
	proto.RegisterExtension(E_GoprotoEnumStringerAll)
	proto.RegisterExtension(E_EnumStringerAll)
	proto.RegisterExtension(E_UnsafeMarshalerAll)
	proto.RegisterExtension(E_UnsafeUnmarshalerAll)
	proto.RegisterExtension(E_GoprotoExtensionsMapAll)
	proto.RegisterExtension(E_GoprotoUnrecognizedAll)
	proto.RegisterExtension(E_GogoprotoImport)
	proto.RegisterExtension(E_ProtosizerAll)
	proto.RegisterExtension(E_CompareAll)
	proto.RegisterExtension(E_TypedeclAll)
	proto.RegisterExtension(E_EnumdeclAll)
	proto.RegisterExtension(E_GoprotoRegistration)
	proto.RegisterExtension(E_MessagenameAll)
	proto.RegisterExtension(E_GoprotoSizecacheAll)
	proto.RegisterExtension(E_GoprotoUnkeyedAll)
	proto.RegisterExtension(E_GoprotoGetters)
	proto.RegisterExtension(E_GoprotoStringer)
	proto.RegisterExtension(E_VerboseEqual)
	proto.RegisterExtension(E_Face)
	proto.RegisterExtension(E_Gostring)
	proto.RegisterExtension(E_Populate)
	proto.RegisterExtension(E_Stringer)
	proto.RegisterExtension(E_Onlyone)
	proto.RegisterExtension(E_Equal)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_Testgen)
	proto.RegisterExtension(E_Benchgen)
	proto.RegisterExtension(E_Marshaler)
	proto.RegisterExtension(E_Unmarshaler)
	proto.RegisterExtension(E_StableMarshaler)
	proto.RegisterExtension(E_Sizer)
	proto.RegisterExtension(E_UnsafeMarshaler)
	proto.RegisterExtension(E_UnsafeUnmarshaler)
	proto.RegisterExtension(E_GoprotoExtensionsMap)
	proto.RegisterExtension(E_GoprotoUnrecognized)
	proto.RegisterExtension(E_Protosizer)
	proto.RegisterExtension(E_Compare)
	proto.RegisterExtension(E_Typedecl)
	proto.RegisterExtension(E_Messagename)
	proto.RegisterExtension(E_GoprotoSizecache)
	proto.RegisterExtension(E_GoprotoUnkeyed)
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Customtype)
	proto.RegisterExtension(E_Customname)
	proto.RegisterExtension(E_Jsontag)
	proto.RegisterExtension(E_Moretags)
	proto.RegisterExtension(E_Casttype)
	proto.RegisterExtension(E_Castkey)
	proto.RegisterExtension(E_Castvalue)
	proto.RegisterExtension(E_Stdtime)
	proto.RegisterExtension(E_Stdduration)
	proto.RegisterExtension(E_Wktpointer)
}

func init() { proto.RegisterFile("gogoproto/gogo.proto", fileDescriptor_c586470e9b64aee7) }

var fileDescriptor_c586470e9b64aee7 = []byte{
	// 1381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0xc9, 0x6f, 0x1c, 0x45,
	0x17, 0xc0, 0x3f, 0xeb, 0x4b, 0x14, 0x4f, 0xd9, 0x8e, 0xe3, 0xb1, 0x09, 0x21, 0x02, 0x13, 0x6e,
	0x5c, 0xe2, 0x39, 0x45, 0x28, 0x65, 0x45, 0xc6, 0xb1, 0x1c, 0x2b, 0x08, 0x07, 0xe3, 0xc4, 0x61,
	0x3b, 0x8c, 0x7a, 0x7a, 0xca, 0xed, 0xc6, 0xdd, 0x5d, 0x4d, 0x57, 0x75, 0x88, 0x73, 0x43, 0x61,
	0x11, 0x42, 0xec, 0x48, 0x24, 0x21, 0x09, 0x04, 0xc4, 0xbe, 0x86, 0x7d, 0xb9, 0x70, 0x01, 0x72,
	0x84, 0xff, 0x00, 0x19, 0x2e, 0x80, 0xd9, 0x7d, 0xf3, 0x05, 0x55, 0xf7, 0x7b, 0x3d, 0x35, 0xed,
	0x91, 0xaa, 0xe6, 0xd6, 0x1e, 0xd7, 0xef, 0x37, 0xd5, 0xef, 0x75, 0xbd, 0xf7, 0xa6, 0xc9, 0x88,
	0xc7, 0x3d, 0x1e, 0x27, 0x5c, 0xf2, 0x9a, 0xba, 0x1a, 0xcb, 0x2e, 0xab, 0x95, 0xe2, 0xd3, 0xdd,
	0x7b, 0x3c, 0xce, 0xbd, 0x80, 0xd5, 0xb2, 0xbf, 0x1a, 0xe9, 0x62, 0xad, 0xc9, 0x84, 0x9b, 0xf8,
	0xb1, 0xe4, 0x49, 0xbe, 0x98, 0x1e, 0x21, 0xc3, 0xb0, 0xb8, 0xce, 0xa2, 0x34, 0xac, 0xc7, 0x09,
	0x5b, 0xf4, 0x4f, 0x56, 0xaf, 0x1d, 0xcb, 0xc9, 0x31, 0x24, 0xc7, 0xa6, 0xa3, 0x34, 0xbc, 0x2d,
	0x96, 0x3e, 0x8f, 0xc4, 0xae, 0xcb, 0x3f, 0xff, 0x7f, 0x4f, 0xcf, 0x8d, 0xbd, 0xf3, 0x43, 0x80,
	0xaa, 0xff, 0xcd, 0x65, 0x20, 0x9d, 0x27, 0x57, 0xb5, 0xf9, 0x84, 0x4c, 0xfc, 0xc8, 0x63, 0x89,
	0xc1, 0xf8, 0x0d, 0x18, 0x87, 0x35, 0xe3, 0x51, 0x40, 0xe9, 0x14, 0x19, 0xe8, 0xc6, 0xf5, 0x2d,
	0xb8, 0xfa, 0x99, 0x2e, 0x99, 0x21, 0x83, 0x99, 0xc4, 0x4d, 0x85, 0xe4, 0x61, 0xe4, 0x84, 0xcc,
	0xa0, 0xf9, 0x2e, 0xd3, 0x54, 0xe6, 0xb7, 0x2b, 0x6c, 0xaa, 0xa0, 0x28, 0x25, 0xbd, 0xea, 0x93,
	0x26, 0x73, 0x03, 0x83, 0xe1, 0x0a, 0x6c, 0xa4, 0x58, 0x4f, 0x8f, 0x93, 0x11, 0x75, 0x7d, 0xc2,
	0x09, 0x52, 0xa6, 0xef, 0xe4, 0x86, 0x8e, 0x9e, 0xe3, 0x6a, 0x19, 0xca, 0x7e, 0x38, 0xbd, 0x25,
	0xdb, 0xce, 0x70, 0x21, 0xd0, 0xf6, 0xa4, 0x65, 0xd1, 0x63, 0x52, 0xb2, 0x44, 0xd4, 0x9d, 0xa0,
	0xd3, 0xf6, 0x0e, 0xf9, 0x41, 0x61, 0x3c, 0xbb, 0xd6, 0x9e, 0xc5, 0x99, 0x9c, 0x9c, 0x0c, 0x02,
	0xba, 0x40, 0xae, 0xee, 0xf0, 0x54, 0x58, 0x38, 0xcf, 0x81, 0x73, 0x64, 0xd3, 0x93, 0xa1, 0xb4,
	0x73, 0x04, 0x3f, 0x2f, 0x72, 0x69, 0xe1, 0x7c, 0x01, 0x9c, 0x55, 0x60, 0x31, 0xa5, 0xca, 0x78,
	0x0b, 0x19, 0x3a, 0xc1, 0x92, 0x06, 0x17, 0xac, 0xce, 0xee, 0x4b, 0x9d, 0xc0, 0x42, 0x77, 0x1e,
	0x74, 0x83, 0x00, 0x4e, 0x2b, 0x4e, 0xb9, 0xf6, 0x93, 0xde, 0x45, 0xc7, 0x65, 0x16, 0x8a, 0x0b,
	0xa0, 0xd8, 0xa6, 0xd6, 0x2b, 0x74, 0x92, 0xf4, 0x7b, 0x3c, 0xbf, 0x25, 0x0b, 0xfc, 0x22, 0xe0,
	0x7d, 0xc8, 0x80, 0x22, 0xe6, 0x71, 0x1a, 0x38, 0xd2, 0x66, 0x07, 0x2f, 0xa2, 0x02, 0x19, 0x50,
	0x74, 0x11, 0xd6, 0x97, 0x50, 0x21, 0xb4, 0x78, 0x4e, 0x90, 0x3e, 0x1e, 0x05, 0x2b, 0x3c, 0xb2,
	0xd9, 0xc4, 0x25, 0x30, 0x10, 0x40, 0x94, 0x60, 0x9c, 0x54, 0x6c, 0x13, 0xf1, 0xea, 0x1a, 0x1e,
	0x0f, 0xcc, 0xc0, 0x0c, 0x19, 0xc4, 0x02, 0xe5, 0xf3, 0xc8, 0x42, 0xf1, 0x1a, 0x28, 0xb6, 0x6b,
	0x18, 0xdc, 0x86, 0x64, 0x42, 0x7a, 0xcc, 0x46, 0xf2, 0x3a, 0xde, 0x06, 0x20, 0x10, 0xca, 0x06,
	0x8b, 0xdc, 0x25, 0x3b, 0xc3, 0x1b, 0x18, 0x4a, 0x64, 0x94, 0x62, 0x8a, 0x0c, 0x84, 0x4e, 0x22,
	0x96, 0x9c, 0xc0, 0x2a, 0x1d, 0x6f, 0x82, 0xa3, 0xbf, 0x80, 0x20, 0x22, 0x69, 0xd4, 0x8d, 0xe6,
	0x2d, 0x8c, 0x88, 0x86, 0xc1, 0xd1, 0x13, 0xd2, 0x69, 0x04, 0xac, 0xde, 0x8d, 0xed, 0x6d, 0x3c,
	0x7a, 0x39, 0x3b, 0xab, 0x1b, 0xc7, 0x49, 0x45, 0xf8, 0xa7, 0xac, 0x34, 0xef, 0x60, 0xa6, 0x33,
	0x40, 0xc1, 0x77, 0x91, 0x6b, 0x3a, 0xb6, 0x09, 0x0b, 0xd9, 0xbb, 0x20, 0xdb, 0xd9, 0xa1, 0x55,
	0x40, 0x49, 0xe8, 0x56, 0xf9, 0x1e, 0x96, 0x04, 0x56, 0x72, 0xcd, 0x91, 0x91, 0x34, 0x12, 0xce,
	0x62, 0x77, 0x51, 0x7b, 0x1f, 0xa3, 0x96, 0xb3, 0x6d, 0x51, 0x3b, 0x46, 0x76, 0x82, 0xb1, 0xbb,
	0xbc, 0x7e, 0x80, 0x85, 0x35, 0xa7, 0x17, 0xda, 0xb3, 0x7b, 0x0f, 0xd9, 0x5d, 0x84, 0xf3, 0xa4,
	0x64, 0x91, 0x50, 0x4c, 0x3d, 0x74, 0x62, 0x0b, 0xf3, 0x65, 0x30, 0x63, 0xc5, 0x9f, 0x2e, 0x04,
	0xb3, 0x4e, 0xac, 0xe4, 0x77, 0x92, 0x5d, 0x28, 0x4f, 0xa3, 0x84, 0xb9, 0xdc, 0x8b, 0xfc, 0x53,
	0xac, 0x69, 0xa1, 0xfe, 0xb0, 0x94, 0xaa, 0x05, 0x0d, 0x57, 0xe6, 0xc3, 0x64, 0x47, 0x31, 0xab,
	0xd4, 0xfd, 0x30, 0xe6, 0x89, 0x34, 0x18, 0x3f, 0xc2, 0x4c, 0x15, 0xdc, 0xe1, 0x0c, 0xa3, 0xd3,
	0x64, 0x7b, 0xf6, 0xa7, 0xed, 0x23, 0xf9, 0x31, 0x88, 0x06, 0x5a, 0x14, 0x14, 0x0e, 0x97, 0x87,
	0xb1, 0x93, 0xd8, 0xd4, 0xbf, 0x4f, 0xb0, 0x70, 0x00, 0x02, 0x85, 0x43, 0xae, 0xc4, 0x4c, 0x75,
	0x7b, 0x0b, 0xc3, 0xa7, 0x58, 0x38, 0x90, 0x01, 0x05, 0x0e, 0x0c, 0x16, 0x8a, 0xcf, 0x50, 0x81,
	0x8c, 0x52, 0xdc, 0xde, 0x6a, 0xb4, 0x09, 0xf3, 0x7c, 0x21, 0x13, 0x47, 0xad, 0x36, 0xa8, 0x3e,
	0x5f, 0x6b, 0x1f, 0xc2, 0xe6, 0x35, 0x54, 0x55, 0xa2, 0x90, 0x09, 0xe1, 0x78, 0x4c, 0x4d, 0x1c,
	0x16, 0x1b, 0xfb, 0x02, 0x2b, 0x91, 0x86, 0xa9, 0xbd, 0x69, 0x13, 0xa2, 0x0a, 0xbb, 0xeb, 0xb8,
	0x4b, 0x36, 0xba, 0x2f, 0x4b, 0x9b, 0x3b, 0x8a, 0xac, 0x72, 0x6a, 0xf3, 0x4f, 0x1a, 0x2d, 0xb3,
	0x15, 0xab, 0xa7, 0xf3, 0xab, 0xd2, 0xfc, 0xb3, 0x90, 0x93, 0x79, 0x0d, 0x19, 0x2c, 0xcd, 0x53,
	0xd5, 0xeb, 0x37, 0xb9, 0x66, 0xf3, 0xfb, 0x42, 0xdd, 0x03, 0xeb, 0x70, 0xbf, 0xed, 0xe3, 0x14,
	0xbd, 0x55, 0x3d, 0xe4, 0xed, 0x43, 0x8f, 0x59, 0x76, 0x7a, 0xbd, 0x78, 0xce, 0xdb, 0x66, 0x1e,
	0x7a, 0x88, 0x0c, 0xb4, 0x0d, 0x3c, 0x66, 0xd5, 0x83, 0xa0, 0xea, 0xd7, 0xe7, 0x1d, 0xba, 0x8f,
	0x6c, 0x51, 0xc3, 0x8b, 0x19, 0x7f, 0x08, 0xf0, 0x6c, 0x39, 0x3d, 0x40, 0x7a, 0x71, 0x68, 0x31,
	0xa3, 0x0f, 0x03, 0x5a, 0x20, 0x0a, 0xc7, 0x81, 0xc5, 0x8c, 0x3f, 0x82, 0x38, 0x22, 0x0a, 0xb7,
	0x0f, 0xe1, 0xd7, 0x8f, 0x6d, 0x81, 0xa6, 0x83, 0xb1, 0x1b, 0x27, 0xdb, 0x60, 0x52, 0x31, 0xd3,
	0x8f, 0xc2, 0x97, 0x23, 0x41, 0x6f, 0x22, 0x5b, 0x2d, 0x03, 0xfe, 0x38, 0xa0, 0xf9, 0x7a, 0x3a,
	0x45, 0xfa, 0xb4, 0xe9, 0xc4, 0x8c, 0x3f, 0x01, 0xb8, 0x4e, 0xa9, 0xad, 0xc3, 0x74, 0x62, 0x16,
	0x3c, 0x89, 0x5b, 0x07, 0x42, 0x85, 0x0d, 0x07, 0x13, 0x33, 0xfd, 0x14, 0x46, 0x1d, 0x11, 0x3a,
	0x41, 0x2a, 0x45, 0xb3, 0x31, 0xf3, 0x4f, 0x03, 0xdf, 0x62, 0x54, 0x04, 0xb4, 0x66, 0x67, 0x56,
	0x3c, 0x83, 0x11, 0xd0, 0x28, 0x75, 0x8c, 0xca, 0x03, 0x8c, 0xd9, 0xf4, 0x2c, 0x1e, 0xa3, 0xd2,
	0xfc, 0xa2, 0xb2, 0x99, 0xd5, 0x7c, 0xb3, 0xe2, 0x39, 0xcc, 0x66, 0xb6, 0x5e, 0x6d, 0xa3, 0x3c,
	0x11, 0x98, 0x1d, 0xcf, 0xe3, 0x36, 0x4a, 0x03, 0x01, 0x9d, 0x23, 0xd5, 0xcd, 0xd3, 0x80, 0xd9,
	0x77, 0x06, 0x7c, 0x43, 0x9b, 0x86, 0x01, 0x7a, 0x07, 0xd9, 0xd9, 0x79, 0x12, 0x30, 0x5b, 0xcf,
	0xae, 0x97, 0x7e, 0xbb, 0xe9, 0x83, 0x00, 0x3d, 0xd6, 0x6a, 0x29, 0xfa, 0x14, 0x60, 0xd6, 0x9e,
	0x5b, 0x6f, 0x2f, 0xdc, 0xfa, 0x10, 0x40, 0x27, 0x09, 0x69, 0x35, 0x60, 0xb3, 0xeb, 0x3c, 0xb8,
	0x34, 0x48, 0x1d, 0x0d, 0xe8, 0xbf, 0x66, 0xfe, 0x02, 0x1e, 0x0d, 0x20, 0xd4, 0xd1, 0xc0, 0xd6,
	0x6b, 0xa6, 0x2f, 0xe2, 0xd1, 0x40, 0x44, 0x3d, 0xd9, 0x5a, 0x77, 0x33, 0x1b, 0x2e, 0xe1, 0x93,
	0xad, 0x51, 0xf4, 0x08, 0x19, 0xda, 0xd4, 0x10, 0xcd, 0xaa, 0x97, 0x41, 0xb5, 0xa3, 0xdc, 0x0f,
	0xf5, 0xe6, 0x05, 0xcd, 0xd0, 0x6c, 0x7b, 0xa5, 0xd4, 0xbc, 0xa0, 0x17, 0xd2, 0x71, 0xd2, 0x1b,
	0xa5, 0x41, 0xa0, 0x0e, 0x4f, 0xf5, 0xba, 0x0e, 0xdd, 0x94, 0x05, 0x4d, 0x54, 0xfc, 0xb2, 0x01,
	0xd1, 0x41, 0x80, 0xee, 0x23, 0x5b, 0x59, 0xd8, 0x60, 0x4d, 0x13, 0xf9, 0xeb, 0x06, 0x16, 0x4c,
	0xb5, 0x9a, 0x4e, 0x10, 0x92, 0xbf, 0x1a, 0x51, 0x61, 0x36, 0xb1, 0xbf, 0x6d, 0xe4, 0x6f, 0x69,
	0x34, 0xa4, 0x25, 0xc8, 0x92, 0x62, 0x10, 0xac, 0xb5, 0x0b, 0xb2, 0x8c, 0xec, 0x27, 0xdb, 0xee,
	0x15, 0x3c, 0x92, 0x8e, 0x67, 0xa2, 0x7f, 0x07, 0x1a, 0xd7, 0xab, 0x80, 0x85, 0x3c, 0x61, 0xd2,
	0xf1, 0x84, 0x89, 0xfd, 0x03, 0xd8, 0x02, 0x50, 0xb0, 0xeb, 0x08, 0x69, 0x73, 0xdf, 0x7f, 0x22,
	0x8c, 0x80, 0xda, 0xb4, 0xba, 0x5e, 0x66, 0x2b, 0x26, 0xf6, 0x2f, 0xdc, 0x34, 0xac, 0xa7, 0x07,
	0x48, 0x45, 0x5d, 0x66, 0x6f, 0x95, 0x4c, 0xf0, 0xdf, 0x00, 0xb7, 0x08, 0xf5, 0xcd, 0x42, 0x36,
	0xa5, 0x6f, 0x0e, 0xf6, 0x3f, 0x90, 0x69, 0x5c, 0x4f, 0x27, 0x49, 0x9f, 0x90, 0xcd, 0x66, 0x0a,
	0xf3, 0xa9, 0x01, 0xff, 0x77, 0xa3, 0x78, 0x65, 0x51, 0x30, 0x2a, 0xdb, 0xf7, 0x2f, 0xcb, 0x98,
	0xfb, 0x91, 0x64, 0x89, 0xc9, 0xb0, 0x0e, 0x06, 0x0d, 0x39, 0x28, 0xae, 0xac, 0x8e, 0xf6, 0x7c,
	0xbf, 0x3a, 0xda, 0xf3, 0xe3, 0xea, 0x68, 0xcf, 0x99, 0x9f, 0x46, 0xff, 0x47, 0x86, 0x5d, 0x1e,
	0x96, 0x3d, 0x07, 0xc9, 0x0c, 0x9f, 0xe1, 0x73, 0x59, 0xdd, 0xb9, 0xfb, 0x66, 0xcf, 0x97, 0x4b,
	0x69, 0x63, 0xcc, 0xe5, 0x61, 0xcd, 0x0d, 0x78, 0xda, 0xdc, 0xdb, 0x70, 0x12, 0x5f, 0x48, 0xa7,
	0xe6, 0x36, 0xf6, 0x86, 0xee, 0xb2, 0xa8, 0x89, 0xc4, 0xad, 0x79, 0x49, 0xec, 0xee, 0x75, 0x62,
	0xbf, 0xf5, 0xfa, 0xb5, 0xf8, 0xd5, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x38, 0xd4,
	0x1b, 0xbb, 0x15, 0x00, 0x00,
}
