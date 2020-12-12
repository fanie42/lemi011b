package stan

import (
    "context"

    "github.com/nats-io/stan.go"
)

// Controller TODO
type Controller struct {
    sc  stan.Conn
    log iLogger
    svc iService
}

// New TODO
func New(
    stanConn stan.Conn,
) *Controller {
    return &Controller{
        sc: stanConn,
    }
}

// Subscribe TODO
func (ctrl *Controller) Subscribe(
    ctx context.Context,
    subject string,
    handler func(m *stan.Msg),
    durableName string,
) {
    sub, err := ctrl.sc.Subscribe(
        subject,
        func(m *stan.Msg) {
            data := ctrl.dec(m.Data)
            ctrl.svc.Create()
        },
        stan.DurableName(durableName),
    )

    if err != nil {
        ctrl.log.Error("failed to subscribe", err)
    }
}
