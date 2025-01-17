package main

import (
	"fmt"
	"todo/database"
	"todo/handlers"
)

func main() {
	err := database.InitDB()
	if err != nil {
		fmt.Print("DB error : ", err)
	}
	database.InitTodoDAO()
	handlers.InitHandler()
}
