package main

import "fmt"

// 1. Single Responsibility Principle

type User struct {
	name    string
	age     int
	address string
}
type userRepository struct {
}

func (u *User) getName() string {
	fmt.Println(u.name)
	return u.name
}
func (u *User) getAge() int {
	fmt.Println(u.age)
	return u.age
}
func (u *User) getAddress() string {
	fmt.Println(u.address)
	return u.address
}
func (userRepo *userRepository) SaveUser(u *User) {
	// save user to DB
	fmt.Println("Save user: %s %v %d ", u.name, u.age, u.address)
}
func main() {
	user := User{
		name:    "Jack",
		age:     20,
		address: "123",
	}
	user.getName()
	user.getAge()
	user.getAddress()
	// func  saveUser() is responsibility by Repo not dependency on struct User
	userRepo := &userRepository{}
	userRepo.SaveUser(&user)
	// Example for Single Responsibility Principle
}
