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

// GetArchivedStoriesRequest represents TL type `getArchivedStories#6a864fe7`.
type GetArchivedStoriesRequest struct {
	// Identifier of the story starting from which stories must be returned; use 0 to get
	// results from the last story
	FromStoryID int32
	// The maximum number of stories to be returned
	Limit int32
}

// GetArchivedStoriesRequestTypeID is TL type id of GetArchivedStoriesRequest.
const GetArchivedStoriesRequestTypeID = 0x6a864fe7

// Ensuring interfaces in compile-time for GetArchivedStoriesRequest.
var (
	_ bin.Encoder     = &GetArchivedStoriesRequest{}
	_ bin.Decoder     = &GetArchivedStoriesRequest{}
	_ bin.BareEncoder = &GetArchivedStoriesRequest{}
	_ bin.BareDecoder = &GetArchivedStoriesRequest{}
)

func (g *GetArchivedStoriesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.FromStoryID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetArchivedStoriesRequest) String() string {
	if g == nil {
		return "GetArchivedStoriesRequest(nil)"
	}
	type Alias GetArchivedStoriesRequest
	return fmt.Sprintf("GetArchivedStoriesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetArchivedStoriesRequest) TypeID() uint32 {
	return GetArchivedStoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetArchivedStoriesRequest) TypeName() string {
	return "getArchivedStories"
}

// TypeInfo returns info about TL type.
func (g *GetArchivedStoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getArchivedStories",
		ID:   GetArchivedStoriesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FromStoryID",
			SchemaName: "from_story_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetArchivedStoriesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getArchivedStories#6a864fe7 as nil")
	}
	b.PutID(GetArchivedStoriesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetArchivedStoriesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getArchivedStories#6a864fe7 as nil")
	}
	b.PutInt32(g.FromStoryID)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetArchivedStoriesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getArchivedStories#6a864fe7 to nil")
	}
	if err := b.ConsumeID(GetArchivedStoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getArchivedStories#6a864fe7: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetArchivedStoriesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getArchivedStories#6a864fe7 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getArchivedStories#6a864fe7: field from_story_id: %w", err)
		}
		g.FromStoryID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getArchivedStories#6a864fe7: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetArchivedStoriesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getArchivedStories#6a864fe7 as nil")
	}
	b.ObjStart()
	b.PutID("getArchivedStories")
	b.Comma()
	b.FieldStart("from_story_id")
	b.PutInt32(g.FromStoryID)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetArchivedStoriesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getArchivedStories#6a864fe7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getArchivedStories"); err != nil {
				return fmt.Errorf("unable to decode getArchivedStories#6a864fe7: %w", err)
			}
		case "from_story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getArchivedStories#6a864fe7: field from_story_id: %w", err)
			}
			g.FromStoryID = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getArchivedStories#6a864fe7: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFromStoryID returns value of FromStoryID field.
func (g *GetArchivedStoriesRequest) GetFromStoryID() (value int32) {
	if g == nil {
		return
	}
	return g.FromStoryID
}

// GetLimit returns value of Limit field.
func (g *GetArchivedStoriesRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetArchivedStories invokes method getArchivedStories#6a864fe7 returning error if any.
func (c *Client) GetArchivedStories(ctx context.Context, request *GetArchivedStoriesRequest) (*Stories, error) {
	var result Stories

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
