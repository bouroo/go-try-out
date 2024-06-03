package app

import (
	"go-try-out/00_basic_project_structure/internal/handler"
	"go-try-out/00_basic_project_structure/internal/repository"
	"go-try-out/00_basic_project_structure/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func (srv *server) SetupRoutes(app *fiber.App) {

	appRepo := repository.NewRepository(srv.db)
	appUsecase := usecase.NewUsecase(appRepo)
	appHandler := handler.NewHandler(appUsecase)

	app.Get("/readiness", appHandler.Readiness)
	app.Get("/liveness", appHandler.Liveness)

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Get("/tasks", appHandler.GetTasks)
		apiV1.Get("/tasks/:id", appHandler.GetTask)
		apiV1.Post("/tasks", appHandler.CreateTask)
		apiV1.Put("/tasks/:id", appHandler.UpdateTask)
		apiV1.Delete("/tasks/:id", appHandler.DeleteTask)
	}
}
