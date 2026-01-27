package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

type Task struct {
	ID primitive.ObjectID `bson:"_id" json:"-"`
	Title primitive.ObjectID `bson:"title" form:"title" binding:"required" json:"title"`
	UserId primitive.ObjectID `bson:"userID" json:"-"`
}

type TaskRepository interface {
	Create(c context.Context ,task *Task) error
	FetchByUserId(c context.Context, userId string) ([]Task, error)
}

type TaskUseCase interface {
	Create(c context.Context, task *Task) error
	FetchByUserId(c context.Context, userId string) ([]Task, error)
}