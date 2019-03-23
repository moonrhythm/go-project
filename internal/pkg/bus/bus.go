package bus

import (
	"context"

	"github.com/moonrhythm/dispatcher"
)

var mux = dispatcher.NewMux()

// Alias types
type (
	Handler = dispatcher.Handler
	Message = dispatcher.Message
)

// Register registers handler into bus
func Register(hs ...Handler) {
	mux.Register(hs...)
}

// Dispatch dispatches messages
func Dispatch(ctx context.Context, msg ...Message) error {
	return dispatcher.Dispatch(ctx, mux, msg...)
}
