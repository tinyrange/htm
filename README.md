# HTM

## WARNING

This is a very early alpha not ready for production use.

## Sample Code

See [examples/hello/main.go](examples/hello/main.go) for a full sample.

```go
htm.ListenAndServe("127.0.0.1:1512", h.Html(
    h.Head(
        h.Title("Hello, World"),
    ),
    h.Body(
        h.Div(htm.Text("Hello, World")),
    ),
))
```
