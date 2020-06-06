package main

import (
	"errors"
	"fmt"
)

type ErrChild string

func (e *ErrChild) Error() string {
	return fmt.Sprintf("This is child error.")
}

type ErrParent struct {
	err error
}

func (e *ErrParent) Error() string {
	return fmt.Sprintf("This is parent error.")
}

func (e *ErrParent) Unwrap() string {
	return fmt.Sprintf("This is parent error and include %+v.", e.err)
}

func TestWrap() {
	err := &ErrParent{
		err: new(ErrChild),
	}
	fmt.Println(err.Error())
	fmt.Println(err.Unwrap())
	fmt.Println(errors.Unwrap(err))
}

var (
	MyError    = myError()
	OtherError = otherError()
)

func myError() error    { return errors.New("myErr") }
func otherError() error { return errors.New("others") }

func simpleError() error {
	// return OtherError
	return MyError
}

func wrappedError() error {
	err := simpleError()
	return fmt.Errorf("%w", err)

}

func TestIs() {
	err := wrappedError()
	if err != nil {
		switch err {
		case MyError:
			fmt.Println("MyError:", err)
		default:
			fmt.Println("default", err)
		}
	}

	if errors.Is(err, MyError) {
		fmt.Println("err")
	}
}

func main() {
	// TestWrap()
	TestIs()
}
