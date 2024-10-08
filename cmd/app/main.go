package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest-api-crud/internal/config"
	"rest-api-crud/internal/delivery/http/v1"
	"rest-api-crud/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	router := httprouter.New()
	handlers := v1.NewHandler(logger)
	handlers.Register(router)
	logger.Info("Handlers successfully registered in http router")

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()

	var listener net.Listener
	var listenErr error
	if cfg.Listen.Type == "sock" {
		logger.Info(`Listen type -> sock <- received from config.yml <---- Creating unix socket listener ---->`)
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		socketPath := path.Join(appDir, "apа.sock")

		listener, listenErr = net.Listen("unix", socketPath)
	} else {
		logger.Infof(`Listen type -> port <- received from config.yml <---- Creating tcp listener on port %s ---->`,
			cfg.Listen.Port)
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		logger.Fatalf("Error happend while creating http server: %v", listenErr)
	}
	logger.Infof("Listener with type -> %s <- created successfully", cfg.Listen.Type)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatalln(server.Serve(listener))
}
