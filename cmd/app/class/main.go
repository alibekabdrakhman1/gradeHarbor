package main

import (
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/applicator"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title Class Service
// @version 1.0
// @description Class Service

// @host localhost:8082
// @BasePath /

// @securityDefinitions.apikey	BearerAuth
// @name Authorization
// @in header
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With(zap.String("app", "class-service"))

	cfg, err := loadConfig("./config/class")
	if err != nil {
		l.Error(err)
		l.Fatalf("failed to load config err: %v", err)
	}

	app := applicator.New(l, &cfg)
	app.Run()
}
func loadConfig(path string) (config config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to ReadInConfig err: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to Unmarshal config err: %w", err)
	}
	fmt.Println(config.Auth)

	return config, nil
}
