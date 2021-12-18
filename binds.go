package main

import "os"

type bindData struct {
	path string
	*editor
}

type keyDownResult struct {
	Offset       int    `json:"offset"`
	BufferString string `json:"bufferString"`
}

func (bd bindData) keyDown(key string) keyDownResult {
	if len(key) == 1 {
		bd.insert(key)
	} else {
		switch key {
		case "ArrowUp":
			bd.moveUp()
		case "ArrowDown":
			bd.moveDown()
		case "ArrowLeft":
			bd.moveLeft()
		case "ArrowRight":
			bd.moveRight()
		case "Backspace":
			bd.moveLeft()
			bd.delete()
		case "Enter":
			bd.insert("\n")
		case "Tab":
			bd.insert("\t")
		}
	}
	return keyDownResult{
		Offset:       bd.offset,
		BufferString: string(bd.buffer.Value()),
	}
}

func (bd bindData) ctrlKeyDown(key string) error {
	switch key {
	case "s":
		return os.WriteFile(bd.path, bd.buffer.Value(), 0777)
	}
	return nil
}

func (bd bindData) getBufferString() string {
	return string(bd.buffer.Value())
}
