package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	//u := &models.User{}
	//u.Name = "lamp"
	//u.Email = "lamp@example.com"
	//u.PassWord = "password"
	//u.CreateUser()
	//
	//user, _ := models.GetUser(2)
	//user.CreateTodo("first todo")

	//t, _ := models.GetTodo(1)
	//fmt.Println(t)
	//
	//user, _ := models.GetUser(2)
	//user.CreateTodo("second todo")
	//
	todos, _ := models.GetTodos()
	for _, todo := range todos {
		fmt.Println(todo)
	}
}
