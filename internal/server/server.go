package server

import (
	"api-service/internal/cache"
	"api-service/internal/config"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	Cfg         *config.Config
	Gin         *gin.Engine
	DB          *gorm.DB
	RedisClient *cache.RedisClient
}

func (s *Server) Run(addr string) error {
	logrus.Infof("Starting server on :%s", addr)
	return s.Gin.Run(":" + addr)
}

func NewServer(cfg *config.Config, httpHandle *gin.Engine) *Server {
	return &Server{
		Cfg: cfg,
		Gin: httpHandle,
	}
}
func (s *Server) Shutdown(ctx context.Context, srv *http.Server) error {
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down Gin: %v", err)
	}
	if s.DB != nil {
		if err := s.CloseDB(); err != nil {
			return fmt.Errorf("error closing DB: %v", err)
		}
		logrus.Info("Database connection closed successfully.")
	} else {
		logrus.Info("Database connection is nil. Skipping close.")
	}

	return nil
}
func (s *Server) CloseDB() error {
	sqlDB, err := s.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
