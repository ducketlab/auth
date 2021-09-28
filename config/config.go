package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	mgoclient *mongo.Client
)

type Config struct {
	App   *app     `toml:"app"`
	Log   *log     `toml:"log"`
	Mongo *mongodb `toml:"mongodb"`
}

func newConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		Log:   newDefaultLog(),
		Mongo: newDefaultMongoDB(),
	}
}

func (c *Config) InitGlobal() error {
	global = c
	mclient, err := c.Mongo.getClient()
	if err != nil {
		return err
	}
	mgoclient = mclient
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

type mongodb struct {
	Endpoints []string `toml:"endpoints" env:"AUTH_MONGO_ENDPOINTS" envSeparator:","`
	UserName  string   `toml:"username" env:"AUTH_MONGO_USERNAME"`
	Password  string   `toml:"password" env:"AUTH_MONGO_PASSWORD"`
	Database  string   `toml:"database" env:"AUTH_MONGO_DATABASE"`
	AuthDB    string   `toml:"auth_db" env:"AUTH_MONGO_AUTHDB"`
}

func newDefaultApp() *app {
	return &app{
		Name:       "auth",
		Host:       "127.0.0.1",
		HttpPort:   "8050",
		HttpPrefix: "/",
		GrpcPort:   "18050",
		Key:        "default",
	}
}

func newDefaultMongoDB() *mongodb {
	return &mongodb{
		Database:  "auth",
		Endpoints: []string{"127.0.0.1:27017"},
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

func (m *mongodb) Client() *mongo.Client {
	if mgoclient == nil {
		panic("please load mongo client first")
	}

	return mgoclient
}

func (m *mongodb) authDB() string {
	if m.AuthDB != "" {
		return m.AuthDB
	}

	return m.Database
}

func (m *mongodb) GetDB() *mongo.Database {
	return m.Client().Database(m.Database)
}


func (m *mongodb) getClient() (*mongo.Client, error) {
	opts := options.Client()

	cred := options.Credential{
		AuthSource: m.authDB(),
	}

	if m.UserName != "" && m.Password != "" {
		cred.Username = m.UserName
		cred.Password = m.Password
		cred.PasswordSet = true
		opts.SetAuth(cred)
	}
	opts.SetHosts(m.Endpoints)
	opts.SetConnectTimeout(5 * time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("new mongodb client error, %s", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("ping mongodb server(%s) error, %s", m.Endpoints, err)
	}

	return client, nil
}