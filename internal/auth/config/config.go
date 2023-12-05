package config

import "time"

type Config struct {
	HttpServer HttpServer `yaml:"HttpServer"`
	Database   Database   `yaml:"Database"`
	Auth       Auth       `yaml:"JwtSecretKey"`
	Transport  Transport  `yaml:"Transport"`
	Kafka      Kafka      `yaml:"Kafka"`
}

type Database struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
	SslMode  string `yaml:"SslMode"`
}
type Auth struct {
	PasswordSecretKey string `yaml:"PasswordSecretKey"`
	JwtSecretKey      string `yaml:"JwtSecretKey"`
}
type HttpServer struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}
type Transport struct {
	UserHttpTransport UserHttpTransport `yaml:"UserHttpTransport"`
	UserGrpcTransport UserGrpcTransport `yaml:"ClassGrpcTransport"`
}

type UserHttpTransport struct {
	Host    string        `yaml:"Host"`
	Timeout time.Duration `yaml:"Timeout"`
}
type UserGrpcTransport struct {
	Port string `yaml:"Port"`
}

type Kafka struct {
	Brokers  []string `yaml:"brokers"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

type Producer struct {
	Topic string `yaml:"topic"`
}

type Consumer struct {
	Topics []string `yaml:"topics"`
}
