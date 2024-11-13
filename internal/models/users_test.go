package models

import (
	"testing"
	"snippetbox.nijat.net/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	tests := []struct {
		name	string
		userID	int
		want 	bool
	}{
		{
			name: "ValidID",
			userID: 1,
			want: true,
		},
		{
			name: "Zero ID",
			userID: 0,
			want: false,
		},
		{
			name: "Not exixtent ID",
			userID: 2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}
			exists, err := m.Exists(tt.userID)

			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}