package config

type Config struct {
	ServerPorts []string
	MinAcks     int
	MinReads    int
}

func New() *Config {
	return &Config{
		ServerPorts: []string{"localhost:8080", "localhost:8081", "localhost:8082"},
		MinAcks:     2,
		MinReads:    2,
	}
}
