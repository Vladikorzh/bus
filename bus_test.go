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

	bus.Handle(Ping{}, func(ctx context.Context, ping *Ping) {
		ping.Result = "pong"
	})

	msg := Ping{}

	if err := bus.Publish(context.Background(), &msg); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, msg.Result, "pong", "unexpected result")
}

func BenchmarkBus_Publish(b *testing.B) {
	type Ping struct {
		Result string
	}

	bus.Handle(Ping{}, func(ctx context.Context, ping *Ping) {
		ping.Result = "pong"
	})

	for i := 0; i < b.N; i++ {
		msg := Ping{}

		if err := bus.Publish(context.Background(), &msg); err != nil {
			b.Fatal(err)
		}
	}
}
