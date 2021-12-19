package editor

import (
	"github.com/zyedidia/rope"
)

type Editor struct {
	buffer *rope.Node
	offset int
}

func New(b []byte) *Editor {
	return &Editor{buffer: rope.New(b)}
}

func (e *Editor) MoveUp() {
	for i := e.offset - 1; i > 0; i-- {
		if e.buffer.At(i) == '\n' {
			e.offset = i
			break
		}
	}
}

func (e *Editor) MoveDown() {
	for i := e.offset + 1; i < e.buffer.Len(); i++ {
		if e.buffer.At(i) == '\n' {
			e.offset = i
			return
		}
	}
	e.offset = e.buffer.Len()
}

func (e *Editor) MoveLeft() {
	if e.offset > 0 {
		e.offset -= 1
	}
}

func (e *Editor) MoveRight() {
	if e.offset < e.buffer.Len() {
		e.offset += 1
	}
}

func (e *Editor) Insert(s string) {
	e.buffer.Insert(e.offset, []byte(s))
	e.offset += len(s)
}

func (e *Editor) Delete() {
	if e.buffer.Len() > 0 {
		e.buffer.Remove(e.offset, e.offset+1)
	}
}

func (e *Editor) Offset() int {
	return e.offset
}

func (e *Editor) Buffer() []byte {
	return e.buffer.Value()
}
