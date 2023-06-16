package bus_test

import (
	"bus"
	"context"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBus_Publish(t *testing.T) {
	type Ping struct {
		Result string
	}

	messages := bus.New()

	messages.Handle(bus.Func(func(ctx context.Context, msg *Ping) error {
		msg.Result = "pong"

		return nil
	}))

	msg := Ping{}

	if err := messages.Publish(context.Background(), &msg); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, msg.Result, "pong", "unexpected result")
}

func BenchmarkBus_Publish(b *testing.B) {
	type Ping struct {
		Result string
	}

	messages := bus.New()

	messages.Handle(bus.Func(func(ctx context.Context, msg *Ping) error {
		msg.Result = "pong"

		return nil
	}))

	for i := 0; i < b.N; i++ {
		msg := Ping{}

		if err := messages.Publish(context.Background(), &msg); err != nil {
			b.Fatal(err)
		}
	}
}
