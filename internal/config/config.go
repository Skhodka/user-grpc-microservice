package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env     string         `yaml:"env"`
	Timeout time.Duration  `yaml:"timeout"`
	GRPC    GRPCConfig     `yaml:"grpc"`
	Storage PostgresConfig `yaml:"postgres"`
}

type GRPCConfig struct {
	Port int `yaml:"grpc_port"`
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	SslMode  string `yaml:"sslmode"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()

	if configPath == "" {
		panic("empty config path")
	}

	if _, err := os.Stat(configPath); err != nil {
		panic("unable to find config file")
	}

	var config *Config

	data, err := os.ReadFile(configPath)

	if err != nil {
		panic("unable to read config file")
	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		panic("unable to parse config file")
	}

	posConf := loadPostgresConf()

	config.Storage.Host = posConf.Host
	config.Storage.Port = posConf.Port
	config.Storage.User = posConf.User
	config.Storage.Password = posConf.Password
	config.Storage.Dbname = posConf.Dbname

	if posConf == nil {
		panic("unable to parse postgres config")
	}

	config.Storage = *posConf

	return config
}

func loadPostgresConf() *PostgresConfig {
	port, _ := strconv.Atoi(mustGetenv("DB_PORT"))

	return &PostgresConfig{
		Host:     mustGetenv("DB_HOST"),
		Port:     port,
		User:     mustGetenv("DB_APP_USER"),
		Password: mustGetenv("DB_APP_USER_PASS"),
		Dbname:   mustGetenv("DB_NAME"),
	}
}

func mustGetenv(key string) string {
	val := os.Getenv(key)

	if val == "" {
		panic(fmt.Sprintf("empty %s", key))
	}

	return val
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "input config path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
