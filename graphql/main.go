package main

import (
    "log"
    "os"

    "github.com/fanie42/lemi011b/daq/app"
    "github.com/fanie42/lemi011b/daq/stan"
)

func main() {
    logger := log.New(os.Stdout)

    repo := timescale.New(config.Timescale, logger)
    handler := stan.New(config.STAN, logger)

    app.Run(logger, repo, handler)
}
