package repo

import (
	"context"

	"github.com/vasolovev/secret_santa/internal/entity"
)

type GroupRepo struct {
	groups []entity.Group
}

// New -.
func NewGroup(groups []entity.Group) *GroupRepo {
	return &GroupRepo{groups}
}

func (r *GroupRepo) Create(ctx context.Context, p entity.Group) error {
	panic("implementation error")
	return nil
}
func (r *GroupRepo) Read(ctx context.Context) ([]entity.Group, error) {
	panic("implementation error")
	return nil, nil
}
func (r *GroupRepo) Update(ctx context.Context, p entity.Group) error {
	panic("implementation error")
	return nil
}
func (r *GroupRepo) Delete(ctx context.Context, pID int) error {
	panic("implementation error")
	return nil
}
