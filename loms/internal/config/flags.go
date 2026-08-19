package config

import (
	"flag"
	"fmt"
)

func NewConfigFromFlags() *Config {
	const (
		defaultHTTPPort = 8080
		defaultGRPCPort = 8082
	)

	config := Config{}
	flag.IntVar(&config.HTTPPort, "http_port", defaultHTTPPort, fmt.Sprintf("HTTP server address, default: %d", defaultHTTPPort))
	flag.IntVar(&config.GRPCPort, "grpg_port", defaultGRPCPort, fmt.Sprintf("gRPC server address, default: %d", defaultGRPCPort))
	flag.Parse()

	return &config
}
