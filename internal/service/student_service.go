package service

import (
	"context"

	"github.com/ThisAintNishant/sre-one2n/internal/models"
	"github.com/ThisAintNishant/sre-one2n/internal/repository"
)

type StudentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{
		repo: repo,
	}
}

func (s *StudentService) Create(ctx context.Context, student *models.Student) error {
	return s.repo.Create(ctx, student)
}

func (s *StudentService) GetAll(ctx context.Context) ([]models.Student, error) {
	return s.repo.GetAll(ctx)
}

func (s *StudentService) GetByID(ctx context.Context, id string) (*models.Student, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *StudentService) Update(ctx context.Context, id string, student *models.Student) error {
	return s.repo.Update(ctx, id, student)
}

func (s *StudentService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}