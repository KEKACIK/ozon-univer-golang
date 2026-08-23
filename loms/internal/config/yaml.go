package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.yaml.in/yaml/v3"
)

type ServerYamlConfig struct {
	Host     string `yaml:"host"`
	HttpPort int16  `yaml:"http_port"`
	GrpcPort int16  `yaml:"grpc_port"`
}

type DatabaseYamlConfig struct {
	Host     string `yaml:"host"`
	Port     int16  `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type YamlConfig struct {
	Server   ServerYamlConfig   `yaml:"server"`
	Database DatabaseYamlConfig `yaml:"database"`
}

func NewConfigFromYaml() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile("loms/internal/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data = []byte(os.ExpandEnv(string(data)))

	var yamlCfg YamlConfig
	if err := yaml.Unmarshal(data, &yamlCfg); err != nil {
		log.Fatal(err)
	}

	return &Config{
		GRPCPort: int(yamlCfg.Server.GrpcPort),
		HTTPPort: int(yamlCfg.Server.HttpPort),

		DBHost:     yamlCfg.Database.Host,
		DBPort:     int(yamlCfg.Database.Port),
		DBName:     yamlCfg.Database.Name,
		DBUser:     yamlCfg.Database.User,
		DBPassword: yamlCfg.Database.Password,
	}
}
