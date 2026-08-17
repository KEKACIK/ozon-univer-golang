package config

import "flag"

func NewConfigFromFlags() *Config {
	const (
		defaultAddr = ":8080"
	)

	config := Config{}
	flag.StringVar(&config.Addr, "addr", defaultAddr, "server address, default: "+defaultAddr)
	flag.Parse()

	return &config
}
