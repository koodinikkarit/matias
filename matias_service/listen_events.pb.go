// Code generated by protoc-gen-go. DO NOT EDIT.
// source: listen_events.proto

/*
Package MatiasService is a generated protocol buffer package.

It is generated from these files:
	listen_events.proto
	matias_service.proto

It has these top-level messages:
	ListenEventsRequest
	EwDatabase
	SongDatabaseVariation
	SongDatabaseTag
	TagVariation
	Variation
	Song
	EventItem
	RequestMatiasKeyRequest
	RequestMatiasKeyResponse
*/
package MatiasService

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

type ListenEventsRequest struct {
	Key               string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	HostName          string `protobuf:"bytes,2,opt,name=hostName" json:"hostName,omitempty"`
	LastestConnection int64  `protobuf:"varint,3,opt,name=lastestConnection" json:"lastestConnection,omitempty"`
}

func (m *ListenEventsRequest) Reset()                    { *m = ListenEventsRequest{} }
func (m *ListenEventsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenEventsRequest) ProtoMessage()               {}
func (*ListenEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListenEventsRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ListenEventsRequest) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *ListenEventsRequest) GetLastestConnection() int64 {
	if m != nil {
		return m.LastestConnection
	}
	return 0
}

type EwDatabase struct {
	Id             uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FilesystemPath string `protobuf:"bytes,2,opt,name=filesystemPath" json:"filesystemPath,omitempty"`
	SongDatabaseId uint32 `protobuf:"varint,3,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
}

func (m *EwDatabase) Reset()                    { *m = EwDatabase{} }
func (m *EwDatabase) String() string            { return proto.CompactTextString(m) }
func (*EwDatabase) ProtoMessage()               {}
func (*EwDatabase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EwDatabase) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EwDatabase) GetFilesystemPath() string {
	if m != nil {
		return m.FilesystemPath
	}
	return ""
}

func (m *EwDatabase) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

type SongDatabaseVariation struct {
	Id             uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SongDatabaseId uint32 `protobuf:"varint,2,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	VariationId    uint32 `protobuf:"varint,3,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *SongDatabaseVariation) Reset()                    { *m = SongDatabaseVariation{} }
func (m *SongDatabaseVariation) String() string            { return proto.CompactTextString(m) }
func (*SongDatabaseVariation) ProtoMessage()               {}
func (*SongDatabaseVariation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SongDatabaseVariation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SongDatabaseVariation) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *SongDatabaseVariation) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type SongDatabaseTag struct {
	Id             uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SongDatabaseId uint32 `protobuf:"varint,2,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	TagId          uint32 `protobuf:"varint,3,opt,name=tagId" json:"tagId,omitempty"`
}

func (m *SongDatabaseTag) Reset()                    { *m = SongDatabaseTag{} }
func (m *SongDatabaseTag) String() string            { return proto.CompactTextString(m) }
func (*SongDatabaseTag) ProtoMessage()               {}
func (*SongDatabaseTag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SongDatabaseTag) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SongDatabaseTag) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *SongDatabaseTag) GetTagId() uint32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

type TagVariation struct {
	Id          uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TagId       uint32 `protobuf:"varint,2,opt,name=tagId" json:"tagId,omitempty"`
	VariationId uint32 `protobuf:"varint,3,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *TagVariation) Reset()                    { *m = TagVariation{} }
func (m *TagVariation) String() string            { return proto.CompactTextString(m) }
func (*TagVariation) ProtoMessage()               {}
func (*TagVariation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TagVariation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TagVariation) GetTagId() uint32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *TagVariation) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type Variation struct {
	Id          uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Text        string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	Version     uint32 `protobuf:"varint,4,opt,name=version" json:"version,omitempty"`
	LanguageId  uint32 `protobuf:"varint,5,opt,name=languageId" json:"languageId,omitempty"`
	AuthorId    uint32 `protobuf:"varint,6,opt,name=authorId" json:"authorId,omitempty"`
	CopyrightId uint32 `protobuf:"varint,7,opt,name=copyrightId" json:"copyrightId,omitempty"`
	SongId      uint32 `protobuf:"varint,8,opt,name=songId" json:"songId,omitempty"`
}

