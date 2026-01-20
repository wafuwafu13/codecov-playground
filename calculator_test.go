package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed numbers", 5, -3, 2},
		{"zero", 0, 0, 0},
		{"decimal", 1.5, 2.5, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -5, -3, -2},
		{"mixed numbers", 5, -3, 8},
		{"zero", 0, 0, 0},
		{"decimal", 5.5, 2.5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative numbers", -5, -3, 15},
		{"mixed numbers", 5, -3, -15},
		{"zero", 0, 5, 0},
		{"decimal", 2.5, 4, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// func TestDivide(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		a        float64
// 		b        float64
// 		expected float64
// 	}{
// 		{"positive numbers", 10, 2, 5},
// 		{"negative numbers", -10, -2, 5},
// 		{"mixed numbers", 10, -2, -5},
// 		{"decimal result", 5, 2, 2.5},
// 		{"divide by zero", 10, 0, 0},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := Divide(tt.a, tt.b)
// 			if result != tt.expected {
// 				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
// 			}
// 		})
// 	}
// }
