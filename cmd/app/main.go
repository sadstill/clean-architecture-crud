package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest-api-crud/internal/config"
	"rest-api-crud/internal/delivery/http/v1"
	mongodb2 "rest-api-crud/internal/repository"
	"rest-api-crud/pkg/database/mongodb"
	"rest-api-crud/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()

	cfg := config.GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	mongodbClient, err := mongodb.NewClient(ctx, cfg.MongoDB.Host, cfg.MongoDB.Port,
		cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	if err != nil {
		panic(err)
	}
	storage := mongodb2.NewStorage(mongodbClient, cfg.MongoDB.Collection, logger)

	users, err := storage.FindAll(context.Background())
	fmt.Println(users)

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

		socketPath := path.Join(appDir, "app.sock")

		listener, listenErr = net.Listen("unix", socketPath)
	} else {
		logger.Infof(`Listen type -> port <- received from config.yml <---- Creating tcp listener on port %s ---->`,
			cfg.Listen.Port)
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}
	logger.Infof("Listener with type -> %s <- created successfully", cfg.Listen.Type)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatalln(server.Serve(listener))
}
