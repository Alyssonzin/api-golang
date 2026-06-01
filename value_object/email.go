package value_object

import (
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) *Email {
	email := &Email{value: value}
	if !email.IsValid() {
		panic("Invalid email address: " + value)
	}
	return email
}

func (e *Email) IsValid() bool {
	containsAt := strings.Contains(e.value, "@")
	containsDot := strings.Contains(e.value, ".")
	return containsAt && containsDot
}
