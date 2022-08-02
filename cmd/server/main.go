package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0xlilnas/go-rest-api/database"
	"github.com/0xlilnas/go-rest-api/handler"
	"github.com/0xlilnas/go-rest-api/initializer"
	"github.com/gin-gonic/gin"
)

// struct which contains pointer
// database connections
type App struct {
}

//setting up application
func (app *App) Run() error {
	fmt.Println("Setting up our app")
	gin.SetMode(gin.ReleaseMode)

	//load data from env
	initializer.Init()

	//connecting to database
	database.NewDatabase()

	//router setup
	router := gin.Default()

	handler.SetupRoutes(&handler.Config{
		R: router,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// set up handler to go routine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// make channel for listen os.Siganal and setup notify Signal
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// setup withTimeout to preserve connection before close
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(c); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")

	return nil
}

func main() {
	fmt.Println("My App")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our rest api")
		fmt.Println(err)
	}
}
