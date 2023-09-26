package repository

import (
	"context"
	"example/backend-github-trending/model"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	GetListUser(context context.Context) ([]model.User, error)
}