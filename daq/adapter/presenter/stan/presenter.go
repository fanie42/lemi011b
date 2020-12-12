package stan

import (
    "context"

    "github.com/nats-io/stan.go"
)

// Presenter TODO
type Presenter struct {
    sc stan.Conn
}

// New TODO
func New(
    stanConn stan.Conn,
) *Presenter {
    return &Presenter{
        sc: stanConn,
    }
}

// PublishDatum TODO
func (ctrl *Controller) PublishDatum(
    ctx context.Context,
    datum *lemi011b.Datum,
) {

    ctrl.sc.Publish(subject)
}
