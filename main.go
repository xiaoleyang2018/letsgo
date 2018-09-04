package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/qinhao/letsgo/config"
	"github.com/qinhao/letsgo/logger"
	"github.com/qinhao/letsgo/models"
	"github.com/qinhao/letsgo/ormx"
	"github.com/qinhao/letsgo/router"
)

func main() {
	cfg := config.New()
	err := cfg.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("parse config args error: %v\n", err)
		os.Exit(1)
	}

	if cfg.Version {
		config.Version()
		os.Exit(0)
	}

	// setting logger
	err = logger.Parse(cfg.LogPath, cfg.RunMode, cfg.LogFormat)
	if err != nil {
		fmt.Printf("parse logger error: %v\n", err)
		os.Exit(1)
	}

	// init models
	err = models.Init(ormx.DB(cfg.DB), cfg.RunMode)
	if err != nil {
		logger.Errorf("init database error: %v\n", err)
		os.Exit(1)
	}

	// init routers
	e, err := router.Init(cfg.HandlerLogPath, cfg.RunMode)
	if err != nil {
		logger.Errorf("init router error: %v", err)
		os.Exit(1)
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	s := &http.Server{
		Addr: addr,
	}

	// start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
		return
	}

	logger.Infof("Server exiting")
}
