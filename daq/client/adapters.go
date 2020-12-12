package client

import (
    "bufio"
    "context"
    "fmt"
    "io"
    "strconv"
    "strings"
    "time"

    "github.com/fanie42/lemi011b"
    "github.com/nats-io/stan.go"
)

// datum TODO
type datum struct {
    timestamp time.Time
    bytes     []byte
}

func (d *datum) toLemi011bDatum() *lemi011b.Datum {
    result := new(lemi011b.Datum)

    s := strings.Split(string(d.bytes), ", ")
    x, _ := strconv.ParseInt(s[0], 10, 64)
    y, _ := strconv.ParseInt(s[1], 10, 64)
    z, _ := strconv.ParseInt(s[2], 10, 64)
    temperature, _ := strconv.ParseInt(s[3], 10, 64)

    result.Timestamp = d.timestamp
    result.X = &x
    result.Y = &y
    result.Z = &z
    result.Temperature = &temperature

    return result
}

func (r io.Reader) control(svc lemi011b.ICreator) error {
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        d := &datum{
            timestamp: time.Now(),
            bytes:     scanner.Bytes(),
        }
        err := svc.Create(context.Background(), d.toLemi011bDatum())
        if err != nil {
            fmt.Println(err)
        }
    }

    return scanner.Err()
}

func (sc stan.Conn) present(ctx context.Context, event string, datum *lemi011b.Datum) error {
    data := datum.MarshalBinary()

    err := sc.Publish("lemi011b.datum."+event, data)
    if err != nil {
        return err
    }

    return nil
}
