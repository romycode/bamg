package models

import "testing"

func TestNewIban(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Return new IBAN",
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIban(); len(got) != tt.want {
				t.Errorf("NewIban() = %v, want %v", got, tt.want)
			}
		})
	}
}
