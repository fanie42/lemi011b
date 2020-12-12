package stan

type iLogger interface {
    Error(string, error)
    Info(string)
}
