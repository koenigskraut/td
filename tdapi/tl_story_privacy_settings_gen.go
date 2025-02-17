// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// StoryPrivacySettingsEveryone represents TL type `storyPrivacySettingsEveryone#a8204874`.
type StoryPrivacySettingsEveryone struct {
}

// StoryPrivacySettingsEveryoneTypeID is TL type id of StoryPrivacySettingsEveryone.
const StoryPrivacySettingsEveryoneTypeID = 0xa8204874

// construct implements constructor of StoryPrivacySettingsClass.
func (s StoryPrivacySettingsEveryone) construct() StoryPrivacySettingsClass { return &s }

// Ensuring interfaces in compile-time for StoryPrivacySettingsEveryone.
var (
	_ bin.Encoder     = &StoryPrivacySettingsEveryone{}
	_ bin.Decoder     = &StoryPrivacySettingsEveryone{}
	_ bin.BareEncoder = &StoryPrivacySettingsEveryone{}
	_ bin.BareDecoder = &StoryPrivacySettingsEveryone{}

	_ StoryPrivacySettingsClass = &StoryPrivacySettingsEveryone{}
)

func (s *StoryPrivacySettingsEveryone) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryPrivacySettingsEveryone) String() string {
	if s == nil {
		return "StoryPrivacySettingsEveryone(nil)"
	}
	type Alias StoryPrivacySettingsEveryone
	return fmt.Sprintf("StoryPrivacySettingsEveryone%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryPrivacySettingsEveryone) TypeID() uint32 {
	return StoryPrivacySettingsEveryoneTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryPrivacySettingsEveryone) TypeName() string {
	return "storyPrivacySettingsEveryone"
}

// TypeInfo returns info about TL type.
func (s *StoryPrivacySettingsEveryone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyPrivacySettingsEveryone",
		ID:   StoryPrivacySettingsEveryoneTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryPrivacySettingsEveryone) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsEveryone#a8204874 as nil")
	}
	b.PutID(StoryPrivacySettingsEveryoneTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryPrivacySettingsEveryone) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsEveryone#a8204874 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryPrivacySettingsEveryone) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsEveryone#a8204874 to nil")
	}
	if err := b.ConsumeID(StoryPrivacySettingsEveryoneTypeID); err != nil {
		return fmt.Errorf("unable to decode storyPrivacySettingsEveryone#a8204874: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryPrivacySettingsEveryone) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsEveryone#a8204874 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryPrivacySettingsEveryone) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsEveryone#a8204874 as nil")
	}
	b.ObjStart()
	b.PutID("storyPrivacySettingsEveryone")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryPrivacySettingsEveryone) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsEveryone#a8204874 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyPrivacySettingsEveryone"); err != nil {
				return fmt.Errorf("unable to decode storyPrivacySettingsEveryone#a8204874: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StoryPrivacySettingsContacts represents TL type `storyPrivacySettingsContacts#35515d71`.
type StoryPrivacySettingsContacts struct {
	// User identifiers of the contacts that can't see the story; always empty for non-owned
	// stories
	ExceptUserIDs []int64
}

// StoryPrivacySettingsContactsTypeID is TL type id of StoryPrivacySettingsContacts.
const StoryPrivacySettingsContactsTypeID = 0x35515d71

// construct implements constructor of StoryPrivacySettingsClass.
func (s StoryPrivacySettingsContacts) construct() StoryPrivacySettingsClass { return &s }

// Ensuring interfaces in compile-time for StoryPrivacySettingsContacts.
var (
	_ bin.Encoder     = &StoryPrivacySettingsContacts{}
	_ bin.Decoder     = &StoryPrivacySettingsContacts{}
	_ bin.BareEncoder = &StoryPrivacySettingsContacts{}
	_ bin.BareDecoder = &StoryPrivacySettingsContacts{}

	_ StoryPrivacySettingsClass = &StoryPrivacySettingsContacts{}
)

func (s *StoryPrivacySettingsContacts) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ExceptUserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryPrivacySettingsContacts) String() string {
	if s == nil {
		return "StoryPrivacySettingsContacts(nil)"
	}
	type Alias StoryPrivacySettingsContacts
	return fmt.Sprintf("StoryPrivacySettingsContacts%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryPrivacySettingsContacts) TypeID() uint32 {
	return StoryPrivacySettingsContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryPrivacySettingsContacts) TypeName() string {
	return "storyPrivacySettingsContacts"
}

// TypeInfo returns info about TL type.
func (s *StoryPrivacySettingsContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyPrivacySettingsContacts",
		ID:   StoryPrivacySettingsContactsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ExceptUserIDs",
			SchemaName: "except_user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryPrivacySettingsContacts) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsContacts#35515d71 as nil")
	}
	b.PutID(StoryPrivacySettingsContactsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryPrivacySettingsContacts) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsContacts#35515d71 as nil")
	}
	b.PutInt(len(s.ExceptUserIDs))
	for _, v := range s.ExceptUserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryPrivacySettingsContacts) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsContacts#35515d71 to nil")
	}
	if err := b.ConsumeID(StoryPrivacySettingsContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode storyPrivacySettingsContacts#35515d71: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryPrivacySettingsContacts) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsContacts#35515d71 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyPrivacySettingsContacts#35515d71: field except_user_ids: %w", err)
		}

		if headerLen > 0 {
			s.ExceptUserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storyPrivacySettingsContacts#35515d71: field except_user_ids: %w", err)
			}
			s.ExceptUserIDs = append(s.ExceptUserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryPrivacySettingsContacts) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsContacts#35515d71 as nil")
	}
	b.ObjStart()
	b.PutID("storyPrivacySettingsContacts")
	b.Comma()
	b.FieldStart("except_user_ids")
	b.ArrStart()
	for _, v := range s.ExceptUserIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryPrivacySettingsContacts) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsContacts#35515d71 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyPrivacySettingsContacts"); err != nil {
				return fmt.Errorf("unable to decode storyPrivacySettingsContacts#35515d71: %w", err)
			}
		case "except_user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode storyPrivacySettingsContacts#35515d71: field except_user_ids: %w", err)
				}
				s.ExceptUserIDs = append(s.ExceptUserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode storyPrivacySettingsContacts#35515d71: field except_user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetExceptUserIDs returns value of ExceptUserIDs field.
func (s *StoryPrivacySettingsContacts) GetExceptUserIDs() (value []int64) {
	if s == nil {
		return
	}
	return s.ExceptUserIDs
}

// StoryPrivacySettingsCloseFriends represents TL type `storyPrivacySettingsCloseFriends#7cff8b60`.
type StoryPrivacySettingsCloseFriends struct {
}

// StoryPrivacySettingsCloseFriendsTypeID is TL type id of StoryPrivacySettingsCloseFriends.
const StoryPrivacySettingsCloseFriendsTypeID = 0x7cff8b60

// construct implements constructor of StoryPrivacySettingsClass.
func (s StoryPrivacySettingsCloseFriends) construct() StoryPrivacySettingsClass { return &s }

// Ensuring interfaces in compile-time for StoryPrivacySettingsCloseFriends.
var (
	_ bin.Encoder     = &StoryPrivacySettingsCloseFriends{}
	_ bin.Decoder     = &StoryPrivacySettingsCloseFriends{}
	_ bin.BareEncoder = &StoryPrivacySettingsCloseFriends{}
	_ bin.BareDecoder = &StoryPrivacySettingsCloseFriends{}

	_ StoryPrivacySettingsClass = &StoryPrivacySettingsCloseFriends{}
)

func (s *StoryPrivacySettingsCloseFriends) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryPrivacySettingsCloseFriends) String() string {
	if s == nil {
		return "StoryPrivacySettingsCloseFriends(nil)"
	}
	type Alias StoryPrivacySettingsCloseFriends
	return fmt.Sprintf("StoryPrivacySettingsCloseFriends%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryPrivacySettingsCloseFriends) TypeID() uint32 {
	return StoryPrivacySettingsCloseFriendsTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryPrivacySettingsCloseFriends) TypeName() string {
	return "storyPrivacySettingsCloseFriends"
}

// TypeInfo returns info about TL type.
func (s *StoryPrivacySettingsCloseFriends) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyPrivacySettingsCloseFriends",
		ID:   StoryPrivacySettingsCloseFriendsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryPrivacySettingsCloseFriends) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsCloseFriends#7cff8b60 as nil")
	}
	b.PutID(StoryPrivacySettingsCloseFriendsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryPrivacySettingsCloseFriends) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsCloseFriends#7cff8b60 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryPrivacySettingsCloseFriends) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsCloseFriends#7cff8b60 to nil")
	}
	if err := b.ConsumeID(StoryPrivacySettingsCloseFriendsTypeID); err != nil {
		return fmt.Errorf("unable to decode storyPrivacySettingsCloseFriends#7cff8b60: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryPrivacySettingsCloseFriends) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsCloseFriends#7cff8b60 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryPrivacySettingsCloseFriends) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsCloseFriends#7cff8b60 as nil")
	}
	b.ObjStart()
	b.PutID("storyPrivacySettingsCloseFriends")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryPrivacySettingsCloseFriends) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsCloseFriends#7cff8b60 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyPrivacySettingsCloseFriends"); err != nil {
				return fmt.Errorf("unable to decode storyPrivacySettingsCloseFriends#7cff8b60: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StoryPrivacySettingsSelectedContacts represents TL type `storyPrivacySettingsSelectedContacts#25401fc7`.
type StoryPrivacySettingsSelectedContacts struct {
	// Identifiers of the users; always empty for non-owned stories
	UserIDs []int64
}

// StoryPrivacySettingsSelectedContactsTypeID is TL type id of StoryPrivacySettingsSelectedContacts.
const StoryPrivacySettingsSelectedContactsTypeID = 0x25401fc7

// construct implements constructor of StoryPrivacySettingsClass.
func (s StoryPrivacySettingsSelectedContacts) construct() StoryPrivacySettingsClass { return &s }

// Ensuring interfaces in compile-time for StoryPrivacySettingsSelectedContacts.
var (
	_ bin.Encoder     = &StoryPrivacySettingsSelectedContacts{}
	_ bin.Decoder     = &StoryPrivacySettingsSelectedContacts{}
	_ bin.BareEncoder = &StoryPrivacySettingsSelectedContacts{}
	_ bin.BareDecoder = &StoryPrivacySettingsSelectedContacts{}

	_ StoryPrivacySettingsClass = &StoryPrivacySettingsSelectedContacts{}
)

func (s *StoryPrivacySettingsSelectedContacts) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryPrivacySettingsSelectedContacts) String() string {
	if s == nil {
		return "StoryPrivacySettingsSelectedContacts(nil)"
	}
	type Alias StoryPrivacySettingsSelectedContacts
	return fmt.Sprintf("StoryPrivacySettingsSelectedContacts%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryPrivacySettingsSelectedContacts) TypeID() uint32 {
	return StoryPrivacySettingsSelectedContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryPrivacySettingsSelectedContacts) TypeName() string {
	return "storyPrivacySettingsSelectedContacts"
}

// TypeInfo returns info about TL type.
func (s *StoryPrivacySettingsSelectedContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyPrivacySettingsSelectedContacts",
		ID:   StoryPrivacySettingsSelectedContactsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryPrivacySettingsSelectedContacts) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsSelectedContacts#25401fc7 as nil")
	}
	b.PutID(StoryPrivacySettingsSelectedContactsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryPrivacySettingsSelectedContacts) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsSelectedContacts#25401fc7 as nil")
	}
	b.PutInt(len(s.UserIDs))
	for _, v := range s.UserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryPrivacySettingsSelectedContacts) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsSelectedContacts#25401fc7 to nil")
	}
	if err := b.ConsumeID(StoryPrivacySettingsSelectedContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode storyPrivacySettingsSelectedContacts#25401fc7: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryPrivacySettingsSelectedContacts) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsSelectedContacts#25401fc7 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyPrivacySettingsSelectedContacts#25401fc7: field user_ids: %w", err)
		}

		if headerLen > 0 {
			s.UserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storyPrivacySettingsSelectedContacts#25401fc7: field user_ids: %w", err)
			}
			s.UserIDs = append(s.UserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryPrivacySettingsSelectedContacts) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyPrivacySettingsSelectedContacts#25401fc7 as nil")
	}
	b.ObjStart()
	b.PutID("storyPrivacySettingsSelectedContacts")
	b.Comma()
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range s.UserIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryPrivacySettingsSelectedContacts) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyPrivacySettingsSelectedContacts#25401fc7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyPrivacySettingsSelectedContacts"); err != nil {
				return fmt.Errorf("unable to decode storyPrivacySettingsSelectedContacts#25401fc7: %w", err)
			}
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode storyPrivacySettingsSelectedContacts#25401fc7: field user_ids: %w", err)
				}
				s.UserIDs = append(s.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode storyPrivacySettingsSelectedContacts#25401fc7: field user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserIDs returns value of UserIDs field.
func (s *StoryPrivacySettingsSelectedContacts) GetUserIDs() (value []int64) {
	if s == nil {
		return
	}
	return s.UserIDs
}

// StoryPrivacySettingsClassName is schema name of StoryPrivacySettingsClass.
const StoryPrivacySettingsClassName = "StoryPrivacySettings"

// StoryPrivacySettingsClass represents StoryPrivacySettings generic type.
//
// Example:
//
//	g, err := tdapi.DecodeStoryPrivacySettings(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.StoryPrivacySettingsEveryone: // storyPrivacySettingsEveryone#a8204874
//	case *tdapi.StoryPrivacySettingsContacts: // storyPrivacySettingsContacts#35515d71
//	case *tdapi.StoryPrivacySettingsCloseFriends: // storyPrivacySettingsCloseFriends#7cff8b60
//	case *tdapi.StoryPrivacySettingsSelectedContacts: // storyPrivacySettingsSelectedContacts#25401fc7
//	default: panic(v)
//	}
type StoryPrivacySettingsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoryPrivacySettingsClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeStoryPrivacySettings implements binary de-serialization for StoryPrivacySettingsClass.
func DecodeStoryPrivacySettings(buf *bin.Buffer) (StoryPrivacySettingsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoryPrivacySettingsEveryoneTypeID:
		// Decoding storyPrivacySettingsEveryone#a8204874.
		v := StoryPrivacySettingsEveryone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", err)
		}
		return &v, nil
	case StoryPrivacySettingsContactsTypeID:
		// Decoding storyPrivacySettingsContacts#35515d71.
		v := StoryPrivacySettingsContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", err)
		}
		return &v, nil
	case StoryPrivacySettingsCloseFriendsTypeID:
		// Decoding storyPrivacySettingsCloseFriends#7cff8b60.
		v := StoryPrivacySettingsCloseFriends{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", err)
		}
		return &v, nil
	case StoryPrivacySettingsSelectedContactsTypeID:
		// Decoding storyPrivacySettingsSelectedContacts#25401fc7.
		v := StoryPrivacySettingsSelectedContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONStoryPrivacySettings implements binary de-serialization for StoryPrivacySettingsClass.