func (m *Variation) Reset()                    { *m = Variation{} }
func (m *Variation) String() string            { return proto.CompactTextString(m) }
func (*Variation) ProtoMessage()               {}
func (*Variation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Variation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Variation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Variation) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Variation) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Variation) GetLanguageId() uint32 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *Variation) GetAuthorId() uint32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func (m *Variation) GetCopyrightId() uint32 {
	if m != nil {
		return m.CopyrightId
	}
	return 0
}

func (m *Variation) GetSongId() uint32 {
	if m != nil {
		return m.SongId
	}
	return 0
}

type Song struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Song) Reset()                    { *m = Song{} }
func (m *Song) String() string            { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()               {}
func (*Song) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Song) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EventItem struct {
	// Types that are valid to be assigned to EventMessage:
	//	*EventItem_AcceptedKey
	//	*EventItem_NewEwDatabase
	//	*EventItem_NewSongDatabaseVariation
	//	*EventItem_NewSongDatabaseTag
	//	*EventItem_NewTagVariation
	//	*EventItem_NewVariation
	//	*EventItem_NewSong
	EventMessage isEventItem_EventMessage `protobuf_oneof:"eventMessage"`
}

func (m *EventItem) Reset()                    { *m = EventItem{} }
func (m *EventItem) String() string            { return proto.CompactTextString(m) }
func (*EventItem) ProtoMessage()               {}
func (*EventItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isEventItem_EventMessage interface {
	isEventItem_EventMessage()
}

type EventItem_AcceptedKey struct {
	AcceptedKey bool `protobuf:"varint,1,opt,name=acceptedKey,oneof"`
}
type EventItem_NewEwDatabase struct {
	NewEwDatabase *EwDatabase `protobuf:"bytes,2,opt,name=newEwDatabase,oneof"`
}
type EventItem_NewSongDatabaseVariation struct {
	NewSongDatabaseVariation *SongDatabaseVariation `protobuf:"bytes,3,opt,name=newSongDatabaseVariation,oneof"`
}
type EventItem_NewSongDatabaseTag struct {
	NewSongDatabaseTag *SongDatabaseTag `protobuf:"bytes,4,opt,name=newSongDatabaseTag,oneof"`
}
type EventItem_NewTagVariation struct {
	NewTagVariation *TagVariation `protobuf:"bytes,5,opt,name=newTagVariation,oneof"`
}
type EventItem_NewVariation struct {
	NewVariation *Variation `protobuf:"bytes,6,opt,name=newVariation,oneof"`
}
type EventItem_NewSong struct {
	NewSong *Song `protobuf:"bytes,7,opt,name=newSong,oneof"`
}

func (*EventItem_AcceptedKey) isEventItem_EventMessage()              {}
func (*EventItem_NewEwDatabase) isEventItem_EventMessage()            {}
func (*EventItem_NewSongDatabaseVariation) isEventItem_EventMessage() {}
func (*EventItem_NewSongDatabaseTag) isEventItem_EventMessage()       {}
func (*EventItem_NewTagVariation) isEventItem_EventMessage()          {}
func (*EventItem_NewVariation) isEventItem_EventMessage()             {}
func (*EventItem_NewSong) isEventItem_EventMessage()                  {}

func (m *EventItem) GetEventMessage() isEventItem_EventMessage {
	if m != nil {
		return m.EventMessage
	}
	return nil
}

func (m *EventItem) GetAcceptedKey() bool {
	if x, ok := m.GetEventMessage().(*EventItem_AcceptedKey); ok {
		return x.AcceptedKey
	}
	return false
}

func (m *EventItem) GetNewEwDatabase() *EwDatabase {
	if x, ok := m.GetEventMessage().(*EventItem_NewEwDatabase); ok {
		return x.NewEwDatabase
	}
	return nil
}

func (m *EventItem) GetNewSongDatabaseVariation() *SongDatabaseVariation {
	if x, ok := m.GetEventMessage().(*EventItem_NewSongDatabaseVariation); ok {
		return x.NewSongDatabaseVariation
	}
	return nil
}

func (m *EventItem) GetNewSongDatabaseTag() *SongDatabaseTag {
	if x, ok := m.GetEventMessage().(*EventItem_NewSongDatabaseTag); ok {
		return x.NewSongDatabaseTag
	}
	return nil
}

func (m *EventItem) GetNewTagVariation() *TagVariation {
	if x, ok := m.GetEventMessage().(*EventItem_NewTagVariation); ok {
		return x.NewTagVariation
	}
	return nil
}

func (m *EventItem) GetNewVariation() *Variation {
	if x, ok := m.GetEventMessage().(*EventItem_NewVariation); ok {
		return x.NewVariation
	}
	return nil
}

func (m *EventItem) GetNewSong() *Song {
	if x, ok := m.GetEventMessage().(*EventItem_NewSong); ok {
		return x.NewSong
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EventItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EventItem_OneofMarshaler, _EventItem_OneofUnmarshaler, _EventItem_OneofSizer, []interface{}{
		(*EventItem_AcceptedKey)(nil),
		(*EventItem_NewEwDatabase)(nil),
		(*EventItem_NewSongDatabaseVariation)(nil),
		(*EventItem_NewSongDatabaseTag)(nil),
		(*EventItem_NewTagVariation)(nil),
		(*EventItem_NewVariation)(nil),
		(*EventItem_NewSong)(nil),
	}
}

func _EventItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EventItem)
	// eventMessage
	switch x := m.EventMessage.(type) {
	case *EventItem_AcceptedKey:
		t := uint64(0)
		if x.AcceptedKey {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *EventItem_NewEwDatabase:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewEwDatabase); err != nil {
			return err
		}
	case *EventItem_NewSongDatabaseVariation:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewSongDatabaseVariation); err != nil {
			return err
		}
	case *EventItem_NewSongDatabaseTag:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewSongDatabaseTag); err != nil {
			return err
		}
	case *EventItem_NewTagVariation:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewTagVariation); err != nil {
			return err
		}
	case *EventItem_NewVariation:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewVariation); err != nil {
			return err
		}
	case *EventItem_NewSong:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewSong); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EventItem.EventMessage has unexpected type %T", x)
	}
	return nil
}

