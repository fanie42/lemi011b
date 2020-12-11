package main

import (
    "os"

    "github.com/fanie42/lemi011b/daq/app"
    "github.com/fanie42/lemi011b/daq/serial"
    "github.com/fanie42/lemi011b/daq/stan"
    "github.com/fanie42/sansa/conf"
    "github.com/fanie42/sansa/log"
)

var config *Config

func init() {
    conf.ParseYAML("default.yaml", config)
    conf.ParseENV(config)
}

func main() {
    logger := log.New(os.Stdout)

    reader := serial.New(config.Serial, logger)
    publisher := stan.New(config.STAN, logger)

    app.Run(logger, reader, publisher)
}
