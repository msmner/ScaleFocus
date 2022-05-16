package main

import (
	"errors"
	"fmt"
)

func main() {
	a := SafeExec(noError)
	err := a()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	a = SafeExec(withError)
	err = a()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	a = SafeExec(withPanic)
	err = a()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}

type Action func() error

func SafeExec(a Action) Action {
	return a
}

func noError() error {
	return fmt.Errorf("no error")
}

func withError() error {
	return errors.New("this is a wrapped error")
}

func withPanic() error {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("safe exec: %s", r)
			fmt.Printf("error with panic: %s", err)
		}
	}()

	panic("panics")
}
