package applicator

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/controller/grpc"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
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

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	gracefullyShutdown(cancel)
	repository, err := storage.NewRepository(ctx, a.config)
	if err != nil {
		log.Fatalf("cannot сonnect to mainDB '%s:%d': %v", a.config.Database.Host, a.config.Database.Port, err)
	}
	srv := service.NewManager(repository, a.config)
	grpcServer := grpc.NewServer(srv, &a.config.Transport.UserGrpcTransport)
	err = grpcServer.Run()
	if err != nil {
		log.Fatal(err)
	}
	defer grpcServer.Close()
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
