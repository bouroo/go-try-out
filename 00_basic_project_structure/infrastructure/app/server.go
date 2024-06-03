package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go-try-out/00_basic_project_structure/infrastructure"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type server struct {
	err error
	app *fiber.App
	ctx context.Context
	db  *gorm.DB
}

type Server interface {
	Run() error
}

func NewServer(ctx context.Context) *server {
	// Fiber instance
	app := fiber.New(fiber.Config{
		StreamRequestBody: true,
		Prefork:           bool(runtime.GOMAXPROCS(0) > 1),
	})

	// Connect to database
	db, err := infrastructure.NewSqlite("sqlite.db")
	if err != nil {
		log.Panic(err)
	}

	return &server{
		app: app,
		ctx: ctx,
		db:  db,
	}
}

func (srv *server) Run() error {
	if srv.err != nil {
		return srv.err
	}

	// Setup routes
	srv.SetupRoutes(srv.app)

	// Listen from a different goroutine
	go func() {
		if err := srv.app.Listen("localhost:8080"); err != nil {
			log.Panic(err)
		}
	}()

	// Create channel to signify a signal being sent
	quit := make(chan os.Signal, 1)
	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = srv.app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")

	return nil
}
