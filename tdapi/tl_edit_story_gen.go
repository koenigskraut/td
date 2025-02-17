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

// EditStoryRequest represents TL type `editStory#ca87df56`.
type EditStoryRequest struct {
	// Identifier of the story to edit
	StoryID int32
	// New content of the story; pass null to keep the current content
	Content InputStoryContentClass
	// New story caption; pass null to keep the current caption
	Caption FormattedText
}

// EditStoryRequestTypeID is TL type id of EditStoryRequest.
const EditStoryRequestTypeID = 0xca87df56

// Ensuring interfaces in compile-time for EditStoryRequest.
var (
	_ bin.Encoder     = &EditStoryRequest{}
	_ bin.Decoder     = &EditStoryRequest{}
	_ bin.BareEncoder = &EditStoryRequest{}
	_ bin.BareDecoder = &EditStoryRequest{}
)

func (e *EditStoryRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.StoryID == 0) {
		return false
	}
	if !(e.Content == nil) {
		return false
	}
	if !(e.Caption.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditStoryRequest) String() string {
	if e == nil {
		return "EditStoryRequest(nil)"
	}
	type Alias EditStoryRequest
	return fmt.Sprintf("EditStoryRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditStoryRequest) TypeID() uint32 {
	return EditStoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditStoryRequest) TypeName() string {
	return "editStory"
}

// TypeInfo returns info about TL type.
func (e *EditStoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editStory",
		ID:   EditStoryRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
		{
			Name:       "Content",
			SchemaName: "content",
		},
		{
			Name:       "Caption",
			SchemaName: "caption",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditStoryRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editStory#ca87df56 as nil")
	}
	b.PutID(EditStoryRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditStoryRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editStory#ca87df56 as nil")
	}
	b.PutInt32(e.StoryID)
	if e.Content == nil {
		return fmt.Errorf("unable to encode editStory#ca87df56: field content is nil")
	}
	if err := e.Content.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editStory#ca87df56: field content: %w", err)
	}
	if err := e.Caption.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editStory#ca87df56: field caption: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditStoryRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editStory#ca87df56 to nil")
	}
	if err := b.ConsumeID(EditStoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editStory#ca87df56: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditStoryRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editStory#ca87df56 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode editStory#ca87df56: field story_id: %w", err)
		}
		e.StoryID = value
	}
	{
		value, err := DecodeInputStoryContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode editStory#ca87df56: field content: %w", err)
		}
		e.Content = value
	}
	{
		if err := e.Caption.Decode(b); err != nil {
			return fmt.Errorf("unable to decode editStory#ca87df56: field caption: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditStoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editStory#ca87df56 as nil")
	}
	b.ObjStart()
	b.PutID("editStory")
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(e.StoryID)
	b.Comma()
	b.FieldStart("content")
	if e.Content == nil {
		return fmt.Errorf("unable to encode editStory#ca87df56: field content is nil")
	}
	if err := e.Content.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editStory#ca87df56: field content: %w", err)
	}
	b.Comma()
	b.FieldStart("caption")
	if err := e.Caption.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editStory#ca87df56: field caption: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditStoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editStory#ca87df56 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editStory"); err != nil {
				return fmt.Errorf("unable to decode editStory#ca87df56: %w", err)
			}
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode editStory#ca87df56: field story_id: %w", err)
			}
			e.StoryID = value
		case "content":
			value, err := DecodeTDLibJSONInputStoryContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode editStory#ca87df56: field content: %w", err)
			}
			e.Content = value
		case "caption":
			if err := e.Caption.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode editStory#ca87df56: field caption: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStoryID returns value of StoryID field.
func (e *EditStoryRequest) GetStoryID() (value int32) {
	if e == nil {
		return
	}
	return e.StoryID
}

// GetContent returns value of Content field.
func (e *EditStoryRequest) GetContent() (value InputStoryContentClass) {
	if e == nil {
		return
	}
	return e.Content
}

// GetCaption returns value of Caption field.
func (e *EditStoryRequest) GetCaption() (value FormattedText) {
	if e == nil {
		return
	}
	return e.Caption
}

// EditStory invokes method editStory#ca87df56 returning error if any.
func (c *Client) EditStory(ctx context.Context, request *EditStoryRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
