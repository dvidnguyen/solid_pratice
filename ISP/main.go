package main

//  ISO - Interface Segregation Principle

type UserRepoQuery interface {
	findById(id string)
	findList()
}

type UserRepoCmd interface {
	createUser()
	UpdateUser()
}

type UserQuery struct{}
type UserCmd struct{}

func (userQ UserQuery) findById(id string) {
	//	Do something
}
func (userQ UserQuery) findList() {
	//	Do something
}

func (userCmd UserCmd) createUser() {
	//	Do something
}
func (userCmd UserCmd) UpdateUser() {
	//	Do something
}

// Segregation interface so that flexible and avoid impl func or method do need use
