package application

import (
	"api-service/internal/cache"
	"api-service/internal/config"
	"api-service/internal/server"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServiceClient struct {
	RedisClient *cache.RedisClient
	Database    *gorm.DB
	Err         error
}

func StartService(cfg *config.Config) {
	// client := initServiceClient(cfg)
	// gin.ForceConsoleColor()
	router := gin.Default()
	srv := &http.Server{
		Handler: router,
	}
	app := server.NewServer(cfg, router)
	server.ConfigRoutesAndSchedulers(app)
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- app.Run(cfg.HTTP.Port)
	}()
	fmt.Println(srv)
	waitForShutdown(srv, app, serverErr)

}

// func initServiceClient(cfg *config.Config) *ServiceClient {
// 	client := &ServiceClient{}

// 	dbConn := db.InitDB(cfg.DB)
// 	if dbConn == nil {
// 		logrus.Fatal("Failed to initialize database")
// 	}
// 	client.RedisClient = cache.NewRedisClient(
// 		cfg.Redis.Addr,
// 		cfg.Redis.Auth,
// 	)
// 	client.Database = dbConn
// 	return client
// }

func waitForShutdown(srv *http.Server, app *server.Server, serverErr chan error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx, srv); err != nil {
		logrus.Errorf("Shutdown error: %v", err)
	}

	select {
	case err := <-serverErr:
		if err != nil && err != http.ErrServerClosed {
			logrus.Errorf("Server error: %v", err)
		}
	case <-ctx.Done():
		logrus.Warn("Shutdown timeout exceeded")
	}

	logrus.Info("Server stopped cleanly.")
}
