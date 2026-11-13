package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("falta configurar DATABASE_URL")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	cursoRepo := repository.NewCursoRepository(db)
	cursoService := service.NewCursoService(cursoRepo)
	cursoHandler := handler.NewCursoHandler(cursoService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/curso", cursoHandler.Listar)

	return router.Run(":8000")
}
