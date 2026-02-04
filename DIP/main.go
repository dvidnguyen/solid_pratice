package main

import "fmt"

// DIP - Dependency Inversion Principle

type IUserRepo interface {
	findById(id string) (*User, error)
	saveProfile(u *User) error
}

type UserService struct {
	repo IUserRepo
}

func NewUserService(repo IUserRepo) *UserService {
	return &UserService{repo}
}
func (userService *UserService) findById(id string) error {
	user, _ := userService.repo.findById(id)
	if user == nil {
		fmt.Println("User không tồn tại")
		return err
	}
	return nil
}
func (userService *UserService) UpdateProfile(u *User) error {
	user, _ := userService.repo.findById(u.Id)
	if user == nil {
		return err
	}

	user.Name = u.Name
	user.Email = u.Email
	userService.repo.saveProfile(user)
	return nil
}
