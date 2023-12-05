package config

import "time"

type Config struct {
	HttpServer HttpServer `yaml:"HttpServer"`
	Database   Database   `yaml:"Database"`
	Auth       Auth       `yaml:"JwtSecretKey"`
	Transport  Transport  `yaml:"Transport"`
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
	UserHttpTransport  UserHttpTransport  `yaml:"UserHttpTransport"`
	ClassGrpcTransport ClassGrpcTransport `yaml:"ClassGrpcTransport"`
}

type UserHttpTransport struct {
	Host string `yaml:"Host"`
}

type ClassGrpcTransport struct {
	Port string `yaml:"Port"`
}
