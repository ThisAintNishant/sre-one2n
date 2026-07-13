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