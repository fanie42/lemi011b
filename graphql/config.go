package main

import (
    "github.com/fanie42/lemi011b/daq/stan"
    "github.com/fanie42/sansa/conf"
)

// Config TODO
type Config struct {
    Site string      `env:"SITE", yaml:"site"`
    STAN stan.Config `env:"STAN", yaml:"stan"`
}

var config *Config

func init() {
    conf.ParseYAML("default.yaml", config)
    conf.ParseENV(config)
}
