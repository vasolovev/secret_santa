package usecase

import (
	"context"

	"github.com/vasolovev/secret_santa/internal/entity"
	"github.com/vasolovev/secret_santa/internal/repo"
)

type GroupUseCase struct {
	repo repo.Group
}

func NewGroupUseCase(r repo.Group) *GroupUseCase {
	return &GroupUseCase{
		repo: r,
	}
}

func (uc *GroupUseCase) Create(ctx context.Context, v entity.Group) (uint, error) {
	return 0, nil
}
func (uc *GroupUseCase) GetAll(ctx context.Context) ([]entity.Group, error) {
	return nil, nil
}
func (uc *GroupUseCase) Update(ctx context.Context, v entity.Group) (uint, error) {
	return 0, nil
}
func (uc *GroupUseCase) Delete(ctx context.Context, id uint) error {
	return nil
}
