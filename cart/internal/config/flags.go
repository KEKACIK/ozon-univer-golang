package config

import "flag"

func NewConfigFromFlags() *Config {
	const (
		defaultAddr        = ":8080"
		defaultLomsAddr    = "http://loms:8080"
		defaultProductAddr = "HARDCORE"
	)

	config := Config{}
	flag.StringVar(&config.Addr, "addr", defaultAddr, "server address, default: "+defaultAddr)
	flag.StringVar(&config.LomsAddr, "loms_addr", defaultLomsAddr, "loms address, default: "+defaultLomsAddr)
	flag.StringVar(&config.ProductAddr, "product_addr", defaultProductAddr, "product-service address, default: "+defaultProductAddr)
	flag.Parse()

	return &config
}
