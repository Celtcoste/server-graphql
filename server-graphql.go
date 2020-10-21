package server_graphql

import (
	"flag"
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"github.com/Celtcoste/server-graphql/src/handlers"
	"github.com/Celtcoste/server-graphql/src/health"
	"github.com/Celtcoste/server-graphql/src/logging"
	"github.com/Celtcoste/server-graphql/src/middleware"
	"github.com/Celtcoste/server-graphql/src/postgresql"
	"github.com/Celtcoste/server-graphql/src/setting"
	"github.com/Celtcoste/server-graphql/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	setting.Setup()
	logging.Setup()
	postgresql.Setup(setting.PostgresSetting.Host, setting.PostgresSetting.Port, setting.PostgresSetting.User,
		setting.PostgresSetting.Password, setting.PostgresSetting.DatabaseName)
}

func healthServer(healthAddr *string, errChan chan error) {

	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", health.HealthzHandler)
	hmux.HandleFunc("/readiness", health.ReadinessHandler)
	hmux.HandleFunc("/healthz/status", health.HealthzStatusHandler)
	hmux.HandleFunc("/readiness/status", health.ReadinessStatusHandler)
	healthServer := manners.NewServer()
	healthServer.Addr = *healthAddr
	healthServer.Handler = handlers.LoggingHandler(hmux)

	go func() {
		errChan <- healthServer.ListenAndServe()
	}()
}

func Server(graphqlHandler gin.HandlerFunc, playgroundHandler gin.HandlerFunc) {
	var (
		httpAddr   = flag.String("http", "0.0.0.0:" + os.Getenv("PORT"), "HTTP service address.")
		healthAddr = flag.String("health", "0.0.0.0:" + os.Getenv("HEALTH_PORT"), "Health service address.")
	)
	flag.Parse()

	log.Println("Starting server...")
	log.Printf("Health service listening on %s", *healthAddr)
	log.Printf("HTTP service listening on %s", *httpAddr)

	errChan := make(chan error, 10)

	healthServer(healthAddr, errChan)

	// Setting up Gin
	gin.SetMode(setting.AppSetting.RunMode)
	r := gin.Default()

	r.Use(middleware.GinContextToContextMiddleware())
	r.POST("/query", graphqlHandler)

	if utils.GetEnvStr("APP_ENV") == "TEST" {
		r.GET("/query", graphqlHandler)
	}
	if setting.AppSetting.RunMode == "debug" {
		r.GET("/", playgroundHandler)
	}
	go func() {
		errChan <- r.Run()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			health.SetReadinessStatus(http.StatusServiceUnavailable)
			os.Exit(0)
		}
	}
}