package main

// to install gin => go get github.com/gin-gonic/gin

import (
	"fmt"
	"net/http" //The net/http package not only includes the ability to make HTTP requests, but also provides an HTTP server you can use to handle those requests.

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id     string `json:"id"`
	Item   string `json:"title"`
	Status bool   `json:"status"`
}

var todos = []todo{
	{Id: "1", Item: "wash dishes", Status: true},
	{Id: "2", Item: "buy fruits", Status: false},
	{Id: "3", Item: "Read book", Status: false},
	{Id: "4", Item: "mop floor", Status: true},
}

func getTodos(C *gin.Context) {
	C.IndentedJSON(http.StatusOK, todos) // context convert the data into json
}

func postTodos(C *gin.Context) {
	var newTodos todo

	if err := C.BindJSON(&newTodos); err != nil {
		fmt.Println(err)
		return
	}

	todos = append(todos, newTodos)
	C.IndentedJSON(http.StatusCreated, newTodos)
}

func getbyId(C *gin.Context) {
	id := C.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range todos {
		if a.Id == id {
			C.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	C.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	//fmt.Println("hello...")

	router := gin.Default()        //this is our server
	router.GET("/todos", getTodos) //endpoint
	router.GET("/todos/:id", getbyId)
	router.POST("/todos", postTodos)
	router.Run() // the server runs on this localhost:8080 by default

}

// tested in postman
