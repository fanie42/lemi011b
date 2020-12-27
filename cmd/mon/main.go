package main

import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/fanie42/lemi011b/api/protobuf/pb"
    "github.com/jackc/pgx/v4/pgxpool"
    "google.golang.org/protobuf/proto"

    natsio "github.com/nats-io/nats.go"
    stanio "github.com/nats-io/stan.go"
)

func main() {
    done := make(chan interface{})

    clusterID := "marion"
    clientID := "marl111-mon"

    nc, err := natsio.Connect("nats://172.18.30.100:4222")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer nc.Close()
    sc, err := stanio.Connect(clusterID, clientID, stanio.NatsConn(nc))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer sc.Close()

    // TIMESCALE_DB STUFF
    connStr := "postgres://postgres:admin@172.18.30.100:5432/lemi011b"
    dbpool, err := pgxpool.Connect(context.Background(), connStr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    defer dbpool.Close()

    q := `INSERT INTO data (time, station_id, x, y, z, temperature)
        VALUES ($1, 1, $2, $3, $4, $5);`

    sub, err := sc.Subscribe(
        "lemi011b.datumAdded",
        func(m *stanio.Msg) {
            datum := &pb.Datum{}
            proto.Unmarshal(m.Data, datum)

            tag, err := dbpool.Exec(
                context.Background(),
                q,
                time.Unix(0, datum.Timestamp.Nanoseconds),
                datum.X,
                datum.Y,
                datum.Z,
                datum.Temperature,
            )
            if err != nil {
                fmt.Printf("%s: %s\n", tag, err)
            }

            fmt.Printf("%+v\n", datum)
        },
        stanio.DurableName("marl111-mon-durable"),
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer sub.Close()

    // time.Sleep(time.Second * 10)
    <-done
}
