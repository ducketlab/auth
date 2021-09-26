package config

type Config struct {
	App *app `toml:"app"`
	Log *log `toml:"log"`
}

func newConfig() *Config {
	return &Config{
		App: newDefaultApp(),
		Log: newDefaultLog(),
	}
}

func (c *Config) InitGlobal() error {
	global = c
	return nil
}

type app struct {
	Name       string `toml:"name" env:"AUTH_APP_NAME"`
	Host       string `toml:"host" env:"AUTH_APP_HOST"`
	HttpPort   string `toml:"http_port" env:"AUTH_HTTP_PORT"`
	HttpPrefix string `toml:"http_prefix" env:"AUTH_HTTP_PREFIX"`
	GrpcPort   string `toml:"grpc_port" env:"AUTH_GRPC_PORT"`
	Key        string `toml:"key" env:"AUTH_APP_KEY"`
}

type log struct {
	Level   string    `toml:"level" env:"AUTH_LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"AUTH_LOG_PATH"`
	Format  LogFormat `toml:"format" env:"AUTH_LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"AUTH_LOG_TO"`
}

func newDefaultApp() *app {
	return &app{
		Name:     "auth",
		Host:     "127.0.0.1",
		HttpPort: "8050",
		HttpPrefix: "/",
		GrpcPort: "18050",
		Key:      "default",
	}
}

func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}

func (a *app) GrpcAddr() string {
	return a.Host + ":" + a.GrpcPort
}

func (a *app) HttpAddr() string {
	return a.Host + ":" + a.HttpPort
}
