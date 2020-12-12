package main

import (
    "github.com/fanie42/lemi011b/daq/infra/stan"
    "github.com/fanie42/lemi011b/vendor/github.com/fanie42/sansa/conf"
)

var config *Config

func init() {
    // yamlFile, _ := os.Open("default.yaml")
    // conf.ParseYAML(yamlFile, config)

    config = new(Config)
    conf.ParseENV(config)
}

// Config TODO
type Config struct {
    Site string      `env:"SITE", yaml:"site"`
    STAN stan.Config `env:"STAN", yaml:"stan"`
}
