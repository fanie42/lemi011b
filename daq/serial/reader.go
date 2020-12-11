package serial

import (
    "io"

    tarm "github.com/tarm/serial"
)

// Reader TODO
type Reader struct {
    io.ReadCloser
    log iLogger
}

// New TODO
func New(config Config, logger iLogger) *Reader {
    r := &Reader{
        log: logger,
    }

    serialPort, err := tarm.OpenPort(&tarm.Config{
        Name: config.Name,
        Baud: config.Baud,
    })
    if err != nil {
        r.log.Error("failed to connect to serial port", err)
    }

    r.ReadCloser = serialPort

    return r
}
