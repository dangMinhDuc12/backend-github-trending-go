package repository

import (
	"context"
	"example/backend-github-trending/model"
	"example/backend-github-trending/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	GetListUser(context context.Context) ([]model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
}