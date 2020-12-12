package timescale

type iLogger interface {
    Error(string, error)
    Info(string)
}
