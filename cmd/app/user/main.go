package main

import (
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/applicator"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With(zap.String("app", "user-service"))

	cfg, err := loadConfig("./config/user")
	if err != nil {
		l.Fatalf("failed to load config err: %v", err)
	}
	fmt.Println(cfg.Transport)
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
