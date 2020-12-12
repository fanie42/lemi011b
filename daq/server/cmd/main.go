package server

import (
    "log"
    "os"

    "github.com/nats-io/nats-server/server"
)

func main() {
    logger := log.New(os.Stdout)

    database := timescale.New(config.Timescale, logger)
    transport := stan.New(config.STAN, logger)

    repository := sql.New(logger, database)
    presenter := producer.New(logger, transport)

    application := server.New(logger, repository, presenter)

    controller := consumer.New(logger, transport)
    controller.Handle(application)
}
