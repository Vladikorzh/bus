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

var instance Bus = New()

type Message any

type Handler func(ctx context.Context, msg Message) error

type Bus interface {
	Handle(msg any, handler Handler)
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

func (b *bus) Handle(msg any, handler Handler) {
	b.handlers[reflect.TypeOf(msg).Name()] = handler
}

func (b *bus) Publish(ctx context.Context, msg Message) error {
	handler, ok := b.handlers[reflect.TypeOf(msg).Elem().Name()]
	if ok {
		return handler(ctx, msg)
	}

	return ErrHandlerNotFound
}

func GenericHandler[T any](handler func(ctx context.Context, msg *T)) Handler {
	return func(ctx context.Context, msg Message) error {
		switch m := msg.(type) {
		case *T:
			handler(ctx, m)
		default:
			return ErrInvalidMessageType
		}

		return nil
	}
}

func Handle[T any](msg T, handler func(ctx context.Context, msg *T)) {
	instance.Handle(msg, GenericHandler(handler))
}

func Publish(ctx context.Context, msg Message) error {
	return instance.Publish(ctx, msg)
}
