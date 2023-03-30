package usecase

import (
	"context"

	"github.com/vasolovev/secret_santa/internal/entity"
	"github.com/vasolovev/secret_santa/internal/repo"
)

type ParticipantUseCase struct {
	repo repo.Participant
}

func NewParticipantUseCase(r repo.Participant) *ParticipantUseCase {
	return &ParticipantUseCase{
		repo: r,
	}
}

func (uc *ParticipantUseCase) Create(ctx context.Context, v entity.Participant) (uint, error) {
	return 0, nil
}
func (uc *ParticipantUseCase) GetAll(ctx context.Context) ([]entity.Participant, error) {
	return nil, nil
}
func (uc *ParticipantUseCase) Update(ctx context.Context, v entity.Participant) (uint, error) {
	return 0, nil
}
func (uc *ParticipantUseCase) Delete(ctx context.Context, id uint) error {
	return nil
}
