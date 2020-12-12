package daq

// Server TODO
type Server struct {
    log  iLogger
    gw   iGateway
    ctrl iController
}

// NewServer TODO
func NewServer(
    logger iLogger,
    presenter iPresenter,
    controller iController,
) *Server {
    return &Server{
        log:  logger,
        pres: presenter,
        ctrl: controller,
    }
}

// Run TODO
func (c *Client) Run() {
    svc := NewDAQService(c.log, c.gw)
    c.ctrl.Control(c.log, svc)
}
