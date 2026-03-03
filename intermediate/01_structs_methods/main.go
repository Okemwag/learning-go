package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func (u User) Label() string {
	return fmt.Sprintf("%d:%s", u.ID, u.Name)
}

func (u *User) Rename(next string) {
	u.Name = next
}

func main() {
	user := User{ID: 1, Name: "Asha"}
	fmt.Println("Label:", user.Label())

	user.Rename("Asha N.")
	fmt.Println("Updated:", user.Label())
}
