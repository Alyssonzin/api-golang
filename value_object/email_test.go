package value_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailIsValid(t *testing.T) {
	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{name: "email valido", value: "user@example.com", valid: true},
		{name: "email invalido", value: "user.example.com", valid: false},
		{name: "email vazio", value: "", valid: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := &Email{value: tt.value}
			assert.Equal(t, tt.valid, email.IsValid())
		})
	}
}

func TestNewEmailPanicsOnInvalid(t *testing.T) {
	assert.Panics(t, func() {
		_ = NewEmail("invalid-email")
	})
}

func TestNewEmailDoesNotPanicOnValid(t *testing.T) {
	assert.NotPanics(t, func() {
		_ = NewEmail("user@example.com")
	})
}
