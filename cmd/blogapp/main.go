package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"github.com/anazcodes/blog-crud-api/internal/api/http/blogapp"
	"github.com/anazcodes/blog-crud-api/internal/business/blogbus"
	"github.com/anazcodes/blog-crud-api/internal/repository/blogrepo"
)

func main() {
	port := flag.String("port", "3000", "Server Port")
	capacity := flag.Int("cache-capacity", 30, "Cache Capacity")

	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	log.Printf("Starting server on port: %s with cache-capacity: %d", *port, *capacity)

	app := di(*port, *capacity)
	go app.Serve()

	<-ctx.Done()

	log.Println("Graceful shutdown triggered")

	if err := app.Shutdown(context.Background()); err != nil {
		log.Fatalln(err)
	}
}

// di injects dependencies and initializes the application.
func di(port string, capacity int) blogapp.App {
	repo := blogrepo.NewRepository(capacity)
	bus := blogbus.NewBusiness(repo)
	app := blogapp.NewApp(port, bus)

	return app
}
