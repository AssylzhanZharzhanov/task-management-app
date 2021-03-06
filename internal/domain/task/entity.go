package domain

import (
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
)

const (
	TasksTableName = "users"
)

type TaskID int64

type Task struct {
	ID             TaskID        `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserID         domain.UserID `json:"user_id" gorm:"not null;column:user_id"`
	Title          string        `json:"title" gorm:"not null;column:title"`
	Description    string        `json:"description" gorm:"not null;column:description"`
	StartDate      int64         `json:"start_date" gorm:"not null;column:start_date"`
	EndDate        int64         `json:"end_date" gorm:"not null;column:end_date"`
	ReminderPeriod int64         `json:"reminder_period" gorm:"not null;column:reminder_period"`
	IsCompleted    bool          `json:"is_completed" gorm:"not null;column:is_completed"`
	CreatedAt      int64         `json:"created_at" gorm:"not null;column:created_at"`
}

func NewCreatedTask(dto *CreateTaskDTO) *Task {
	return &Task{
		UserID:         dto.UserID,
		Title:          dto.Title,
		Description:    dto.Description,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
		ReminderPeriod: dto.ReminderPeriod,
	}
}

func NewUpdatedTask(dto *UpdateTaskDTO) *Task {
	return &Task{
		ID:             dto.ID,
		UserID:         dto.UserID,
		Title:          dto.Title,
		Description:    dto.Description,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
		IsCompleted:    dto.IsCompleted,
		ReminderPeriod: dto.ReminderPeriod,
	}
}
