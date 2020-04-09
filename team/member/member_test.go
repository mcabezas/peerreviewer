package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMember(t *testing.T) {
	member := NewMember("Marce", WithCategory(Experienced), WithName("Marcelon"))
	assert.EqualValues(t, member.Category, Experienced, "category option")
	assert.EqualValues(t, member.Name, "Marcelon", "name option")
	member2 := NewMember("Marce", WithCategory(NeedsCoaching), WithName("Marcelon"))
	assert.NotEqual(t, member2.Category, Experienced, "category option")
}
