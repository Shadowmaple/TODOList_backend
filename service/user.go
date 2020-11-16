package service

import "todolist_backend/model"

func CreateNewAccount(qq string) (*model.UserModel, error) {
	user := &model.UserModel{
		Username: "",
		Nickname: "todolist-user",
		QQ:       qq,
	}
	if err := user.Create(); err != nil {
		return nil, err
	}
	return user, nil
}
