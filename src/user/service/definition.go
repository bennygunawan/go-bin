package service

import (
	"context"

	"github.com/bennygunawan/go-bin/src/user/repository"
)

type UserServices struct {
	ctx  context.Context
	repo repository.IUserRepository
}

func NewUserServices(context context.Context, repository repository.IUserRepository) IUserServices {
	return &UserServices{ctx: context, repo: repository}
}

type IUserServices interface {
	GetUser() error
	SetUser() error
	UpsertUser() error
	DeleteUser() error
}
