# Bus

```go
import "github.com/Vladikorzh/bus"
```

Example:
```go
import "github.com/Vladikorzh/bus"

type Ping struct {
    Result string
}

func main()  {
    msg := Ping{}
    
    bus.Handle(Ping{}, func(ctx context.Context, ping *Ping) {
        ping.Result = "pong"
    })
    
    _ = bus.Publish(context.Background(), &msg)
    
    fmt.Println(msg.Result)
}
```