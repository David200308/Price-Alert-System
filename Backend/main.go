package main

import (
	"fmt"
	"os"

	_ "github.com/David200308/go-api/Backend/docs"
	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/routers"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.InitRedis()

	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		fmt.Println("SENTRY_DSN is not set")
		return
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		EnableTracing:    true,
		TracesSampleRate: 0.1,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	r := routers.SetupRouter()
	r.Use(sentrygin.New(sentrygin.Options{}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}
