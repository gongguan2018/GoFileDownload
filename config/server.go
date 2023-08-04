package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Server grpcserver `yaml:"server"`
}
type grpcserver struct {
	Port     int    `yaml:"port"`
	DownPath string `yaml:"downpath"`
}

func ReadServerConfig() (*ServerConfig, error) {
	by, err := ioutil.ReadFile("config/server.yml")
	if err != nil {
		return nil, err
	}
	var sc *ServerConfig
	err = yaml.Unmarshal(by, &sc)
	if err != nil {
		return nil, err
	}
	return sc, nil
}
