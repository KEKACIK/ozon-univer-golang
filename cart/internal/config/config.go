package config

type Config struct {
	// Client
	LomsGRPCAddr string

	// Server
	GRPCPort int
	HTTPPort int

	// Database
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}
