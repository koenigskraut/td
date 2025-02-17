// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// StoriesExportStoryLinkRequest represents TL type `stories.exportStoryLink#16e443ce`.
//
// See https://core.telegram.org/method/stories.exportStoryLink for reference.
type StoriesExportStoryLinkRequest struct {
	// UserID field of StoriesExportStoryLinkRequest.
	UserID InputUserClass
	// ID field of StoriesExportStoryLinkRequest.
	ID int
}

// StoriesExportStoryLinkRequestTypeID is TL type id of StoriesExportStoryLinkRequest.
const StoriesExportStoryLinkRequestTypeID = 0x16e443ce

// Ensuring interfaces in compile-time for StoriesExportStoryLinkRequest.
var (
	_ bin.Encoder     = &StoriesExportStoryLinkRequest{}
	_ bin.Decoder     = &StoriesExportStoryLinkRequest{}
	_ bin.BareEncoder = &StoriesExportStoryLinkRequest{}
	_ bin.BareDecoder = &StoriesExportStoryLinkRequest{}
)

func (e *StoriesExportStoryLinkRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.UserID == nil) {
		return false
	}
	if !(e.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *StoriesExportStoryLinkRequest) String() string {
	if e == nil {
		return "StoriesExportStoryLinkRequest(nil)"
	}
	type Alias StoriesExportStoryLinkRequest
	return fmt.Sprintf("StoriesExportStoryLinkRequest%+v", Alias(*e))
}

// FillFrom fills StoriesExportStoryLinkRequest from given interface.
func (e *StoriesExportStoryLinkRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetID() (value int)
}) {
	e.UserID = from.GetUserID()
	e.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesExportStoryLinkRequest) TypeID() uint32 {
	return StoriesExportStoryLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesExportStoryLinkRequest) TypeName() string {
	return "stories.exportStoryLink"
}

// TypeInfo returns info about TL type.
func (e *StoriesExportStoryLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.exportStoryLink",
		ID:   StoriesExportStoryLinkRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *StoriesExportStoryLinkRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode stories.exportStoryLink#16e443ce as nil")
	}
	b.PutID(StoriesExportStoryLinkRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *StoriesExportStoryLinkRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode stories.exportStoryLink#16e443ce as nil")
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode stories.exportStoryLink#16e443ce: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.exportStoryLink#16e443ce: field user_id: %w", err)
	}
	b.PutInt(e.ID)
	return nil
}

// Decode implements bin.Decoder.
func (e *StoriesExportStoryLinkRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode stories.exportStoryLink#16e443ce to nil")
	}
	if err := b.ConsumeID(StoriesExportStoryLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.exportStoryLink#16e443ce: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *StoriesExportStoryLinkRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode stories.exportStoryLink#16e443ce to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.exportStoryLink#16e443ce: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.exportStoryLink#16e443ce: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (e *StoriesExportStoryLinkRequest) GetUserID() (value InputUserClass) {
	if e == nil {
		return
	}
	return e.UserID
}

// GetID returns value of ID field.
func (e *StoriesExportStoryLinkRequest) GetID() (value int) {
	if e == nil {
		return
	}
	return e.ID
}

// StoriesExportStoryLink invokes method stories.exportStoryLink#16e443ce returning error if any.
//
// See https://core.telegram.org/method/stories.exportStoryLink for reference.
func (c *Client) StoriesExportStoryLink(ctx context.Context, request *StoriesExportStoryLinkRequest) (*ExportedStoryLink, error) {
	var result ExportedStoryLink

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
