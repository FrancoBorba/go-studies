package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Rectangle", func(t *testing.T) {

		tests := []struct {
			name     string
			form     Form
			expected float64
		}{
			{
				name: "10x20 rectangle",
				form: Rectangle{
					Height: 10,
					Large:  20,
				},
				expected: 200,
			},
			{
				name: "1x5 rectangle",
				form: Rectangle{
					Height: 1,
					Large:  5,
				},
				expected: 5,
			},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				result := tt.form.Area()

				if result != tt.expected {
					t.Errorf("expected %.2f, got %.2f", tt.expected, result)
				}
			})
		}
	})

	t.Run("Circle", func(t *testing.T) {

		tests := []struct {
			name     string
			form     Form
			expected float64
		}{
			{
				name: "radius 1",
				form: Circle{
					Radius: 1,
				},
				expected: math.Pi,
			},
			{
				name: "radius 2",
				form: Circle{
					Radius: 2,
				},
				expected: math.Pi * 4,
			},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {

				result := tt.form.Area()

				if result != tt.expected {
					t.Errorf("expected %.4f, got %.4f", tt.expected, result)
				}
			})
		}
	})
}

