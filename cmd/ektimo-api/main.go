package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/auxcube/ektimo-api/docs"
	"github.com/auxcube/ektimo-api/internal/health"
	"github.com/auxcube/ektimo-api/pkg/config"
	"github.com/auxcube/ektimo-api/pkg/database"
	"github.com/auxcube/ektimo-api/pkg/httpserver"
	"github.com/auxcube/ektimo-api/pkg/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Ektimo API
// @version         1.0
// @host            localhost:8080

func main() {
	config.Initialize(os.Getenv("ENV"))
	log := logger.New(config.Global.Log.Level)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Global.PG.Host, 5432, config.Global.PG.User, config.Global.PG.Password, config.Global.PG.Database)
	db, err := database.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed opening connection to postgres: %v", err)
	}
	defer db.Close()
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatal("failed creating schema resources: %v", err)
	}

	// TODO: initialize MVC components/dependencies here
	healthController := health.NewController(db)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()
		log.Info("%s %s - %s - %d", c.Request.Method, c.Request.RequestURI, latency, status)
	})

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	router.GET("/swagger/*any", swaggerHandler)

	// TODO: register routes
	healthController.RegisterRoutes(router)

	httpServer := httpserver.New(router, httpserver.Port(config.Global.HTTP.Port))
	log.Info("Server started listening on port %d", config.Global.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("Server closed via interrupt signal: %s", s)
	case err := <-httpServer.Notify():
		log.Error("Server closed via httpserver.Notify: %w", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Error("Server closed via httpserver.Shutdown: %w", err)
	}
}
