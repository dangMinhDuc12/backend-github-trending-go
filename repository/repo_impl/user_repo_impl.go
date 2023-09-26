package repoimpl

import (
	"context"
	"example/backend-github-trending/banana"
	"example/backend-github-trending/db"
	"example/backend-github-trending/model"
	"example/backend-github-trending/repository"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl {
		sql: sql,
	}
}

func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	query := `
		INSERT INTO public.users (user_id, full_name, email, password, role, created_at, updated_at)
		VALUES (:user_id, :full_name, :email, :password, :role, :created_at, :updated_at)
	`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, query, user)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
				if (err.Code.Name() == "unique_violation") {
					return user, banana.UserConflict
				}
		}

		return user, banana.SignUpFail
	}

	return user, nil
}

func (u *UserRepoImpl) GetListUser(context context.Context) ([]model.User, error) {
	query := `
		SELECT
			user_id,
			full_name,
			email,
			role,
			created_at,
			updated_at
		FROM public.users
	`

	users := []model.User{}

	err := u.sql.Db.SelectContext(context, &users, query) 

	if err != nil {
		return users, banana.GetListUserFail
	}

	return users, nil
}