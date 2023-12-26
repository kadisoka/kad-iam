package iam

import (
	"context"
	"time"

	"github.com/alloyzeus/go-azfl/v2/errors"
)

var (
	ErrCallMetadataContextTypeInvalid = errors.CtxM("context call metadata type invalid")
	ErrCallMetadataContextNotFound    = errors.CtxM("context has no call metadata")
)

const CallMetadataKey = contextKey("CallMetadata")

// CallMetadata holds metadata about a service method call.
type CallMetadata struct {
	contextErr    error
	authorization *Authorization

	// ReceiveTime returns the time when request was accepted by
	// the handler.
	receiveTime time.Time
}

func (md CallMetadata) ContextError() error { return md.contextErr }

func (md CallMetadata) Authorization() *Authorization {
	return md.authorization
}

func (md CallMetadata) ReceiveTime() time.Time { return md.receiveTime }

// CallMetadataOf extracts CallMetadata from ctx.
//
// It returns valid instance of CallMetadata if it exists. It returns nil
// if no instance of CallMetadata existed or it exists but nil. It
// returns an instance of CallMetadata which ContextError returns an error
// if the value from the ctx is of unexpected type.
func CallMetadataOf(ctx context.Context) *CallMetadata {
	val := ctx.Value(CallMetadataKey)
	if val == nil {
		return nil
	}
	if md, ok := val.(CallMetadata); ok {
		return &md
	}
	if md, ok := val.(*CallMetadata); ok {
		return md
	}
	return &CallMetadata{contextErr: ErrCallMetadataContextTypeInvalid}
}
