package main

import (
	"github.com/lazdotdigital/bluepencil/addon"
	"github.com/lazdotdigital/bluepencil/editor"
)

func makeExecutor(editor *editor.Editor) addon.Executor {
	return func(cmd string, fa addon.FunctionArgs) (interface{}, error) {
		switch cmd {
		case "moveUp":
			editor.MoveUp()
		case "moveDown":
			editor.MoveDown()
		case "moveLeft":
			editor.MoveLeft()
		case "moveRight":
			editor.MoveRight()
		case "delete":
			editor.Delete()
		case "currentChar":
			return editor.CurrentChar(), nil
		case "insert":
			text := fa["text"].(string)
			editor.Insert(text)
		}
		return nil, nil
	}
}
