package config

import "time"

type Config struct {
	Server   Server         `mapstructure:"server"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
}

type Server struct {
	Address         string        `mapstructure:"address"`
	Timeout         time.Duration `mapstructure:"timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type RabbitMQConfig struct {
	URL   string            `mapstructure:"url"`
	Queue string            `mapstructure:"queue"`
	Retry RabbitRetryConfig `mapstructure:"retry"`
}

type RabbitRetryConfig struct {
	Attempts int     `mapstructure:"attempts"`
	DelayMS  int     `mapstructure:"delay_ms"`
	Backoff  float64 `mapstructure:"backoff"`
}
