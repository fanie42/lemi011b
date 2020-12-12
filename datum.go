package lemi011b

import (
    "time"
)

// Datum TODO
type Datum struct {
    Timestamp   time.Time `json:"timestamp"`
    SensorID    sensor.ID `json:"sensor_id"`
    X           int64     `json:"x"`
    Y           int64     `json:"y"`
    Z           int64     `json:"z"`
    Temperature int64     `json:"temperature"`
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
