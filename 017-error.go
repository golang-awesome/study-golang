package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func echo(name string, request string) (response string, err error) {
	if request == "" {
		//err = errors.New("empty content")
		err = fmt.Errorf("empty content for %s", name)
		return
	}

	response = fmt.Sprintf("echo: %s", request)
	return
}

// todo .(type)
func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func main() {
	for i, req := range []string{"", "hello"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(string(i), req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}

	printError := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}

		err = underlyingError(err)
		switch err {
		case os.ErrClosed:
			fmt.Println("error(closed)[%d]: %s\n", i, err)
		case os.ErrInvalid:
			fmt.Println("error(invalid)[%d]: %s\n", i, err)
		case os.ErrPermission:
			fmt.Println("error(permission)[%d]: %s\n", i, err)
		default:
			fmt.Println("error(unknown)[%d]: %s\n", i, err)
		}
	}

	// todo fix this
	var pathError os.PathError = os.PathError{Err: errors.New("path invalid error")}
	printError(10, pathError.Err)
}
