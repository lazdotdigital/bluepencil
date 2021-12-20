package addon

import (
	"errors"
	"os"

	"github.com/dop251/goja"
)

type Addon struct {
	vm     *goja.Runtime
	events EventMap
}

var ErrEventNotFound = errors.New("event not found")

func New(path string, executor Executor) (Addon, error) {
	script, err := os.ReadFile(path)
	if err != nil {
		return Addon{}, err
	}
	vm := goja.New()
	events := make(map[string]EventCallback)
	initVM(vm, events, executor)
	if _, err := vm.RunString(string(script)); err != nil {
		return Addon{}, err
	}
	return Addon{vm, events}, nil
}

func (a Addon) Fire(eventName string, args FunctionArgs) error {
	event, ok := a.events[eventName]
	if !ok {
		return ErrEventNotFound
	}
	return event(args)
}
