package server

// Server TODO
type Server struct {
    log  iLogger
    ctrl iController
    repo iRepository
}

// New TODO
func New(logger iLogger, controller iController, repository iRepository) *Server {
    s := &Server{
        log:  logger,
        ctrl: controller,
        repo: repository,
    }

    return s
}

// Start TODO
func (s *Server) Start() {
    s.ctrl.Control(s.repo)
}
