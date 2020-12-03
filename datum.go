package lemi011b

import (
    "time"

    "github.com/fanie42/sansa"
)

const api uint8 = 1

// Datum TODO
type Datum struct {
    ID          sansa.UID `json:"id"`
    Time        time.Time `json:"time"`
    X           *int32    `json:"x"`
    Y           *int32    `json:"y"`
    Z           *int32    `json:"z"`
    Temperature *int32    `json:"temperature"`
}
