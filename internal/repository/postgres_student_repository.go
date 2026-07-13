package repository

import (
	"context"

	"github.com/ThisAintNishant/sre-one2n/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStudentRepository struct {
	db *pgxpool.Pool
}

func NewPostgresStudentRepository(db *pgxpool.Pool) *PostgresStudentRepository {
	return &PostgresStudentRepository{
		db: db,
	}
}

func (r *PostgresStudentRepository) Create(ctx context.Context, student *models.Student) error {

	query := `
	INSERT INTO students
	(first_name, last_name, email, age)
	VALUES ($1,$2,$3,$4)
	RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		student.FirstName,
		student.LastName,
		student.Email,
		student.Age,
	).Scan(
		&student.ID,
		&student.CreatedAt,
		&student.UpdatedAt,
	)
}

func (r *PostgresStudentRepository) GetAll(ctx context.Context) ([]models.Student, error) {
	return nil, nil
}

func (r *PostgresStudentRepository) GetByID(ctx context.Context, id string) (*models.Student, error) {
	return nil, nil
}

func (r *PostgresStudentRepository) Update(ctx context.Context, id string, student *models.Student) error {
	return nil
}

func (r *PostgresStudentRepository) Delete(ctx context.Context, id string) error {
	return nil
}