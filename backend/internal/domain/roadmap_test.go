package domain_test

import (
	"testing"
	"backend/internal/domain"
)

func TestChangeStateInvalidTransitionDraftToPublicedError(t *testing.T) {
	roadmap := &domain.Roadmap{
		ID: "urs-1",
		Cursos : []domain.Curso{},
	}
	cursoBorrador := domain.Curso{
		ID: "curso-1",
		Titulo: "Curso de prueba",
		Estado: domain.EstadoBorrador,
	}

	err := roadmap.AddCurso(cursoBorrador)

	if err == nil {
		t.Fatalf("Se esperaba error de curso no publicado, se obtuvo nil")
	}
	if err != domain.ErrNotPublicCourse {
		t.Fatalf("Se esperaba estado %v, se obtuvo %v", domain.EstadoBorrador, err)
	}
}

func TestAddCursoDuplicateCourseError(t *testing.T) {
	cursoPublicado := domain.Curso{
		ID: "c-1", Titulo: "React 19", 
		Estado: domain.EstadoPublicado}
	
	roadmap := &domain.Roadmap{
		ID: "usr-1",
		Cursos:    []domain.Curso{cursoPublicado},
	}

	err := roadmap.AddCurso(cursoPublicado)

	if err != domain.ErrNotDuplicateCourse {
		t.Errorf("se esperaba error %v, se obtuvo: %v", domain.ErrNotDuplicateCourse, err)
	}
}