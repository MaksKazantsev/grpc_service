package config

import (
	"flag"
	"github.com/goccy/go-yaml"
	"os"
)

type Config struct {
	Port int `yaml:"port"`

	Env string `yaml:"env"`

	Database Database
}

type Database struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	User     string `yaml:"users"`
	Name     string `yaml:"name"`
}

func MustLoad() *Config {
	var cfg *Config

	path := fetchPath()

	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read cfg path" + err.Error())
	}

	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to unmarshal cfg" + err.Error())
	}

	return cfg
}

func fetchPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path to cfg")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
