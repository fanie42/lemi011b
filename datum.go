package lemi011b

import (
    "time"
)

// const datumKind string = 1

// Datum TODO
type Datum struct {
    Time        time.Time `json:"time"`
    X           *int32    `json:"x"`
    Y           *int32    `json:"y"`
    Z           *int32    `json:"z"`
    Temperature *int32    `json:"temperature"`
}

// // Tags TODO
// type Tags struct {
//     Site       string `json:"site"`
//     Instrument string `json:"instrument"`
//     Iteration  uint8  `json:"iteration"`
// }

// // ID TODO
// func (d *Datum) ID() *sansa.UID {
//     sansa.NewUID(api, "datum")
// }

// // Kind TODO
// func (d *Datum) Kind(iteration uint8) string {
//     return "MAR_L11_1_Datum"
// }
