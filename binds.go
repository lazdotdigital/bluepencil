package main

import (
	"os"

	"github.com/lazdotdigital/bluepencil/editor"
)

type bindData struct {
	path string
	*editor.Editor
}

type keyDownResult struct {
	Offset       int    `json:"offset"`
	BufferString string `json:"bufferString"`
}

func (bd bindData) keyDown(key string) keyDownResult {
	if len(key) == 1 {
		bd.Insert(key)
	} else {
		switch key {
		case "ArrowUp":
			bd.MoveUp()
		case "ArrowDown":
			bd.MoveDown()
		case "ArrowLeft":
			bd.MoveLeft()
		case "ArrowRight":
			bd.MoveRight()
		case "Backspace":
			bd.MoveLeft()
			bd.Delete()
		case "Enter":
			bd.Insert("\n")
		case "Tab":
			bd.Insert("\t")
		}
	}
	return keyDownResult{
		Offset:       bd.Offset(),
		BufferString: string(bd.Buffer()),
	}
}

func (bd bindData) ctrlKeyDown(key string) error {
	switch key {
	case "s":
		return os.WriteFile(bd.path, bd.Buffer(), 0777)
	}
	return nil
}

func (bd bindData) getBufferString() string {
	return string(bd.Buffer())
}
