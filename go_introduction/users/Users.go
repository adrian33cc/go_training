package users

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id      int
	Name    string
	friends []User
}

type Employee struct {
	User
	Active bool
}

type Cashier interface {
	CalcularTotal(item ...float32) float32
	deactivate()
}

type Admin interface {
	DeactivateEmployee(c Cashier)
}

func (u User) SayHello() {
	fmt.Println("Hola me llamo:", u.Name)
}

func (u *User) AddFriend(friend User) {
	u.friends = append(u.friends, friend)
}

func (u User) ListFriends() {
	fmt.Println("Lista de amigos de", u.Name)
	for i, friend := range u.friends {
		//fmt.Printf("%+v\n", friend)
		fmt.Printf("%d. %s\n", i+1, friend.Name)
	}
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

func (e *Employee) DeactivateEmployee(c Cashier) {
	c.deactivate()
}

func (e *Employee) deactivate() {
	e.Active = false
}

func (e *Employee) CalcularTotal(items ...float32) float32 {
	var sum float32

	if !e.Active {
		fmt.Println("Empleado desactivado")
		return 0
	}

	for _, item := range items {
		sum += item
	}

	return sum
}
