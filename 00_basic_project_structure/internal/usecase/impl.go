package usecase

import "go-try-out/00_basic_project_structure/internal/repository"

func (u *Usecase) CreateTask(task *repository.Task) error {
	return u.repo.CreateTask(task)
}

func (u *Usecase) UpdateTask(id uint, task *repository.Task) error {
	return u.repo.UpdateTask(id, task)
}

func (u *Usecase) DeleteTask(id uint) error {
	return u.repo.DeleteTask(id)
}

func (u *Usecase) GetTask(id uint) (*repository.Task, error) {
	return u.repo.GetTask(id)
}

func (u *Usecase) GetTasks(offset int, limit int) ([]repository.Task, error) {
	return u.repo.GetTasks(offset, limit)
}

func (u *Usecase) Readiness() error {
	return u.repo.Readiness()
}
