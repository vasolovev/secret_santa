package repo

import (
	"context"

	"github.com/vasolovev/secret_santa/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test
type Repositories struct {
	Participant Participant
	Group       Group
}

func NewRepositories(groups []entity.Group, participants []entity.Participant) *Repositories {
	return &Repositories{
		Participant: NewParticipant(participants),
		Group:       NewGroup(groups),
	}
}

type Participant interface {
	Create(context.Context, entity.Participant) error
	Read(context.Context) ([]entity.Participant, error)
	Update(context.Context, entity.Participant) error
	Delete(context.Context, int) error
}

type Group interface {
	Create(context.Context, entity.Group) error
	Read(context.Context) ([]entity.Group, error)
	Update(context.Context, entity.Group) error
	Delete(context.Context, int) error
}
