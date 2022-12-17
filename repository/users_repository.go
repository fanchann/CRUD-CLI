package repository

import (
	"context"

	"crud_mysqli_cli/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, name string, email string) (entity.User, error)
	Update(ctx context.Context, name string, email string, id int32) (entity.User, error)
	Delete(ctx context.Context, id int32) (string, error)
	FindById(ctx context.Context, id int32) (entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
}
