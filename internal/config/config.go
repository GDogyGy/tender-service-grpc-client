package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"ENV" env:"ENV" env-default:"local" env-required:"true"`
	HTTPServer `yaml:"HTTP_SERVER"`
	Redis      `yaml:"REDIS_SERVER"`
	Kafka      `yaml:"KAFKA"`
	GrpcRemote string `yaml:"GRPC_REMOTE" env:"ENV" env-default:"localhost:8081" env-required:"true"`
	DebugLevel string `yaml:"DEBUG_LEVEL" env:"DEBUG_LEVEL" env-default:"info"`
}

type HTTPServer struct {
	Address     string        `yaml:"SERVER_ADDRESS" env:"SERVER_ADDRESS" env-default:"localhost:8082" env-required:"true"`
	Timeout     time.Duration `yaml:"TIMEOUT" env:"TIMEOUT" env-default:"6s"`
	IdleTimeout time.Duration `yaml:"IDLE_TIMEOUT" env:"IDLE_TIMEOUT" env-default:"60s"`
}

type Kafka struct {
	Address      string `yaml:"KAFKA_ADDRESS" env:"KAFKA_ADDRESS" env-default:"localhost:29092" env-required:"true"`
	DefaultTopic string `yaml:"KAFKA_DEFAULT_TOPIC" env:"KAFKA_DEFAULT_TOPIC" env-default:"model-events" env-required:"true"`
}

type Redis struct {
	Address       string `yaml:"REDIS_ADDRESS" env:"REDIS_ADDRESS" env-default:"localhost:6379" env-required:"true"`
	RedisPassword string `yaml:"REDIS_PASSWORD" env:"REDIS_PASSWORD" env-default:"123"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH id not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH is not set: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return &cfg
}
