package repository

func (r *Repository) CreateTask(task *Task) error {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	err := dbTx.Create(task).Error
	if err != nil {
		return err
	}
	return dbTx.Commit().Error
}

func (r *Repository) UpdateTask(id uint, task *Task) error {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	err := dbTx.Model(&Task{}).Where(Task{ID: id}).Updates(task).Error
	if err != nil {
		return err
	}
	return dbTx.Commit().Error
}

func (r *Repository) DeleteTask(id uint) error {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	err := dbTx.Delete(&Task{}, id).Error
	if err != nil {
		return err
	}
	return dbTx.Commit().Error
}

func (r *Repository) GetTask(id uint) (*Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *Repository) GetTasks(offset int, limit int) ([]Task, error) {
	var tasks []Task
	err := r.db.Offset(offset).Limit(limit).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) Readiness() error {
	err := r.db.Raw("SELECT 1").Scan(nil).Error
	return err
}
