package repository

import (
	"amirullazmi/belajar-restfull-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx sql.Tx, category domain.User) domain.User
	Update(ctx context.Context, tx sql.Tx, category domain.User) domain.User
	Delete(ctx context.Context, tx sql.Tx, category domain.User) domain.User
	FindById(ctx context.Context, tx sql.Tx, category domain.User) domain.User
	FindAll(ctx context.Context, tx sql.Tx, category domain.User) []domain.User
}
