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
		Name:     "polygon-rpc",
		Usage:    "HTTP provider URL for L1",
		EnvVars:  prefixEnvVars("POLYGON_RPC"),
		Required: true,
	}
	HttpHostFlag = &cli.StringFlag{
		Name:     "http-host",
		Usage:    "The host of the api",
		EnvVars:  prefixEnvVars("HTTP_HOST"),
		Required: true,
	}
	HttpPortFlag = &cli.IntFlag{
		Name:     "http-port",
		Usage:    "The port of the api",
		EnvVars:  prefixEnvVars("HTTP_PORT"),
		Value:    8987,
		Required: true,
	}
	DbHostFlag = &cli.StringFlag{
		Name:     "db-host",
		Usage:    "The host of the database",
		EnvVars:  prefixEnvVars("DB_HOST"),
		Required: true,
	}
	DbPortFlag = &cli.IntFlag{
		Name:     "db-port",
		Usage:    "The port of the database",
		EnvVars:  prefixEnvVars("DB_PORT"),
		Required: true,
	}
	DbUserFlag = &cli.StringFlag{
		Name:     "db-user",
		Usage:    "The user of the database",
		EnvVars:  prefixEnvVars("DB_USER"),
		Required: true,
	}
	DbPasswordFlag = &cli.StringFlag{
		Name:     "db-password",
		Usage:    "The host of the database",
		EnvVars:  prefixEnvVars("DB_PASSWORD"),
		Required: true,
	}
	DbNameFlag = &cli.StringFlag{
		Name:     "db-name",
		Usage:    "The db name of the database",
		EnvVars:  prefixEnvVars("DB_NAME"),
		Required: true,
	}
	MetricsHostFlag = &cli.StringFlag{
		Name:     "metrics-host",
		Usage:    "The host of the metrics",
		EnvVars:  prefixEnvVars("METRICS_HOST"),
		Required: true,
	}
	MetricsPortFlag = &cli.IntFlag{
		Name:     "metrics-port",
		Usage:    "The port of the metrics",
		EnvVars:  prefixEnvVars("METRICS_PORT"),
		Value:    7214,
		Required: true,
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
