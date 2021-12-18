package main

import (
	"github.com/zyedidia/rope"
)

type editor struct {
	buffer *rope.Node
	offset int
}

func newEditor(b []byte) *editor {
	return &editor{buffer: rope.New(b)}
}

func (e *editor) moveUp() {
	for i := e.offset - 1; i > 0; i-- {
		if e.buffer.At(i) == '\n' {
			e.offset = i
			break
		}
	}
}

func (e *editor) moveDown() {
	for i := e.offset + 1; i < e.buffer.Len(); i++ {
		if e.buffer.At(i) == '\n' {
			e.offset = i
			return
		}
	}
	e.offset = e.buffer.Len()
}

func (e *editor) moveLeft() {
	if e.offset > 0 {
		e.offset -= 1
	}
}

func (e *editor) moveRight() {
	if e.offset < e.buffer.Len() {
		e.offset += 1
	}
}

func (e *editor) insert(s string) {
	e.buffer.Insert(e.offset, []byte(s))
	e.offset += len(s)
}

func (e *editor) delete() {
	if e.buffer.Len() > 0 {
		e.buffer.Remove(e.offset, e.offset+1)
	}
}
