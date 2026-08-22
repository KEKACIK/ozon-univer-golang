package config

import (
	"flag"
	"fmt"
)

func NewConfigFromFlags() *Config {
	const (
		defaultLomsGRPCAddr = "loms:8082"
		defaultHTTPPort     = 8080
		defaultGRPCPort     = 8082
	)

	config := Config{}
	flag.StringVar(&config.LomsGRPCAddr, "loms_grpc_addr", defaultLomsGRPCAddr, fmt.Sprintf("Loms GRPS server address, default: %s", defaultLomsGRPCAddr))
	flag.IntVar(&config.HTTPPort, "http_port", defaultHTTPPort, fmt.Sprintf("HTTP server port, default: %d", defaultHTTPPort))
	flag.IntVar(&config.GRPCPort, "grpg_port", defaultGRPCPort, fmt.Sprintf("gRPC server port, default: %d", defaultGRPCPort))
	flag.Parse()

	return &config
}
