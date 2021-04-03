package phttp

import (
	"context"
	"fmt"
	"github.com/LiveScraper/models"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func GracefullyServe(r http.Handler, conf *models.Config) {
	//Graceful Shutdown
	srv := &http.Server{
		Addr:     conf.Port.ToString(),
		Handler:  r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("listen: %s\n", err)
			os.Exit(1)
		} else {
			fmt.Println("Listening On :", conf.Port)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutdown Server ...")
	//Timeout of 10 sec
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(conf.ShutdownTimeout))
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown Error:", err)
	}
	fmt.Println("Server Shutdown Success")
}
