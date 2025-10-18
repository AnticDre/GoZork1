package main

import (
	"testing"
)

func TestObjectRespondTo(t *testing.T) {
	tests := []struct {
		name        string
		object      Object
		args        []string
		expected    bool
		description string
	}{
		{
			name: "exact name match",
			object: Object{
				name: "SWORD",
			},
			args:        []string{"SWORD"},
			expected:    true,
			description: "should match exact object name",
		},
		{
			name: "alias match",
			object: Object{
				name:    "SWORD",
				aliases: []string{"BLADE", "WEAPON"},
			},
			args:        []string{"BLADE"},
			expected:    true,
			description: "should match object alias",
		},
		{
			name: "adjective with name",
			object: Object{
				name:       "SWORD",
				adjectives: []string{"SHARP", "STEEL"},
			},
			args:        []string{"SHARP", "SWORD"},
			expected:    true,
			description: "should match adjective followed by name",
		},
		{
			name: "multiple adjectives with name",
			object: Object{
				name:       "SWORD",
				adjectives: []string{"SHARP", "STEEL"},
			},
			args:        []string{"SHARP", "STEEL", "SWORD"},
			expected:    true,
			description: "should match multiple adjectives followed by name",
		},
		{
			name: "no match",
			object: Object{
				name: "SWORD",
			},
			args:        []string{"SHIELD"},
			expected:    false,
			description: "should not match different object",
		},
		{
			name: "case insensitive match",
			object: Object{
				name: "SWORD",
			},
			args:        []string{"sword"}, // input will be converted to uppercase in actual game
			expected:    false,             // since RespondTo expects uppercase
			description: "should handle case sensitivity correctly",
		},
		{
			name: "empty args",
			object: Object{
				name: "SWORD",
			},
			args:        []string{},
			expected:    false,
			description: "should handle empty arguments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.object.RespondTo(tt.args)
			if result != tt.expected {
				t.Errorf("%s: expected %v, got %v", tt.description, tt.expected, result)
			}
		})
	}
}

func TestObjectGetName(t *testing.T) {
	tests := []struct {
		name        string
		object      Object
		expected    string
		description string
	}{
		{
			name: "simple object",
			object: Object{
				name: "sword",
			},
			expected:    "a sword",
			description: "should return name with article",
		},
		{
			name: "openable closed object",
			object: Object{
				name:     "chest",
				openable: true,
				open:     false,
			},
			expected:    "a chest (closed)",
			description: "should show closed status for openable objects",
		},
		{
			name: "openable open object",
			object: Object{
				name:     "chest",
				openable: true,
				open:     true,
			},
			expected:    "a chest (open)",
			description: "should show open status for openable objects",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.object.GetName()
			if result != tt.expected {
				t.Errorf("%s: expected '%s', got '%s'", tt.description, tt.expected, result)
			}
		})
	}
}

func TestObjectGetDesc(t *testing.T) {
	tests := []struct {
		name        string
		object      Object
		expected    string
		description string
	}{
		{
			name: "object with description only",
			object: Object{
				name: "sword",
				desc: "A sharp metal blade.",
			},
			expected:    "A sharp metal blade.",
			description: "should return object description",
		},
		{
			name: "openable closed object with description",
			object: Object{
				name:     "chest",
				desc:     "A wooden treasure chest.",
				openable: true,
				open:     false,
			},
			expected:    "A wooden treasure chest.\nThe chest is closed.",
			description: "should return description with closed status",
		},
		{
			name: "openable open object with description",
			object: Object{
				name:     "chest",
				desc:     "A wooden treasure chest.",
				openable: true,
				open:     true,
			},
			expected:    "A wooden treasure chest.\nThe chest is open.",
			description: "should return description with open status",
		},
		{
			name: "openable object without description",
			object: Object{
				name:     "chest",
				openable: true,
				open:     false,
			},
			expected:    "The chest is closed.",
			description: "should return only status for openable objects without description",
		},
		{
			name: "empty description",
			object: Object{
				name: "rock",
				desc: "",
			},
			expected:    "",
			description: "should handle empty description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.object.GetDesc()
			if result != tt.expected {
				t.Errorf("%s: expected '%s', got '%s'", tt.description, tt.expected, result)
			}
		})
	}
}

