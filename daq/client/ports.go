package client

import (
    "context"
    "io"
    "time"

    "github.com/fanie42/lemi011b"
)

type request struct {
    timestamp time.Time
    bytes     []byte
}

type iLogger interface {
    Error(string, error)
    Info(string)
}

// type iPublisher interface {
//     Publish(*lemi011b.Datum)
//     Close() error
// }

type iReader interface {
    io.Reader
}

type iRepository interface {
    Create(context.Context, *lemi011b.Datum) error
}

type iPublisher interface {
    PublishDatum(context.Context, *lemi011b.Datum)
    PublishError(context.Context, error)
}
