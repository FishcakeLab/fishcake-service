package flags

import (
	"github.com/urfave/cli/v2"
)

const envVarPrefix = "FISHCAKE"

func prefixEnvVars(name string) []string {
	return []string{envVarPrefix + "_" + name}
}

var (
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path to migrations folder",
		EnvVars: prefixEnvVars("MIGRATIONS_DIR"),
	}
	PolygonRpcFlag = &cli.StringFlag{
		Name:    "polygon-rpc",
		Usage:   "HTTP provider URL for L1",
		EnvVars: prefixEnvVars("POLYGON_RPC"),
		Value:   "https://polygon-rpc.com",
		//Required: true,
	}
	HttpHostFlag = &cli.StringFlag{
		Name:    "http-host",
		Usage:   "The host of the api",
		EnvVars: prefixEnvVars("HTTP_HOST"),
		Value:   "0.0.0.0",
	}
	HttpPortFlag = &cli.IntFlag{
		Name:    "http-port",
		Usage:   "The port of the api",
		EnvVars: prefixEnvVars("HTTP_PORT"),
		Value:   8987,
	}
	DbHostFlag = &cli.StringFlag{
		Name:    "migrations-host",
		Usage:   "The host of the database",
		EnvVars: prefixEnvVars("DB_HOST"),
		Value:   "127.0.0.1",
		//Required: true,
	}
	DbPortFlag = &cli.IntFlag{
		Name:    "migrations-port",
		Usage:   "The port of the database",
		EnvVars: prefixEnvVars("DB_PORT"),
		Value:   5432,
		//Required: true,
	}
	DbUserFlag = &cli.StringFlag{
		Name:    "migrations-user",
		Usage:   "The user of the database",
		EnvVars: prefixEnvVars("DB_USER"),
		Value:   "guoshijiang",
		//Required: true,
	}
	DbPasswordFlag = &cli.StringFlag{
		Name:    "migrations-password",
		Usage:   "The host of the database",
		EnvVars: prefixEnvVars("DB_PASSWORD"),
		Value:   "",
		//Required: true,
	}
	DbNameFlag = &cli.StringFlag{
		Name:    "migrations-name",
		Usage:   "The migrations name of the database",
		EnvVars: prefixEnvVars("DB_NAME"),
		Value:   "fishcake",
		//Required: true,
	}
	MetricsHostFlag = &cli.StringFlag{
		Name:    "metrics-host",
		Usage:   "The host of the metrics",
		EnvVars: prefixEnvVars("METRICS_HOST"),
		Value:   "0.0.0.0",
	}
	MetricsPortFlag = &cli.IntFlag{
		Name:    "metrics-port",
		Usage:   "The port of the metrics",
		EnvVars: prefixEnvVars("METRICS_PORT"),
		Value:   7214,
	}
)

var requiredFlags = []cli.Flag{
	MigrationsFlag,
	PolygonRpcFlag,
	HttpPortFlag,
	HttpHostFlag,
	DbHostFlag,
	DbPortFlag,
	DbUserFlag,
	DbPasswordFlag,
	DbNameFlag,
}

var optionalFlags = []cli.Flag{
	MetricsPortFlag,
	MetricsHostFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
