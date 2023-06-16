# Bus

Example:
```go
package main

import (
	"context"
	"fmt"
	"github.com/Vladikorzh/bus"
)

type Ping struct {
    Result string
}

func main() {
    var msg Ping

	messages := bus.New()

	messages.Handle(bus.Func(func(ctx context.Context, msg *Ping) error {
		msg.Result = "pong"

		return nil
	}))

    _ = messages.Publish(context.Background(), &msg)

    fmt.Println(msg.Result)
}

```