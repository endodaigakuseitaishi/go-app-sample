package main

import (
	"fmt"
	"go-todo-sample/app/controllers"
	"go-todo-sample/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartmainServer()
}