package serial

import (
    "context"

    tarm "github.com/tarm/serial"
)

// Controller TODO
type Controller struct {
    sp *tarm.Port
}

// New TODO
func New(
    serialPort *tarm.Port,
) *Controller {
    return &Controller{
        sp: serialPort,
    }
}

// PublishDatum TODO
func (ctrl *Controlelr) PublishDatum(ctx context.Context, datum *lemi011b.Datum) {

}
