package usecase

import "go-try-out/00_basic_project_structure/internal/repository"

type Usecase struct {
	repo *repository.Repository
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}
