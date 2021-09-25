package config

type Config struct {
	App *app `toml:"app"`
}

func newConfig() *Config {
	return &Config{
		App: newDefaultApp(),
	}
}

func (c *Config) InitGlobal() error {
	global = c
	return nil
}

type app struct {
	Name     string `toml:"name" env:"AUTH_APP_NAME"`
	Host     string `toml:"host" env:"AUTH_APP_HOST"`
	HttpPort string `toml:"http_port" env:"AUTH_HTTP_PORT"`
	GrpcPort string `toml:"grpc_port" env:"AUTH_GRPC_PORT"`
	Key      string `toml:"key" env:"AUTH_APP_KEY"`
}

func newDefaultApp() *app {
	return &app{
		Name:     "auth",
		Host:     "127.0.0.1",
		HttpPort: "8050",
		GrpcPort: "18050",
		Key:      "default",
	}
}

func (a *app) GrpcAddr() string {
	return a.Host + ":" + a.GrpcPort
}
