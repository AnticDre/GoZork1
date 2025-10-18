package main

import (
	"testing"
)

func TestRoomExitDirection(t *testing.T) {
	// Create test rooms
	northRoom := &Room{name: "North Room"}
	southRoom := &Room{name: "South Room"}
	eastRoom := &Room{name: "East Room"}
	westRoom := &Room{name: "West Room"}
	upRoom := &Room{name: "Up Room"}
	downRoom := &Room{name: "Down Room"}
	inRoom := &Room{name: "In Room"}
	outRoom := &Room{name: "Out Room"}

	// Create main room with all exits
	room := &Room{
		name: "Test Room",
		n:    northRoom,
		s:    southRoom,
		e:    eastRoom,
		w:    westRoom,
		up:   upRoom,
		down: downRoom,
		in:   inRoom,
		out:  outRoom,
	}

	tests := []struct {
		direction string
		expected  *Room
	}{
		{"NORTH", northRoom},
		{"SOUTH", southRoom},
		{"EAST", eastRoom},
		{"WEST", westRoom},
		{"UP", upRoom},
		{"DOWN", downRoom},
		{"IN", inRoom},
		{"OUT", outRoom},
		{"INVALID", nil},
		{"", nil},
	}

	for _, tt := range tests {
		t.Run(tt.direction, func(t *testing.T) {
			result := room.ExitDirection(tt.direction)
			if result != tt.expected {
				t.Errorf("ExitDirection(%s): expected %v, got %v", tt.direction, tt.expected, result)
			}
		})
	}
}

func TestRoomExitDirectionWithExitFunc(t *testing.T) {
	northRoom := &Room{name: "North Room"}

	// Room with exit function that blocks NORTH direction
	room := &Room{
		name: "Test Room",
		n:    northRoom,
		exitFunc: func(dir string) bool {
			return dir != "NORTH" // Block north direction
		},
	}

	// Test blocked direction
	result := room.ExitDirection("NORTH")
	if result != nil {
		t.Errorf("Expected NORTH to be blocked, but got room: %v", result)
	}

	// Test allowed direction (should return nil since no south room exists)
	result = room.ExitDirection("SOUTH")
	if result != nil {
		t.Errorf("Expected SOUTH to return nil (no room), but got: %v", result)
	}
}

func TestRoomEnter(t *testing.T) {
	room := &Room{name: "Test Room", visited: false}

	if room.visited {
		t.Errorf("Room should not be visited initially")
	}

	room.Enter()

	if !room.visited {
		t.Errorf("Room should be marked as visited after Enter()")
	}
}

func TestRoomLeave(t *testing.T) {
	room := &Room{name: "Test Room"}

	// Leave method exists but doesn't do anything currently
	// This test ensures it doesn't panic
	room.Leave()

	// Test passes if no panic occurs
}

// Integration test for NewGameWorld
func TestNewGameWorldBasicStructure(t *testing.T) {
	startRoom, trollRoom := NewGameWorld()

	if startRoom == nil {
		t.Fatal("NewGameWorld should return a non-nil start room")
	}

	if trollRoom == nil {
		t.Fatal("NewGameWorld should return a non-nil troll room")
	}

	if startRoom.name != "\nWest of House" {
		t.Errorf("Expected start room to be 'West of House', got '%s'", startRoom.name)
	}

	if trollRoom.name != "Troll Room" {
		t.Errorf("Expected troll room to be 'Troll Room', got '%s'", trollRoom.name)
	}
}

func TestNewGameWorldConnections(t *testing.T) {
	startRoom, _ := NewGameWorld()

	// Test some basic connections from West of House
	if startRoom.n == nil {
		t.Errorf("West of House should have a north exit")
	}

	if startRoom.s == nil {
		t.Errorf("West of House should have a south exit")
	}

	if startRoom.e != nil {
		t.Errorf("West of House should not have an east exit")
	}

	// Test that north room has correct name
	if startRoom.n.name != "\nNorth of House" {
		t.Errorf("North room should be 'North of House', got '%s'", startRoom.n.name)
	}
}

func TestNewGameWorldObjects(t *testing.T) {
	startRoom, trollRoom := NewGameWorld()

	// Navigate to kitchen to test objects
	bhouse := startRoom.n.e // North of House -> Behind House
	if bhouse == nil {
		t.Fatal("Could not navigate to Behind House")
	}

	kitchen := bhouse.in // Behind House -> Kitchen
	if kitchen == nil {
		t.Fatal("Could not navigate to Kitchen")
	}

	// Check for window object in kitchen
	window := kitchen.FindObject([]string{"WINDOW"})
	if window == nil {
		t.Errorf("Kitchen should contain a window object")
	}

	// Check for can object in kitchen
	can := kitchen.FindObject([]string{"CAN"})
	if can == nil {
		t.Errorf("Kitchen should contain a can object")
	}

	// Check troll room has troll (will be added by TrollAI.Init, but room should exist)
	if trollRoom.name != "Troll Room" {
		t.Errorf("Troll room should have correct name")
	}
}

func TestRoomObjectContainer(t *testing.T) {
	room := &Room{name: "Test Room"}
	obj := &Object{name: "TEST OBJECT"}

	// Test adding object to room
	room.AddObject(obj)

	found := room.FindObject([]string{"TEST", "OBJECT"})
	if found == nil {
		t.Errorf("Should find object that was added to room")
	}

	// Test removing object from room
	room.RemoveObject(obj)

	found = room.FindObject([]string{"TEST", "OBJECT"})
	if found != nil {
		t.Errorf("Should not find object after it was removed from room")
	}
}
