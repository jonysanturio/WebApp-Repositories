package domain

import (
	"errors"
)

var (
	ErrTransitionInvalid = errors.New("transición del estado no permetida")
	ErrCampIncomplited = errors.New("El titulo y la URL son obligatorios")
)

type EstadoCurso string

const (
	EstadoBorrador 		EstadoCurso = "Borrador"
	EstadoEnRevision	EstadoCurso = "En_Revision"
	EstadoPublicado		EstadoCurso = "Publicado"
	EstadoRechazado		EstadoCurso = "Rechazado"
)

type Curso struct {
	ID			string
	Titulo		string
	URL			string
	Estado	EstadoCurso
}

func (c *Curso) ChangeEstado(nuevoEstado EstadoCurso) error {
	switch c.Estado {
	case EstadoBorrador:
		if nuevoEstado != EstadoEnRevision {
			return ErrTransitionInvalid
		}
	case EstadoEnRevision:
		if nuevoEstado != EstadoPublicado && nuevoEstado != EstadoRechazado {
			return ErrTransitionInvalid
		}
	case EstadoPublicado, EstadoRechazado:
		return ErrTransitionInvalid
	}
	c.Estado = nuevoEstado
	return nil
}