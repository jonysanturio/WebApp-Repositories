package domain_test

import (
	"testing" 
	"backend/internal/domain"
)
func TestChangeEstadoTransitionInvalidInRewiewSuccess(t *testing.T)  {
	curso := &domain.Curso{
		ID:		"1",
		Titulo:	"Prueba",
		Estado: domain.EstadoBorrador,
	}

	err := curso.ChangeEstado(domain.EstadoEnRevision)

	if err != nil {
		t.Fatalf("Se espera error vacio, se obtuvo %v", err)
	}
	if curso.Estado != domain.EstadoEnRevision{
		t.Fatalf("Se esperaba estado %s, se obtuvo %s", domain.EstadoEnRevision, curso.Estado)
	}
}

func TestChangeStatusInvalidTransitionDraftToPublishedError(t *testing.T) {
	curso := &domain.Curso{
		ID:     "1",
		Titulo: "Prueba",
		Estado: domain.EstadoBorrador,
	}

	err := curso.ChangeEstado(domain.EstadoPublicado)

	if err == nil {
		t.Fatal("se esperaba un error de transición inválida, pero fue nil")
	}
	if err != domain.ErrTransitionInvalid {
		t.Errorf("se esperaba %v, se obtuvo: %v", domain.ErrTransitionInvalid, err)
	}
}
