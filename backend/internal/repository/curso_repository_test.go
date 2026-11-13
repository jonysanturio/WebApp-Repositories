package repository_test

import (
	"os"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)


func TestCursoRepository_CreateAndGet(t *testing.T){
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Fatal("Omite la configuracion de TEST_DATABASE_URL")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("Abrir conexion: %v", err)
	}
	t.Cleanup(func() {
		db.Close()
	})
	if err := db.Ping(); err != nil {
		t.Fatalf("Conectando a Postgres: %v", err)
	}

	_, err = db.Exec(`CREATE TEMPORARY TABLE cursos(
		id BIGSERIAL PRIMARY KEY,
		titulo TEXT NOT NULL,
		estado TEXT NOT NULL
	)`)
	if err != nil {
		t.Fatalf("Creando tabla temporal: %v", err)
	}
}





