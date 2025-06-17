package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "正数相加",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "负数相加",
			a:        -1,
			b:        -2,
			expected: -3,
		},
		{
			name:     "零相加",
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; 期望 %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
