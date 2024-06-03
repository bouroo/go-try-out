package main

import (
	"context"
	"log/slog"
	"os"

	"go-try-out/00_basic_project_structure/infrastructure/app"

	_ "go.uber.org/automaxprocs"
)

func main() {
	ctx := context.Background()
	if err := app.NewServer(ctx).Run(); err != nil {
		slog.Error("Main: Run", "error", err)
		os.Exit(1)
	}
}
