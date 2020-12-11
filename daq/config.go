package main

import (
    "github.com/fanie42/lemi011b/daq/stan"
    "github.com/fanie42/sansa/conf"
    "github.com/tarm/serial"
)

// Config TODO
type Config struct {
    Site    string        `env:"SITE", yaml:"site"`
    Version uint8         `env:"VERSION", yaml:"version"`
    Serial  serial.Config `env:"SERIAL", yaml:"serial"`
    STAN    stan.Config   `env:"STAN", yaml:"stan"`
}

var config *Config

func init() {
    conf.ParseYAML("default.yaml", config)
    conf.ParseENV(config)
}
