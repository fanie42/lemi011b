package server

import (
    "context"

    "github.com/fanie42/lemi011b"
)

type iLogger interface {
    Error(string, error)
    Info(string)
}

type iRepository interface {
    Create(context.Context, *lemi011b.Datum) error
}

type iController interface {
    Control(iRepository)
}
