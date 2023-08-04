package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RpcServer   configRpcServer   `yaml:"rpcserver"`
	HttpdServer configHttpdServer `yaml:"httpdserver"`
	RequestURL  []string          `yaml:"requesturl"`
}
type configRpcServer struct {
	ServerIP string `yaml:"serverip"`
	RPCPort  int    `yaml:"rpcport"`
}
type configHttpdServer struct {
	HttpdPort     int    `yaml:"httpdport"`
	HttpdProtocol string `yaml:"protocol"`
}

//读取配置文件，反序列化到结构体中
func ReadClientConfig() (*Config, error) {
	var c *Config
	by, err := ioutil.ReadFile("config/client.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(by, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
