package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type CursoLister interface {
	List(filters repository.ListFilters) ([]domain.Curso, int, error)
}

type CursoService struct {
	repo CursoLister
}

func NewCursoService(repo CursoLister) *CursoService {
	return &CursoService{repo: repo}
}

func (s *CursoService) Listar(
	page int,
	size int,
	query string,
) ([]domain.Curso, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	return s.repo.List(repository.ListFilters{
		Query:    query,
		Page:     page,
		PageSize: size,
	})
}
