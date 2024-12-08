package repository

import (
	"amirullazmi/belajar-restfull-api/helper"
	"amirullazmi/belajar-restfull-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(name, email) values (?, ?)"

	result, err := tx.ExecContext(ctx, SQL, user.Email, user.Name)

	helper.PanicIfError(err)

	ID, err := result.LastInsertId()

	helper.PanicIfError(err)

	user.ID = string(ID)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update User set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.ID)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from User where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.ID)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select id, name from User where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User is not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, name from User"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
