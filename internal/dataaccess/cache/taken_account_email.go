package cache

import "context"

type TakenAccountEmail interface {
	Add(ctx context.Context, email string) error
	Has(ctx context.Context, email string) (bool, error)
}

type takenAccountEmail struct {
	client map[string]bool
}

func NewTakenAccountEmail() TakenAccountEmail {
	return &takenAccountEmail{
		client: make(map[string]bool),
	}
}

func (t *takenAccountEmail) Add(ctx context.Context, email string) error {
	t.client[email] = true
	return nil
}

func (t *takenAccountEmail) Has(ctx context.Context, email string) (bool, error) {
	return t.client[email], nil
}
