package repository

import (
	"context"

	"github.com/ThisAintNishant/sre-one2n/internal/models"
)

type StudentRepository interface {
	Create(ctx context.Context, student *models.Student) error
	GetAll(ctx context.Context) ([]models.Student, error)
	GetByID(ctx context.Context, id string) (*models.Student, error)
	Update(ctx context.Context, id string, student *models.Student) error
	Delete(ctx context.Context, id string) error
}