func _EventItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EventItem)
	switch tag {
	case 1: // eventMessage.acceptedKey
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.EventMessage = &EventItem_AcceptedKey{x != 0}
		return true, err
	case 2: // eventMessage.newEwDatabase
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EwDatabase)
		err := b.DecodeMessage(msg)
		m.EventMessage = &EventItem_NewEwDatabase{msg}
		return true, err
	case 3: // eventMessage.newSongDatabaseVariation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SongDatabaseVariation)
		err := b.DecodeMessage(msg)
		m.EventMessage = &EventItem_NewSongDatabaseVariation{msg}
		return true, err
	case 4: // eventMessage.newSongDatabaseTag
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SongDatabaseTag)
		err := b.DecodeMessage(msg)
		m.EventMessage = &EventItem_NewSongDatabaseTag{msg}
		return true, err
	case 5: // eventMessage.newTagVariation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TagVariation)
		err := b.DecodeMessage(msg)
		m.EventMessage = &EventItem_NewTagVariation{msg}
		return true, err
	case 6: // eventMessage.newVariation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Variation)
		err := b.DecodeMessage(msg)
		m.EventMessage = &EventItem_NewVariation{msg}
		return true, err
	case 7: // eventMessage.newSong
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Song)
		err := b.DecodeMessage(msg)
		m.EventMessage = &EventItem_NewSong{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EventItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EventItem)
	// eventMessage
	switch x := m.EventMessage.(type) {
	case *EventItem_AcceptedKey:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *EventItem_NewEwDatabase:
		s := proto.Size(x.NewEwDatabase)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventItem_NewSongDatabaseVariation:
		s := proto.Size(x.NewSongDatabaseVariation)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventItem_NewSongDatabaseTag:
		s := proto.Size(x.NewSongDatabaseTag)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventItem_NewTagVariation:
		s := proto.Size(x.NewTagVariation)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventItem_NewVariation:
		s := proto.Size(x.NewVariation)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventItem_NewSong:
		s := proto.Size(x.NewSong)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ListenEventsRequest)(nil), "MatiasService.ListenEventsRequest")
	proto.RegisterType((*EwDatabase)(nil), "MatiasService.EwDatabase")
	proto.RegisterType((*SongDatabaseVariation)(nil), "MatiasService.SongDatabaseVariation")
	proto.RegisterType((*SongDatabaseTag)(nil), "MatiasService.SongDatabaseTag")
	proto.RegisterType((*TagVariation)(nil), "MatiasService.TagVariation")
	proto.RegisterType((*Variation)(nil), "MatiasService.Variation")
	proto.RegisterType((*Song)(nil), "MatiasService.Song")
	proto.RegisterType((*EventItem)(nil), "MatiasService.EventItem")
}

