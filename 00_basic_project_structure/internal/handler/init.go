package handler

import "go-try-out/00_basic_project_structure/internal/usecase"

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
