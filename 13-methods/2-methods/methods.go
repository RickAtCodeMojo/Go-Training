package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("%s at %s", u.name, u.email)
}

func (u *user) Set(n string, e string) {
	u.name = n
	u.email = e
}
func (u user) SetInfo(n string, e string) {
	u.name = n
	u.email = e
}
func main() {
	var u user
	fmt.Println(u)
	u.Set("Rick", "rick@gmail.com")
	u.SetInfo("Fred", "fred@yahoo.com")
	fmt.Println(u)
}
