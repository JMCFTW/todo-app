package main

import rest "todo-app/todo/controllers"

func main() {
	c := rest.NewRestController()
	c.Execute()
}
