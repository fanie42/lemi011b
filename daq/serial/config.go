package serial

// Config TODO
type Config struct {
    Name string `env:"NAME", yaml:"name"`
    Baud int    `env:"BAUD", yaml:"baud"`
}
