package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

)


type curso struct {
	ID		string	`json:"id"`
	Nombre	string	`json:"nombre"`
}

var cursos = []curso {
	{ID: "1", Nombre: "Curso 1"},
	{ID: "2", Nombre: "Curso 2"},
	{ID: "3", Nombre: "Curso 3"},
}

func main(){
	router := gin.Default()

	router.GET("/health", getHealth)

	router.GET("/curso", getCurso)

	router.Run("localhost:8000")
}

func getHealth(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func getCurso(c *gin.Context){
	c.IndentedJSON(http.StatusOK, cursos)
}
