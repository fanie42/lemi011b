package daq

// Client TODO
type Client struct {
    log  iLogger
    gw   iGateway
    ctrl iController
}

// New TODO
func New(
    logger iLogger,
    gateway iGateway,
    controller iController,
) *Client {
    return &Client{
        log:  logger,
        gw:   gateway,
        ctrl: controller,
    }
}

// Run TODO
func (c *Client) Run() {
    svc := NewDAQService(c.log, c.gw)
    c.ctrl.Control(c.log, svc)
}
