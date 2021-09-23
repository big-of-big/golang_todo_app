package main

import "todo_app/app/models"

func main() {
	t, _ := models.GetTodo(2)
	t.DeleteTodo()
}