func TestObjectContainerAddObject(t *testing.T) {
	container := ObjectContainer{}
	obj1 := &Object{name: "sword"}
	obj2 := &Object{name: "shield"}

	// Test adding single object
	container.AddObject(obj1)
	if len(container.objects) != 1 {
		t.Errorf("Expected 1 object after adding one, got %d", len(container.objects))
	}
	if container.objects[0] != obj1 {
		t.Errorf("Expected first object to be sword")
	}

	// Test adding multiple objects
	container.AddObject(obj2)
	if len(container.objects) != 2 {
		t.Errorf("Expected 2 objects after adding two, got %d", len(container.objects))
	}
}

func TestObjectContainerRemoveObject(t *testing.T) {
	obj1 := &Object{name: "sword"}
	obj2 := &Object{name: "shield"}
	obj3 := &Object{name: "potion"}

	container := ObjectContainer{
		objects: []*Object{obj1, obj2, obj3},
	}

	// Remove middle object
	container.RemoveObject(obj2)
	if len(container.objects) != 2 {
		t.Errorf("Expected 2 objects after removing one, got %d", len(container.objects))
	}
	if container.objects[0] != obj1 || container.objects[1] != obj3 {
		t.Errorf("Expected remaining objects to be sword and potion")
	}

	// Remove first object
	container.RemoveObject(obj1)
	if len(container.objects) != 1 {
		t.Errorf("Expected 1 object after removing another, got %d", len(container.objects))
	}
	if container.objects[0] != obj3 {
		t.Errorf("Expected remaining object to be potion")
	}
}

func TestObjectContainerFindObject(t *testing.T) {
	obj1 := &Object{name: "SWORD", aliases: []string{"BLADE"}}
	obj2 := &Object{name: "SHIELD"}

	container := ObjectContainer{
		objects: []*Object{obj1, obj2},
	}

	// Test finding by name
	found := container.FindObject([]string{"SWORD"})
	if found != obj1 {
		t.Errorf("Expected to find sword object")
	}

	// Test finding by alias
	found = container.FindObject([]string{"BLADE"})
	if found != obj1 {
		t.Errorf("Expected to find sword object by alias")
	}

	// Test not finding object
	found = container.FindObject([]string{"POTION"})
	if found != nil {
		t.Errorf("Expected not to find non-existent object")
	}

	// Test empty container
	emptyContainer := ObjectContainer{}
	found = emptyContainer.FindObject([]string{"SWORD"})
	if found != nil {
		t.Errorf("Expected not to find object in empty container")
	}
}

func TestObjectContainerObjectNames(t *testing.T) {
	tests := []struct {
		name        string
		objects     []*Object
		expectedRes string
		expectError bool
		description string
	}{
		{
			name: "single object",
			objects: []*Object{
				{name: "sword"},
			},
			expectedRes: "a sword",
			expectError: false,
			description: "should return single object name",
		},
		{
			name: "two objects",
			objects: []*Object{
				{name: "sword"},
				{name: "shield"},
			},
			expectedRes: "a sword and a shield",
			expectError: false,
			description: "should return two objects with 'and'",
		},
		{
			name: "three objects",
			objects: []*Object{
				{name: "sword"},
				{name: "shield"},
				{name: "potion"},
			},
			expectedRes: "a sword, a shield and a potion",
			expectError: false,
			description: "should return three objects with commas and 'and'",
		},
		{
			name: "fixture objects only",
			objects: []*Object{
				{name: "table", fixture: true},
				{name: "chair", fixture: true},
			},
			expectedRes: "",
			expectError: true,
			description: "should return error when only fixture objects present",
		},
		{
			name: "mixed fixture and non-fixture",
			objects: []*Object{
				{name: "table", fixture: true},
				{name: "sword"},
			},
			expectedRes: "a sword",
			expectError: false,
			description: "should only return non-fixture objects",
		},
		{
			name:        "empty container",
			objects:     []*Object{},
			expectedRes: "",
			expectError: true,
			description: "should return error for empty container",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := ObjectContainer{objects: tt.objects}
			result, err := container.ObjectNames()

			if tt.expectError && err == nil {
				t.Errorf("%s: expected error but got none", tt.description)
			}
			if !tt.expectError && err != nil {
				t.Errorf("%s: expected no error but got: %v", tt.description, err)
			}
			if result != tt.expectedRes {
				t.Errorf("%s: expected '%s', got '%s'", tt.description, tt.expectedRes, result)
			}
		})
	}
}
