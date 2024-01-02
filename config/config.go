package config

import (
	"github.com/urfave/cli/v2"

	"github.com/FishcakeLab/fishcake-service/flags"
)

type Config struct {
	Migrations  string
	PolygonRpc  string
	HttpHost    string
	HttpPort    int
	DbHost      string
	DbPort      int
	DbName      string
	DbUser      string
	DbPassword  string
	MetricsHost string
	MetricsPort int
}

func NewConfig(ctx *cli.Context) (*Config, error) {
	return &Config{
		Migrations:  ctx.String(flags.MigrationsFlag.Name),
		PolygonRpc:  ctx.String(flags.PolygonRpcFlag.Name),
		HttpHost:    ctx.String(flags.HttpHostFlag.Name),
		HttpPort:    ctx.Int(flags.HttpHostFlag.Name),
		DbHost:      ctx.String(flags.DbHostFlag.Name),
		DbPort:      ctx.Int(flags.DbPortFlag.Name),
		DbName:      ctx.String(flags.DbNameFlag.Name),
		DbUser:      ctx.String(flags.DbUserFlag.Name),
		DbPassword:  ctx.String(flags.DbPasswordFlag.Name),
		MetricsHost: ctx.String(flags.MetricsHostFlag.Name),
		MetricsPort: ctx.Int(flags.MetricsPortFlag.Name),
	}, nil
}
