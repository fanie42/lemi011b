package graphql

import (
    "time"

    "github.com/fanie42/lemi011b"
)

// Datum TODO
type Datum struct {
    ID          lemi011b.ID `json:"id"`
    Timestamp   time.Time   `json:"timestamp"`
    X           *int32      `json:"x"`
    Y           *int32      `json:"y"`
    Z           *int32      `json:"z"`
    Temperature *int32      `json:"temperature"`
}

func (d *Datum) toLemi011bDatum() *lemi011b.Datum {
    datum := &lemi011b.Datum{
        ID:          d.ID,
        Timestamp:   d.Timestamp,
        X:           d.X,
        Y:           d.Y,
        Z:           d.Z,
        Temperature: d.Temperature,
    }
}

func (d *Datum) fromLemi011bDatum(datum *lemi011b.Datum) {

}
