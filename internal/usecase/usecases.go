package usecase

import (
	"context"

	"github.com/vasolovev/secret_santa/internal/entity"
	"github.com/vasolovev/secret_santa/internal/repo"
)

type Groups interface {
	Create(context.Context, entity.Group) (uint, error)
	Update(context.Context, entity.Group) error
	GetByID(context.Context, []uint) ([]entity.Group, error)
	GetAll(context.Context) ([]entity.Group, error)
	Delete(context.Context, uint) error
}
type Participants interface {
	Create(context.Context, entity.Participant) (uint, error)
	Update(context.Context, entity.Participant) error
	GetByID(context.Context, []uint) ([]entity.Participant, error)
	GetAll(context.Context) ([]entity.Participant, error)
	Delete(context.Context, uint) error
}

type Deps struct {
	Repos *repo.Repositories
}
type UseCases struct {
	Participant Participants
	Group       Groups
}

func NewUsecases(deps Deps) *UseCases {
	participantUseCase := NewParticipantUseCase(deps.Repos.Participant)
	groupUseCase := NewGroupUseCase(deps.Repos.Group)

	return &UseCases{
		Participant: participantUseCase,
		Group:       groupUseCase,
	}
}
