package main

import (
	// srcからの相対パス
	"fmt"
	"log"
	"todo_app/config"
)

func main() {
	fmt.Println(config.Config)
	log.Println("test")
}
