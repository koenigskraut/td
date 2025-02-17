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

// DeleteStoryRequest represents TL type `deleteStory#62332a06`.
type DeleteStoryRequest struct {
	// Identifier of the story to delete
	StoryID int32
}

// DeleteStoryRequestTypeID is TL type id of DeleteStoryRequest.
const DeleteStoryRequestTypeID = 0x62332a06

// Ensuring interfaces in compile-time for DeleteStoryRequest.
var (
	_ bin.Encoder     = &DeleteStoryRequest{}
	_ bin.Decoder     = &DeleteStoryRequest{}
	_ bin.BareEncoder = &DeleteStoryRequest{}
	_ bin.BareDecoder = &DeleteStoryRequest{}
)

func (d *DeleteStoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.StoryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteStoryRequest) String() string {
	if d == nil {
		return "DeleteStoryRequest(nil)"
	}
	type Alias DeleteStoryRequest
	return fmt.Sprintf("DeleteStoryRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteStoryRequest) TypeID() uint32 {
	return DeleteStoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteStoryRequest) TypeName() string {
	return "deleteStory"
}

// TypeInfo returns info about TL type.
func (d *DeleteStoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteStory",
		ID:   DeleteStoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteStoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteStory#62332a06 as nil")
	}
	b.PutID(DeleteStoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteStoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteStory#62332a06 as nil")
	}
	b.PutInt32(d.StoryID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteStoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteStory#62332a06 to nil")
	}
	if err := b.ConsumeID(DeleteStoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteStory#62332a06: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteStoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteStory#62332a06 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode deleteStory#62332a06: field story_id: %w", err)
		}
		d.StoryID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteStoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteStory#62332a06 as nil")
	}
	b.ObjStart()
	b.PutID("deleteStory")
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(d.StoryID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteStoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteStory#62332a06 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteStory"); err != nil {
				return fmt.Errorf("unable to decode deleteStory#62332a06: %w", err)
			}
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode deleteStory#62332a06: field story_id: %w", err)
			}
			d.StoryID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStoryID returns value of StoryID field.
func (d *DeleteStoryRequest) GetStoryID() (value int32) {
	if d == nil {
		return
	}
	return d.StoryID
}

// DeleteStory invokes method deleteStory#62332a06 returning error if any.
func (c *Client) DeleteStory(ctx context.Context, storyid int32) error {
	var ok Ok

	request := &DeleteStoryRequest{
		StoryID: storyid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
