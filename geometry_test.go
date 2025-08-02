package geometry

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHypotenuse(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected string
	}{
		{0, 0, "0.00"},
		{0, 9, "9.00"},
		{3, 4, "5.00"},
		{10, 21, "23.26"},
		{56, 37, "67.12"},
		{102, 67, "122.04"},
		{34, 17, "38.01"},
		{3456, 1089, "3623.51"},
		{478, 201, "518.54"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("a=%f,b=%f", test.a, test.b), func(t *testing.T) {
			result := Hypotenuse(test.a, test.b)
			rounded := fmt.Sprintf("%.2f", result)
			assert.Equal(t, test.expected, rounded,
				"Для катетов %.2f и %.2f ожидалась гипотенуза %s, получено %s",
				test.a, test.b, test.expected, rounded)
		})
	}
}
