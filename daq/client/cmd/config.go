package main

import (
    "fmt"
    "os"

    "github.com/fanie42/lemi011b/infrastructure/serial"
    "github.com/fanie42/lemi011b/infrastructure/stan"
    "github.com/fanie42/sansa/conf"
)

// Config TODO
type Config struct {
    Site    string        `env:"SITE", yaml:"site"`
    Version uint8         `env:"VERSION", yaml:"version"`
    Log     log.config    `env:"LOG", yaml:"log"`
    Serial  serial.Config `env:"SERIAL", yaml:"serial"`
    STAN    stan.Config   `env:"STAN", yaml:"stan"`
}

func init() {
    // yamlFile, _ := os.Open("default.yaml")
    // conf.ParseYAML(yamlFile, config)

    config = new(Config)
    conf.ParseENV(config)

    fmt.Printf("%+v\n", config)

    log := logger.New(&config.Logger, os.Stdout)
    serialPort := serial.New(&config.Serial)
    stanConn := stan.New(&config.STAN)
}
