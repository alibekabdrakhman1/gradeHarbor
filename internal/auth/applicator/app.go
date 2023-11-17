package applicator

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/controller"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/controller/http"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/service"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/transport"
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
	repo, err := storage.NewRepository(ctx, a.config)
	if err != nil {
		log.Fatalf("cannot сonnect to mainDB '%s:%d': %v", a.config.Database.Host, a.config.Database.Port, err)
	}
	userHttpTransport := transport.NewUserHttpTransport(a.config.Transport.UserHttpTransport)
	userGrpcTransport := transport.NewUserGrpcTransport(a.config.Transport.UserGrpcTransport)
	authService := service.NewManager(repo, a.config.Auth, userHttpTransport, userGrpcTransport)
	endPointHandler := http.NewManager(authService)
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
