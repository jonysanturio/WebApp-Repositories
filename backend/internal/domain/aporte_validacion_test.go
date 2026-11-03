package domain_test

import (
	"testing"
	"backend/internal/domain"
)

func TestValidarAporte_URLInvalida_Error(t *testing.T) {
	aporte := domain.Aporte{
		Titulo: "Curso de Go",
		URL:    "ftp://invalid-url.com",
	}

	err := domain.ValidarAporte(aporte)

	if err != domain.ErrURLInvalida {
		t.Errorf("se esperaba %v, se obtuvo %v", domain.ErrURLInvalida, err)
	}
}