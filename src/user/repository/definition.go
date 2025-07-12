package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	ctx    context.Context
	dbw    *sqlx.DB
	dbr    *sqlx.DB
	schema string
	query  map[string]string
}

func NewUserRepository(context context.Context, writer *sqlx.DB, reader *sqlx.DB) IUserRepository {
	return &UserRepository{ctx: context, dbw: writer, dbr: reader}
}

type IUserRepository interface {
	GetUser() error
	SetUser() error
	UpsertUser() error
	DeleteUser() error
}
