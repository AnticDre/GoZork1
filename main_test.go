package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestByLength(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "sort strings by length ascending",
			input:    []string{"hello", "go", "programming", "a"},
			expected: []string{"a", "go", "hello", "programming"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "single element",
			input:    []string{"test"},
			expected: []string{"test"},
		},
		{
			name:     "equal length strings",
			input:    []string{"cat", "dog", "bat"},
			expected: []string{"cat", "dog", "bat"}, // order should be preserved for equal lengths
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]string, len(tt.input))
			copy(result, tt.input)
			sort.Sort(ByLength(result))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ByLength sort failed. Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestByLengthMethods(t *testing.T) {
	strings := ByLength([]string{"hello", "go", "test"})

	// Test Len method
	if strings.Len() != 3 {
		t.Errorf("Expected Len() to return 3, got %d", strings.Len())
	}

	// Test Less method
	if !strings.Less(1, 0) { // "go" (index 1) should be less than "hello" (index 0)
		t.Errorf("Expected Less(1, 0) to return true")
	}
	if strings.Less(0, 1) { // "hello" (index 0) should not be less than "go" (index 1)
		t.Errorf("Expected Less(0, 1) to return false")
	}

	// Test Swap method
	original := make([]string, len(strings))
	copy(original, strings)
	strings.Swap(0, 1)
	if strings[0] != original[1] || strings[1] != original[0] {
		t.Errorf("Swap method failed. Expected elements to be swapped")
	}
}
