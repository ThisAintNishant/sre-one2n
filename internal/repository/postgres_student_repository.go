package repository

import (
	"context"

	"github.com/ThisAintNishant/sre-one2n/internal/models"
	"github.com/jackc/pgx/v5"
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

	query := `
	SELECT
		id,
		first_name,
		last_name,
		email,
		age,
		created_at,
		updated_at
	FROM students
	ORDER BY created_at DESC;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var student models.Student

		err := rows.Scan(
			&student.ID,
			&student.FirstName,
			&student.LastName,
			&student.Email,
			&student.Age,
			&student.CreatedAt,
			&student.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *PostgresStudentRepository) GetByID(ctx context.Context, id string) (*models.Student, error) {

	query := `
	SELECT
		id,
		first_name,
		last_name,
		email,
		age,
		created_at,
		updated_at
	FROM students
	WHERE id = $1;
	`

	var student models.Student

	err := r.db.QueryRow(ctx, query, id).Scan(
		&student.ID,
		&student.FirstName,
		&student.LastName,
		&student.Email,
		&student.Age,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *PostgresStudentRepository) Update(ctx context.Context, id string, student *models.Student) error {

	query := `
	UPDATE students
	SET first_name = $1, last_name = $2, email = $3, age = $4, updated_at = NOW()
	WHERE id = $5
	RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		student.FirstName,
		student.LastName,
		student.Email,
		student.Age,
		id,
	).Scan(
		&student.ID,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresStudentRepository) Delete(ctx context.Context, id string) error {

	query := `DELETE FROM students WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}