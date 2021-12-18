package main

import (
	"fmt"
	"testing"
)

type offsetTest struct {
	initial, expected int
}

func TestMoveUp(t *testing.T) {
	editor := newEditor([]byte("this is a\ntest"))
	tests := []offsetTest{{11, 9}, {3, 3}}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("MoveUp: %v", i), func(t *testing.T) {
			editor.offset = tt.initial
			editor.moveUp()
			if editor.offset != tt.expected {
				t.Errorf("editor.offset: %v, expected: %v", editor.offset, tt.expected)
			}
		})
	}
}

func TestMoveDown(t *testing.T) {
	editor := newEditor([]byte("this is a\ntest"))
	tests := []offsetTest{{3, 9}, {11, 11}}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("MoveDown: %v", i), func(t *testing.T) {
			editor.offset = tt.initial
			editor.moveDown()
			if editor.offset != tt.expected {
				t.Errorf("editor.offset: %v, expected: %v", editor.offset, tt.expected)
			}
		})
	}
}

func TestMoveLeft(t *testing.T) {
	editor := newEditor([]byte("this is a\ntest"))
	tests := []offsetTest{{3, 2}, {1, 0}, {0, 0}}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("MoveLeft: %v", i), func(t *testing.T) {
			editor.offset = tt.initial
			editor.moveLeft()
			if editor.offset != tt.expected {
				t.Errorf("editor.offset: %v, expected: %v", editor.offset, tt.expected)
			}
		})
	}
}

func TestMoveRight(t *testing.T) {
	editor := newEditor([]byte("this is a\ntest"))
	tests := []offsetTest{{0, 1}, {13, 13}}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("MoveRight: %v", i), func(t *testing.T) {
			editor.offset = tt.initial
			editor.moveRight()
			if editor.offset != tt.expected {
				t.Errorf("editor.offset: %v, expected: %v", editor.offset, tt.expected)
			}
		})
	}
}
