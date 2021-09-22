package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	u := &models.User{}
	u.Name = "lamp"
	u.Email = "lamp@example.com"
	u.PassWord = "password"
	fmt.Println(u)

	u.CreateUser()
}
