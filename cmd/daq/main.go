package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/fanie42/lemi011b/api/protobuf/pb"
    "github.com/tarm/serial"
    "google.golang.org/protobuf/proto"

    natsio "github.com/nats-io/nats.go"
    stanio "github.com/nats-io/stan.go"
)

// Observation: Timestamping absolutely HAS to happen in a seperate routine
// than writing to stan!
func main() {
    clusterID := "marion"
    clientID := "marl111-daq"

    nc, err := natsio.Connect("nats://172.18.30.100:4222")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    sc, err := stanio.Connect(clusterID, clientID, stanio.NatsConn(nc))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    port, err := serial.OpenPort(&serial.Config{
        Name: "COM5",
        Baud: 19200,
    })
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    time.Sleep(time.Second)
    err = port.Flush()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    scanner := bufio.NewScanner(port)
    for scanner.Scan() {
        fmt.Println(time.Now())
        timestamp := &pb.Timestamp{
            Nanoseconds: time.Now().UnixNano(),
        }
        b := scanner.Bytes()
        s := strings.Split(string(b), ", ")
        x, _ := strconv.ParseInt(s[0], 10, 64)
        y, _ := strconv.ParseInt(s[0], 10, 64)
        z, _ := strconv.ParseInt(s[0], 10, 64)
        temperature, _ := strconv.ParseInt(s[0], 10, 64)

        datum := &pb.Datum{
            SensorID:    "marl111",
            Timestamp:   timestamp,
            X:           x,
            Y:           y,
            Z:           z,
            Temperature: temperature,
        }

        msg, err := proto.Marshal(datum)
        if err != nil {
            fmt.Println(err)
            continue
        }

        err = sc.Publish(
            "lemi011b.datumAdded",
            msg,
        )
        if err != nil {
            fmt.Println(err)
            continue
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
