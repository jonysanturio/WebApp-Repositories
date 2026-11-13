package repository

import (
	"backend/internal/domain"
	"database/sql"
)

type ListFilters struct {
	Query     string
	Categoria string
	Page      int
	PageSize  int
}

type CursoRepository interface {
	Create(curso domain.Curso) (string, error)
	GetByID(id string) (*domain.Curso, error)
	List(filters ListFilters) ([]domain.Curso, int, error)
	Update(id string, updates map[string]interface{}) error
	Delete(id string) error
}

type postgresCursoRepository struct {
	db *sql.DB
}

func NewCursoRepository(db *sql.DB) CursoRepository {
	return &postgresCursoRepository{db: db}
}

func (r *postgresCursoRepository) Create(c domain.Curso) (string, error) {
	var id string
	query := `INSERT INTO cursos (titulo, estado) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, c.Titulo, c.Estado).Scan(&id)
	return id, err
}

func (r *postgresCursoRepository) GetByID(id string) (*domain.Curso, error) {
	query := `SELECT id, titulo, estado FROM cursos WHERE id = $1`
	var c domain.Curso
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Titulo, &c.Estado)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *postgresCursoRepository) List(filters ListFilters) ([]domain.Curso, int, error) {
	offset := (filters.Page - 1) * filters.PageSize
	query := `SELECT id, titulo, estado FROM cursos 
	          WHERE ($1 = '' OR LOWER(titulo) LIKE '%' || LOWER($1) || '%')
	          LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, filters.Query, filters.PageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cursos []domain.Curso
	for rows.Next() {
		var c domain.Curso
		if err := rows.Scan(&c.ID, &c.Titulo, &c.Estado); err != nil {
			return nil, 0, err
		}
		cursos = append(cursos, c)
	}

	var total int
	countQuery := `SELECT COUNT(*) FROM cursos WHERE ($1 = '' OR LOWER(titulo) LIKE '%' || LOWER($1) || '%')`
	_ = r.db.QueryRow(countQuery, filters.Query).Scan(&total)

	return cursos, total, nil
}

func (r *postgresCursoRepository) Update(id string, updates map[string]interface{}) error {
	query := `UPDATE cursos SET titulo = $1 WHERE id = $2`
	_, err := r.db.Exec(query, updates["titulo"], id)
	return err
}

func (r *postgresCursoRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM cursos WHERE id = $1`, id)
	return err
}
