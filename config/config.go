package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Migrations      string       `yaml:"migrations"`
	PolygonRpc      string       `yaml:"polygon_rpc"`
	RpcUrl          string       `yaml:"rpc_url"`
	PolygonChainId  string       `yaml:"polygon_chain_id"`
	HttpHost        string       `yaml:"http_host"`
	HttpPort        int          `yaml:"http_port"`
	DbHost          string       `yaml:"db_host"`
	DbPort          int          `yaml:"db_port"`
	DbName          string       `yaml:"db_name"`
	DbUser          string       `yaml:"db_user"`
	DbPassword      string       `yaml:"db_password"`
	MetricsHost     string       `yaml:"metrics_host"`
	MetricsPort     int          `yaml:"metrics_port"`
	StartBlock      uint64       `yaml:"start_block"`
	EventStartBlock uint64       `yaml:"event_start_block"`
	Contracts       []string     `yaml:"contracts"`
	AliConfig       AliConfig    `yaml:"ali_config"`
	ContractInfo    ContractInfo `yaml:"contract_info"`
}

type AliConfig struct {
	RegionId        string `yaml:"region_id"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}

type ContractInfo struct {
	FccToken             string `yaml:"fcc_token"`
	UsdtToken            string `yaml:"usdt_token"`
	RedemptionPool       string `yaml:"redemption_pool"`
	DirectSalePool       string `yaml:"direct_sale_pool"`
	InvestorSalePool     string `yaml:"investor_sale_pool"`
	NFTManager           string `yaml:"nft_manager"`
	FishcakeEventManager string `yaml:"fishcake_event_manager"`
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
