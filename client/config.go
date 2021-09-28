package client

type Config struct {
	address string
	*Authentication
}

func NewDefaultConfig() *Config {
	return &Config{
		address:        "localhost:18050",
		Authentication: &Authentication{},
	}
}

func (c *Config) SetAddress(addr string) {
	c.address = addr
}

func (c *Config) Address() string {
	return c.address
}