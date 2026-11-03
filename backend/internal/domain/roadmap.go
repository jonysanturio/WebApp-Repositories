package domain

import (
	"errors"
)

var (
	ErrNotDuplicateCourse = errors.New("No se puede duplicar el curso")
	ErrNotPublicCourse = errors.New("No se puede publicar el curso")
)

type Roadmap struct {
	ID			string
	Cursos		[]Curso
}

func (c *Roadmap) AddCurso(curso Curso) error {
	if curso.Estado != EstadoPublicado {
		return ErrNotPublicCourse
	}
	for _, c := range c.Cursos {
		if c.ID == curso.ID {
			return ErrNotDuplicateCourse
		}
	}
	c.Cursos = append(c.Cursos, curso)
	return nil
}
