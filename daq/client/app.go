package client

import (
    "bufio"
    "time"
)

// Client TODO
type Client struct {
    log  iLogger
    pres iPresenter
    ctrl iController
}

// New TODO
func New(
    logger iLogger,
    presenter iPresenter,
    controller iController,
) *Client {
    c := &Client{
        log:  logger,
        pres: presenter,
        ctrl: controller,
    }

    return c
}

// Start TODO
func (c *Client) Start() {
    scanner := bufio.NewScanner(c.read)
    for scanner.Scan() {
        d := &datum{
            timestamp: time.Now(),
            bytes:     scanner.Bytes(),
        }
        c.pub.Publish(d.toLemi011bDatum())
    }
    c.pub.Close()
}
