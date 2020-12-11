package serial

type iLogger interface {
    Error(string, error)
    Info(string)
}
