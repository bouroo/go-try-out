package handler

import (
	"go-try-out/00_basic_project_structure/internal/repository"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (u *Handler) CreateTask(c *fiber.Ctx) error {
	task := new(repository.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	err := u.usecase.CreateTask(task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(JSend{
		Status:  JSendStatusSuccess,
		Message: "Task created successfully",
		Data:    task,
	})
}

func (u *Handler) GetTask(c *fiber.Ctx) error {
	strID := c.Params("id", "0")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	task, err := u.usecase.GetTask(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(JSend{
		Status:  JSendStatusSuccess,
		Message: "Task retrieved successfully",
		Data:    task,
	})
}

func (u *Handler) GetTasks(c *fiber.Ctx) error {
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	tasks, err := u.usecase.GetTasks(offset, limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(JSend{
		Status:  JSendStatusSuccess,
		Message: "Tasks retrieved successfully",
		Data:    tasks,
	})
}

func (u *Handler) UpdateTask(c *fiber.Ctx) error {
	task := new(repository.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	strID := c.Params("id", "0")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	err = u.usecase.UpdateTask(uint(id), task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(JSend{
		Status:  JSendStatusSuccess,
		Message: "Task updated successfully",
		Data:    task,
	})
}

func (u *Handler) DeleteTask(c *fiber.Ctx) error {
	strID := c.Params("id", "0")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	err = u.usecase.DeleteTask(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(JSend{
		Status:  JSendStatusSuccess,
		Message: "Task deleted successfully",
	})
}

func (u *Handler) Readiness(c *fiber.Ctx) error {
	err := u.usecase.Readiness()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(JSend{
			Status:  JSendStatusError,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(JSend{
		Status:  JSendStatusSuccess,
		Message: "OK",
	})
}

func (u *Handler) Liveness(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(JSend{
		Status:  JSendStatusSuccess,
		Message: "OK",
	})
}
