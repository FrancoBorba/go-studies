package address

import "testing"

// UNIT TEST

func TestTypeAddress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"street address", "Street number 5", "street"},
		{"avenue address", "avenue Paulista", "avenue"},
		{"invalid address", "Square central", "Invalid Type"},
		{"empty string", "", "Invalid Type"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TypeAdress(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