func init() { proto.RegisterFile("listen_events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x75, 0x2e, 0xcd, 0x65, 0x9c, 0xa4, 0xb0, 0x81, 0xca, 0x80, 0x54, 0x45, 0x16, 0x42, 0x7d,
	0x40, 0x41, 0x0a, 0xef, 0x48, 0x5c, 0x2a, 0x1c, 0x41, 0x51, 0xb5, 0x8d, 0xfa, 0x5a, 0x4d, 0xe2,
	0xc1, 0xb1, 0x48, 0xd6, 0x89, 0x77, 0x93, 0x90, 0xef, 0xe4, 0x4f, 0xf8, 0x02, 0xe4, 0x69, 0x2e,
	0x8e, 0x71, 0x10, 0xea, 0xdb, 0xce, 0x99, 0x99, 0x73, 0x66, 0x77, 0x8e, 0x0d, 0xed, 0x49, 0xa8,
	0x0d, 0xa9, 0x3b, 0x5a, 0x92, 0x32, 0xba, 0x3b, 0x8b, 0x23, 0x13, 0x89, 0xe6, 0x15, 0x9a, 0x10,
	0xf5, 0x0d, 0xc5, 0xcb, 0x70, 0x44, 0xee, 0x1c, 0xda, 0x5f, 0xb9, 0xea, 0x92, 0x8b, 0x24, 0xcd,
	0x17, 0xa4, 0x8d, 0x78, 0x04, 0xa5, 0x1f, 0xb4, 0x76, 0x0a, 0x9d, 0xc2, 0x45, 0x5d, 0x26, 0x47,
	0xf1, 0x1c, 0x6a, 0xe3, 0x48, 0x9b, 0x6f, 0x38, 0x25, 0xa7, 0xc8, 0xf0, 0x2e, 0x16, 0xaf, 0xe1,
	0xf1, 0x04, 0xb5, 0x21, 0x6d, 0x3e, 0x46, 0x4a, 0xd1, 0xc8, 0x84, 0x91, 0x72, 0x4a, 0x9d, 0xc2,
	0x45, 0x49, 0xfe, 0x9d, 0x70, 0x27, 0x00, 0x97, 0xab, 0x4f, 0x68, 0x70, 0x88, 0x9a, 0x44, 0x0b,
	0x8a, 0xa1, 0xcf, 0x42, 0x4d, 0x59, 0x0c, 0x7d, 0xf1, 0x0a, 0x5a, 0xdf, 0xc3, 0x09, 0xe9, 0xb5,
	0x36, 0x34, 0xbd, 0x46, 0x33, 0xde, 0xa8, 0x65, 0xd0, 0xa4, 0x4e, 0x47, 0x2a, 0xd8, 0xf2, 0xf4,
	0x7d, 0x16, 0x6c, 0xca, 0x0c, 0xea, 0xce, 0xe1, 0xe9, 0x4d, 0x0a, 0xb9, 0xc5, 0x38, 0xc4, 0x64,
	0x8c, 0x3c, 0xe1, 0x0c, 0x61, 0x31, 0x8f, 0x50, 0x74, 0xc0, 0x5e, 0x6e, 0x49, 0x76, 0xaa, 0x69,
	0xc8, 0xbd, 0x83, 0xd3, 0xb4, 0xe4, 0x00, 0x83, 0x07, 0x8b, 0x3d, 0x81, 0x13, 0x83, 0xc1, 0x4e,
	0xe6, 0x3e, 0x70, 0x6f, 0xa1, 0x31, 0xc0, 0xe0, 0xf8, 0x55, 0x76, 0x5d, 0xc5, 0x54, 0xd7, 0x7f,
	0x0c, 0xfe, 0xab, 0x00, 0xf5, 0xe3, 0xac, 0x02, 0xca, 0x6a, 0xbf, 0x7d, 0x3e, 0x27, 0x98, 0xa1,
	0x9f, 0x86, 0xc9, 0xea, 0x92, 0xcf, 0xc2, 0x81, 0xea, 0x92, 0x62, 0x9d, 0x78, 0xa0, 0xcc, 0xcd,
	0xdb, 0x50, 0x9c, 0x03, 0x4c, 0x50, 0x05, 0x0b, 0x0c, 0x92, 0x1b, 0x9f, 0x70, 0x32, 0x85, 0x24,
	0x1e, 0xc3, 0x85, 0x19, 0x47, 0x71, 0xdf, 0x77, 0x2a, 0x9c, 0xdd, 0xc5, 0xc9, 0xf4, 0xa3, 0x68,
	0xb6, 0x8e, 0xc3, 0x60, 0x6c, 0xfa, 0xbe, 0x53, 0xbd, 0x9f, 0x3e, 0x05, 0x89, 0x33, 0xa8, 0x24,
	0xaf, 0xd7, 0xf7, 0x9d, 0x1a, 0x27, 0x37, 0x91, 0x7b, 0x06, 0xe5, 0x64, 0x1d, 0xd9, 0xfb, 0xb8,
	0xbf, 0x4b, 0x50, 0x67, 0xd7, 0xf7, 0x0d, 0x4d, 0x85, 0x0b, 0x36, 0x8e, 0x46, 0x34, 0x33, 0xe4,
	0x7f, 0xd9, 0x38, 0xbf, 0xe6, 0x59, 0x32, 0x0d, 0x8a, 0xf7, 0xd0, 0x54, 0xb4, 0xda, 0x9b, 0x97,
	0x9f, 0xc2, 0xee, 0x3d, 0xeb, 0x1e, 0x7c, 0x53, 0xdd, 0x7d, 0x81, 0x67, 0xc9, 0xc3, 0x0e, 0x31,
	0x04, 0x47, 0xd1, 0x2a, 0xd7, 0x91, 0xfc, 0x88, 0x76, 0xef, 0x65, 0x86, 0x2d, 0xb7, 0xd6, 0xb3,
	0xe4, 0x51, 0x1e, 0x71, 0x0d, 0x22, 0x93, 0x1b, 0x60, 0xc0, 0xbb, 0xb0, 0x7b, 0xe7, 0xff, 0x60,
	0x1f, 0x60, 0xe0, 0x59, 0x32, 0xa7, 0x57, 0x7c, 0x86, 0x53, 0x45, 0xab, 0xb4, 0xe7, 0x78, 0x7b,
	0x76, 0xef, 0x45, 0x86, 0x2e, 0x5d, 0xe2, 0x59, 0x32, 0xdb, 0x25, 0xde, 0x41, 0x43, 0xd1, 0x6a,
	0xcf, 0x52, 0x61, 0x16, 0x27, 0xc3, 0x92, 0xa6, 0x38, 0xa8, 0x17, 0x6f, 0xa0, 0xba, 0x19, 0x8f,
	0x1d, 0x60, 0xf7, 0xda, 0x39, 0xf7, 0xf1, 0x2c, 0xb9, 0xad, 0xfa, 0xd0, 0x82, 0x06, 0xff, 0xfe,
	0xae, 0x48, 0x6b, 0x0c, 0x68, 0x58, 0xe1, 0xbf, 0xe0, 0xdb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd2, 0xbc, 0xfe, 0xe9, 0x1c, 0x05, 0x00, 0x00,
}