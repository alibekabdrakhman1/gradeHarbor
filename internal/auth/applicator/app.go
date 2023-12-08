package applicator

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/controller"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/controller/consumer"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/controller/http"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/dto"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/service"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/transport"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/kafka"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

type App struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func New(logger *zap.SugaredLogger, cfg *config.Config) *App {
	return &App{
		config: cfg,
		logger: logger,
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	gracefullyShutdown(cancel)
	userVerificationProducer, err := kafka.NewProducer(a.config.Kafka)
	if err != nil {
		a.logger.Panicf("failed NewProducer err: %v", err)
	}
	repo, err := storage.NewRepository(ctx, a.config, a.logger)
	if err != nil {
		log.Fatalf("cannot сonnect to mainDB '%s:%d': %v", a.config.Database.Host, a.config.Database.Port, err)
	}

	userHttpTransport := transport.NewUserHttpTransport(a.config.Transport.UserHttpTransport, a.logger)
	userGrpcTransport := transport.NewUserGrpcTransport(a.config.Transport.UserGrpcTransport, a.logger)
	a.logger.Info("transports success")
	userVerificationConsumerCallback := consumer.NewUserVerificationCallback(a.logger)

	userVerificationConsumer, err := kafka.NewConsumer(a.logger, a.config.Kafka, userVerificationConsumerCallback)
	if err != nil {
		a.logger.Panicf("failed NewConsumer err: %v", err)
	}

	go userVerificationConsumer.Start()

	serviceDTO := &dto.UserTokenServiceDTO{
		Repository:               repo,
		JwtSecretKey:             a.config.Auth.JwtSecretKey,
		PasswordSecretKey:        a.config.Auth.PasswordSecretKey,
		UserHttpTransport:        userHttpTransport,
		UserGrpcTransport:        userGrpcTransport,
		Logger:                   a.logger,
		UserVerificationProducer: userVerificationProducer,
	}
	authService := service.NewManager(serviceDTO)

	endPointHandler := http.NewManager(authService, a.logger)
	HTTPServer := controller.NewServer(a.config, endPointHandler)

	return HTTPServer.StartHTTPServer(ctx)
}
func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
