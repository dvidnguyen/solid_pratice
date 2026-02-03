package main

import "fmt"

// [O] is Open/closed principle

//////////--------------------Bad Practice--------------------/////////////////////////
//type Employee struct {
//	name string
//	role string
//}
//
//func (e Employee) getSalary() int64 {
//	if e.role == "ceo" {
//		return 1000
//	} else if e.role == "developer" {
//		return 300
//	}
//	return 0
//}
//
//func main() {
//	employee1 := Employee{
//		name: "John",
//		role: "developer",
//	}
//	employee2 := Employee{
//		name: "Jane",
//		role: "ceo",
//	}
//
//	print("Salary developer is ", employee1.getSalary())
//	print("Salary ceo is ", employee2.getSalary())
//}

//////////////-----------------------------Good Practice-------------------/////////////////////////////////////////////////////////

type role interface {
	getSalary() int64
}
type employee struct {
	name string
	role role
}

type dev struct{}

func (d dev) getSalary() int64 {
	return 300
}

type ceo struct{}

func (c ceo) getSalary() int64 {
	return 1000
}

func (em employee) getSalary() int64 {
	return em.role.getSalary()
}

func main() {
	em1 := employee{
		name: "John",
		role: dev{},
	}
	em2 := employee{
		name: "Jane",
		role: ceo{},
	}

	// using interface for the role giving the flexibility to extend
	//the code by adding more roles without modifying existing getSalary func
	fmt.Println("Salary", em1.getSalary())
	fmt.Println("Salary", em2.getSalary())
}
