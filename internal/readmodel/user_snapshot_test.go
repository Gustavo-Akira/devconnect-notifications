package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserSnapshot(t *testing.T) {
	tests := []struct {
		name  string
		id    string
		email string
	}{
		{
			name:  "valid input",
			id:    "123",
			email: "user@example.com",
		},
		{
			name:  "empty id",
			id:    "",
			email: "user@example.com",
		},
		{
			name:  "empty email",
			id:    "123",
			email: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserSnapshot(tt.id, tt.name, tt.email)
			if tt.email == "" {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else if tt.id == "" {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else if tt.name == "" {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}
