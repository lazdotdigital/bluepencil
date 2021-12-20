package main

import (
	"errors"
	"os"

	"github.com/lazdotdigital/bluepencil/addon"
)

func readFileFromArgs() ([]byte, string, error) {
	if len(os.Args) != 2 {
		return nil, "", errors.New("invalid number of arguments")
	}
	path := os.Args[1]
	b, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, "", err
	}
	return b, path, nil
}

func loadAddonsFromEnv(executor addon.Executor) (as []addon.Addon, err error) {
	path := os.Getenv("ADDONS_PATH")
	if path == "" {
		return
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, f := range files {
		a, err := addon.New(path+f.Name(), executor)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	return
}
