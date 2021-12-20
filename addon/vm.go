package addon

import (
	"fmt"

	"github.com/dop251/goja"
)

type FunctionArgs map[string]interface{}

type EventCallback func(FunctionArgs) error

type EventMap map[string]EventCallback

type Executor func(string, FunctionArgs) (interface{}, error)

func initVM(vm *goja.Runtime, events EventMap, executor Executor) error {
	if err := vm.Set("on", events.on); err != nil {
		return err
	}
	if err := vm.Set("execute", executor); err != nil {
		return err
	}
	if err := vm.Set("log", log); err != nil {
		return err
	}
	return nil
}

func (em EventMap) on(name string, cb EventCallback) {
	em[name] = cb
}

func log(s string) {
	fmt.Println(s)
}
