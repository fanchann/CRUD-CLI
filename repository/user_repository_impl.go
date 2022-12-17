package repository

import (
	"context"
	"database/sql"
	"strconv"

	"crud_mysqli_cli/entity"
)

type userImplRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userImplRepository{DB: db}
}

func (repo *userImplRepository) Insert(ctx context.Context, name string, email string) (entity.User, error) {
	sqlQuery := "insert into user (name,email) values (?,?)"

	result, err := repo.DB.ExecContext(ctx, sqlQuery, name, email)
	if err != nil {
		return entity.User{}, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.User{}, nil
	}

	return entity.User{Id: int32(id), Name: name, Email: email}, nil
}

func (repo *userImplRepository) Delete(ctx context.Context, id int32) (string, error) {
	sqlQuery := "delete from user where id = ?"

	result, err := repo.DB.ExecContext(ctx, sqlQuery, id)

	if err != nil {
		return "404 not found", err
	}

	res, err := result.RowsAffected()
	if err != nil {
		return "Something went wrong", err
	}
	if res == 0 {
		return "Id " + strconv.Itoa(int(id)) + " not found ", err
	}

	return "Succes delete id " + strconv.Itoa(int(id)), nil

}

func (repo *userImplRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	sqlQuery := "select * from user"

	result, err := repo.DB.QueryContext(ctx, sqlQuery)
	defer result.Close()
	if err != nil {
		return nil, err
	}

	var rows []entity.User

	for result.Next() {
		column := entity.User{}
		result.Scan(&column.Id, &column.Name, &column.Email)
		rows = append(rows, column)
	}

	return rows, nil

}

func (repo *userImplRepository) Update(ctx context.Context, name, email string, id int32) (entity.User, error) {
	sqlQuery := "update user set name = ?, email = ? where id = ?"

	result, err := repo.DB.ExecContext(ctx, sqlQuery, name, email, id)
	if err != nil {
		return entity.User{}, err
	}

	res, err := result.RowsAffected()
	if err != nil {
		return entity.User{}, err
	}

	if res == 0 {
		return entity.User{Id: id, Name: name, Email: email}, nil
	} else {
		return entity.User{}, nil
	}
}

func (repo *userImplRepository) FindById(ctx context.Context, id int32) (entity.User, error) {
	sqlQuery := "select * from user where id = ? limit 1"
	var name, email string
	result, err := repo.DB.QueryContext(ctx, sqlQuery, id)
	if err != nil {
		return entity.User{}, err
	}

	if result.Next() {
		result.Scan(&id, &name, &email)
		return entity.User{Id: id, Name: name, Email: email}, nil
	} else {
		return entity.User{}, nil
	}
}