func DecodeTDLibJSONStoryPrivacySettings(buf tdjson.Decoder) (StoryPrivacySettingsClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "storyPrivacySettingsEveryone":
		// Decoding storyPrivacySettingsEveryone#a8204874.
		v := StoryPrivacySettingsEveryone{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", err)
		}
		return &v, nil
	case "storyPrivacySettingsContacts":
		// Decoding storyPrivacySettingsContacts#35515d71.
		v := StoryPrivacySettingsContacts{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", err)
		}
		return &v, nil
	case "storyPrivacySettingsCloseFriends":
		// Decoding storyPrivacySettingsCloseFriends#7cff8b60.
		v := StoryPrivacySettingsCloseFriends{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", err)
		}
		return &v, nil
	case "storyPrivacySettingsSelectedContacts":
		// Decoding storyPrivacySettingsSelectedContacts#25401fc7.
		v := StoryPrivacySettingsSelectedContacts{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryPrivacySettingsClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// StoryPrivacySettings boxes the StoryPrivacySettingsClass providing a helper.
type StoryPrivacySettingsBox struct {
	StoryPrivacySettings StoryPrivacySettingsClass
}

// Decode implements bin.Decoder for StoryPrivacySettingsBox.
func (b *StoryPrivacySettingsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryPrivacySettingsBox to nil")
	}
	v, err := DecodeStoryPrivacySettings(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryPrivacySettings = v
	return nil
}

// Encode implements bin.Encode for StoryPrivacySettingsBox.
func (b *StoryPrivacySettingsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StoryPrivacySettings == nil {
		return fmt.Errorf("unable to encode StoryPrivacySettingsClass as nil")
	}
	return b.StoryPrivacySettings.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for StoryPrivacySettingsBox.
func (b *StoryPrivacySettingsBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryPrivacySettingsBox to nil")
	}
	v, err := DecodeTDLibJSONStoryPrivacySettings(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryPrivacySettings = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for StoryPrivacySettingsBox.
func (b *StoryPrivacySettingsBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.StoryPrivacySettings == nil {
		return fmt.Errorf("unable to encode StoryPrivacySettingsClass as nil")
	}
	return b.StoryPrivacySettings.EncodeTDLibJSON(buf)
}
