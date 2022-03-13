package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cbodonnell/api-template/cache"
	"github.com/cbodonnell/api-template/config"
	"github.com/cbodonnell/api-template/handlers"
	"github.com/cbodonnell/api-template/logging"
	"github.com/cbodonnell/api-template/repositories"
	"github.com/cbodonnell/api-template/services"
	"github.com/cbodonnell/api-template/workers"
)

func main() {
	// Get configuration
	ENV := os.Getenv("ENV")
	conf, err := config.ReadConfig(ENV)
	if err != nil {
		log.Fatal(err)
	}

	// create cache layer
	cache := cache.NewRedisCache(conf)
	err = cache.FlushDB()
	if err != nil {
		log.Println(err)
	}
	// create repository
	repo := repositories.NewPSQLRepository(conf, cache)
	defer repo.Close()
	// create service
	service := services.NewExampleService(conf, repo)
	// create handler (TODO: Make an options struct??)
	handler := handlers.NewMuxHandler(conf, service)
	if conf.AllowedOrigins != "" {
		handler.AllowCORS(strings.Split(conf.AllowedOrigins, ","))
	}
	r := handler.GetRouter()

	// create worker
	worker := workers.NewExampleWorker(conf)
	go worker.Start()

	// Set log file
	if conf.LogFile != "" {
		logFile := logging.SetLogFile(conf.LogFile)
		defer logFile.Close()
	}

	// Run server
	log.Println(fmt.Sprintf("Serving on port %d", conf.Port))

	// TLS
	if conf.SSLCert == "" {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r))
	}
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", conf.Port), conf.SSLCert, conf.SSLKey, r))
}
