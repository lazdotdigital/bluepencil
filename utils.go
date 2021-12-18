package main

import (
	"errors"
	"os"
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
