package main

import (
    "fmt"
    "os"
    "time"

    "github.com/fanie42/lemi011b/api/protobuf/pb"
    "google.golang.org/protobuf/proto"

    natsio "github.com/nats-io/nats.go"
    stanio "github.com/nats-io/stan.go"
)

func main() {
    done := make(chan interface{})

    // Connect to NATS cluster
    clusterID := "marion"
    clientID := "marl111-fs"

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

    // Subscribe to NATS with durable subscription
    sub, err := sc.Subscribe(
        "lemi011b.datumAdded",
        func(m *stanio.Msg) {
            datum := &pb.Datum{}
            proto.Unmarshal(m.Data, datum)

            timestamp := time.Unix(0, datum.Timestamp.Nanoseconds).UTC()

            // Open file
            file, err := os.OpenFile(
                "marlem1_"+timestamp.Format("2006-01-02")+".dat",
                os.O_WRONLY|os.O_CREATE|os.O_APPEND,
                0755,
            )
            if err != nil {
                fmt.Println(err)
            }
            defer file.Close()

            // Write datum to file
            _, err = file.WriteString(fmt.Sprintf(
                "%s, %d, %d, %d, %d\n",
                timestamp.Format("2006-01-02 15:04:05.000000"),
                datum.X,
                datum.Y,
                datum.Z,
                datum.Temperature,
            ))
            if err != nil {
                fmt.Println(err)
            }

            fmt.Printf("%+v\n", datum)
        },
        stanio.DurableName("marl111-fs-durable"),
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer sub.Close()

    // time.Sleep(time.Second * 10)
    <-done
}
