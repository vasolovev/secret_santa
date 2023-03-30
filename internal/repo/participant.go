package repo

import (
	"context"

	"github.com/vasolovev/secret_santa/internal/entity"
)

type ParticipantRepo struct {
	participants []entity.Participant
}

// New -.
func NewParticipant(participants []entity.Participant) *ParticipantRepo {
	return &ParticipantRepo{participants}
}

func (r *ParticipantRepo) Create(ctx context.Context, p entity.Participant) error {
	panic("implementation error")
	return nil
}
func (r *ParticipantRepo) Read(ctx context.Context) ([]entity.Participant, error) {
	panic("implementation error")
	return nil, nil
}
func (r *ParticipantRepo) Update(ctx context.Context, p entity.Participant) error {
	panic("implementation error")
	return nil
}
func (r *ParticipantRepo) Delete(ctx context.Context, pID int) error {
	panic("implementation error")
	return nil
}
