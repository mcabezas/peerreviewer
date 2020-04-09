package member

import "github.com/google/uuid"

type Category int

const (
	NeedsCoaching = iota
	Experienced
)

type Member struct {
	ID       string
	Name     string
	Email    string
	Category Category
}

type MemberOption func(*Member)

func NewMember(email string, opts ...MemberOption) *Member {
	member := &Member{ID: uuid.New().String(),  Email:email}
	for _, opt := range opts {
		opt(member)
	}
	return member
}

func WithName(name string) MemberOption {
	return func(m *Member) {
		m.Name = name
	}
}

func WithCategory(category Category) MemberOption {
	return func(m *Member) {
		m.Category = category
	}
}
