package domain

import (
	"errors"
	"strings"
)

var (
	ErrURLInvalida     = errors.New("la URL debe comenzar con http:// o https://")
	ErrTituloVacio     = errors.New("el título no puede estar vacío")
)

type Aporte struct {
	Titulo string
	URL    string
}

func ValidarAporte(a Aporte) error {
	if strings.TrimSpace(a.Titulo) == "" {
		return ErrTituloVacio
	}
	if !strings.HasPrefix(a.URL, "http://") && !strings.HasPrefix(a.URL, "https://") {
		return ErrURLInvalida
	}
	return nil
}
