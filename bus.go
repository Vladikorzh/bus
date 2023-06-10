package bus

import (
	"context"
	"errors"
	"reflect"
)

var (
	ErrHandlerNotFound    = errors.New("handler not found")
	ErrInvalidMessageType = errors.New("invalid message type")
)

type Message any

type Handler interface {
	handle(ctx context.Context, msg Message) error
}

type HandlerFunc[T any] func(ctx context.Context, msg *T) error

func (fn HandlerFunc[T]) handle(ctx context.Context, msg Message) error {
	switch m := msg.(type) {
	case *T:
		return fn(ctx, m)
	default:
		return ErrInvalidMessageType
	}
}

type Bus interface {
	Handle(handler Handler)
	Publish(ctx context.Context, msg Message) error
}

type bus struct {
	handlers map[string]Handler
}

func New() Bus {
	return &bus{
		handlers: map[string]Handler{},
	}
}

func (b *bus) Handle(handler Handler) {
	b.handlers[reflect.TypeOf(handler).In(1).Elem().Name()] = handler
}

func (b *bus) Publish(ctx context.Context, msg Message) error {
	handler, ok := b.handlers[reflect.TypeOf(msg).Elem().Name()]
	if ok {
		return handler.handle(ctx, msg)
	}

	return ErrHandlerNotFound
}
