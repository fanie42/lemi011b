package app

import (
    "bufio"
    "io"
)

// Run TODO
func Run(log iLogger, reader iReader, publisher iPublisher) {
    scanner := bufio.NewScanner(reader)
    for scanner.Scan() {
        b := scanner.Bytes()
        sink.Publish(b)
    }
    reader.Close()
    publisher.Close()
}

type iLogger interface {
    Error(string, error)
    Info(string)
}

type iReader interface {
    io.ReadCloser
}

type iProducer interface {
    Publish([]byte)
    Close() error
}
