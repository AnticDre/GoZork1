package main

import (
	"testing"
)

func TestTrollAIBasics(t *testing.T) {
	// Create test setup
	_, trollRoom := NewGameWorld()
	player := &Player{
		console:   nil,
		room:      trollRoom,
		maxPoints: 11,
		points:    0,
	}

	var trollAI TrollAI

	// Test initialization
	trollAI.Init(trollRoom, player)

	if trollAI.room != trollRoom {
		t.Errorf("TrollAI room should be set to troll room")
	}

	if trollAI.player != player {
		t.Errorf("TrollAI player should be set to the player")
	}

	if trollAI.aggro != trollDifficulty {
		t.Errorf("TrollAI aggro should be initialized to %d, got %d", trollDifficulty, trollAI.aggro)
	}

	// Check that troll object was added to room
	troll := trollRoom.FindObject([]string{"TROLL"})
	if troll == nil {
		t.Errorf("Troll object should be added to troll room during initialization")
	}

	if troll.carryable {
		t.Errorf("Troll should not be carryable")
	}
}

func TestTrollAIFollowLogic(t *testing.T) {
	_, trollRoom := NewGameWorld()
	player := &Player{
		console:   nil,
		room:      trollRoom,
		maxPoints: 11,
		points:    0,
	}

	var trollAI TrollAI
	trollAI.Init(trollRoom, player)

	// Test initial state - should not be following
	if trollAI.follow {
		t.Errorf("TrollAI should not be following initially")
	}

	// Test starting to follow when player enters troll room
	player.room = trollRoom
	trollAI.PlayerMove()

	if !trollAI.follow {
		t.Errorf("TrollAI should start following when player enters troll room")
	}
}

func TestTrollDifficultyConstant(t *testing.T) {
	// Test that the difficulty constant is reasonable
	if trollDifficulty <= 0 {
		t.Errorf("Troll difficulty should be positive, got %d", trollDifficulty)
	}

	if trollDifficulty > 10 {
		t.Errorf("Troll difficulty seems too high: %d", trollDifficulty)
	}
}
