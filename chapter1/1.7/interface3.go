package main

import "fmt"

type Message interface {
	sending()
}

type User struct {
	name  string
	phone string
}

func (u *User) sending() {
	fmt.Printf("User (name: %s) is sending...\n", u.name)
}

type Admin struct {
	name string
	sex  string
}

func (a *Admin) sending() {
	fmt.Printf("Admin (name: %s) is sending...\n", a.name)
}

func send(m Message) {
	m.sending()
}

func main() {
	u := User{"jaluik user", "18581111111"}
	a := Admin{"jaluik admin", "male"}

	send(&u)
	send(&a)

}
