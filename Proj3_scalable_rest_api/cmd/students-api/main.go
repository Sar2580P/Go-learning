package main

import (
	"Proj3_scalable_rest_api/internal/config"
	"Proj3_scalable_rest_api/internal/http/handlers/student"
	"Proj3_scalable_rest_api/internal/storage/sqlite"
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err:= sqlite.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /apis/students", student.GetList(storage))


	// setup server
	server := http.Server{
		Addr: cfg.HTTPServer.Addr, 
		Handler: router,
	}

	slog.Info("server started %s", slog.String("address", cfg.HTTPServer.Addr))

	done:= make(chan os.Signal, 1)

	// listen for interrupt signals
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)  // listen for interrupt signals to gracefully shutdown the server

	// graceful shutdown: listen for interrupt signal, then shutdown the server gracefully with a timeout context
	// run the server in a goroutine so that it doesn't block the main thread, allowing us to listen for shutdown signals
	go func(){
		err:= server.ListenAndServe()
		if err!= nil{
			log.Fatalf("failed to start server: %s", err.Error())	
		}
	}()


	<-done

	slog.Info("shutting down server...")

	// timer to force shutdown the server if it doesn't shutdown gracefully within the timeout duration
	ctx, cancel:= context.WithTimeout(context.Background(), 5*time.Second)
	
	defer cancel()

	if err:= server.Shutdown(ctx); err!= nil{
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
