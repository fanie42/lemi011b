package stan

// Config TODO
type Config struct {
    Servers   []ServerConfig `env:"SERVERS", yaml:"servers"`
    ClusterID string         `env:"CLUSTER_ID", yaml:"clusterID"`
}

// ServerConfig TODO
type ServerConfig struct {
    Host string `env:"HOST", yaml:"host"`
    Port uint   `env:"PORT", yaml:"port"`
}
