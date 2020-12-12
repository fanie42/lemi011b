package main

import (
    "github.com/fanie42/sansa/logger"
    "github.com/nats-io/stan.go"
    tarm "github.com/tarm/serial"
)

var (
    config     *Config
    log        *logger.Logger
    serialPort *tarm.Port
    stanConn   stan.Conn
)

func main() {
    controller := client.NewSerial(serialPort)
    presenter := client.NewSTAN(stanConn)

    service := client.New(logger, presenter)
    controller.control(service)
}
