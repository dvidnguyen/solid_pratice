package repo

import "fmt"

type User struct {
	Id    string
	Name  string
	Email string
}

type MySQLRepo struct{}

func (mysql MySQLRepo) findById(id string) (*User, error) {
	// Select * from user Where id =
	//	Do something else
	return &User{ID: id, Name: "Nguyen Van A", Email: "a@gmail.com"}, nil
}
func (mysql MySQLRepo) saveProfile(u *User) error {
	fmt.Printf("-> Repo: Đang execute 'UPDATE users SET name=%s WHERE id=%d'\n", u.Name, u.ID)
	return nil
}
