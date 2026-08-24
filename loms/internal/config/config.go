package config

type Config struct {
	// Server
	GRPCPort int
	HTTPPort int

	// Database
	DBUrl      string
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}
