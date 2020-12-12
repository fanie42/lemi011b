package stan

import (
    "context"

    "github.com/nats-io/stan.go"
)

type iServer interface {
    Create(context.Context, *lemi011b.Datum) error
}

type controller struct {
    srv iServer
    enc func() ([]byte, error)
    dec func([]byte) error
}

func newController(server iServer) *controller {
    return &controller{
        srv: server,
    }
}

func (ctrl *controller) stan(m *stan.Msg) {
    ctrl.dec(m)
}
