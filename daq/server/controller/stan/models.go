package stan

import "time"

type datum struct {
    timestamp   time.Time
    sensorID    string
    x           int64
    y           int64
    z           int64
    temeprature int64
}

func (d *datum) MarshalBinary() (data []byte, err error) {

}

func (d *datum) UnmarshalBinary(data []byte) error {

}
