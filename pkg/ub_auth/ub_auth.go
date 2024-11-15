package ub_auth

import (
	u "github.com/ahmdyaasiin/ub-auth-without-notification/v2"
	"github.com/yogarn/parkirkuy/model"
)

type IUbAuth interface {
	Login(identifier string, password string) (user *model.UBAuthRes, err error)
}

type ubAuth struct{}

func Init() IUbAuth {
	return &ubAuth{}
}

func (a *ubAuth) Login(identifier string, password string) (user *model.UBAuthRes, err error) {
	resp, err := u.AuthUB(identifier, password)
	if err != nil {
		return nil, err
	}

	user = &model.UBAuthRes{
		NIM:          resp.NIM,
		FullName:     resp.FullName,
		Email:        resp.Email,
		Faculty:      resp.Fakultas,
		StudyProgram: resp.ProgramStudi,
	}

	return user, nil
}
