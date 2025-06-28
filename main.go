package main

import (
	"fmt"
	"net/http"
)

var firstTask = "Learn Go Now"
var secondTask = "Learn DSA"
var thirdTask = "Eat a cup cake"
var taskItems = []string{firstTask, secondTask, thirdTask}

func main() {

	fmt.Println("### Welcome to ToDoList app")
	fmt.Println("#### List of TodoList")
	// newTask := "Learn K8s"
	// taskItems = addTask(taskItems, newTask)
	// printTasks(taskItems)
	http.HandleFunc("/", hellouser)
	http.HandleFunc("/show-users", showtask)
	http.ListenAndServe(":8080", nil)

}

func hellouser(writer http.ResponseWriter, response *http.Request) {
	var greeting = "### Welcome to ToDoList app"
	fmt.Fprintln(writer, greeting)
}

func showtask(writer http.ResponseWriter, respone *http.Request) {
	for ind, task := range taskItems {
		fmt.Fprintf(writer, "%d.%s\n", ind+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	updateTask := append(taskItems, newTask)
	return updateTask
}
