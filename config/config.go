package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Migrations      string   `yaml:"migrations"`
	PolygonRpc      string   `yaml:"polygon_rpc"`
	RpcUrl          string   `yaml:"rpc_url"`
	PolygonChainId  string   `yaml:"polygon_chain_id"`
	HttpHost        string   `yaml:"http_host"`
	HttpPort        int      `yaml:"http_port"`
	DbHost          string   `yaml:"db_host"`
	DbPort          int      `yaml:"db_port"`
	DbName          string   `yaml:"db_name"`
	DbUser          string   `yaml:"db_user"`
	DbPassword      string   `yaml:"db_password"`
	MetricsHost     string   `yaml:"metrics_host"`
	MetricsPort     int      `yaml:"metrics_port"`
	StartBlock      uint64   `yaml:"start_block"`
	EventStartBlock uint64   `yaml:"event_start_block"`
	Contracts       []string `yaml:"contracts"`
}

func New(path string) (*Config, error) {
	var config = new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
