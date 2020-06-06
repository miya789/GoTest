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

// wrapするには%w
func TestIs() {
	err := simpleError()
	fmt.Printf("simpleError():\t %T\n", err)

	err = wrappedError()
	fmt.Printf("wrappedError():\t %T\n", err)
	// switchでは捉えられない
	if err != nil {
		switch err {
		case MyError:
			fmt.Println("MyError:", err)
		default:
			fmt.Println("default", err)
		}
	}

	// errors.Is()では捉えられる
	fmt.Printf("%T\n", err)
	if errors.Is(err, MyError) {
		fmt.Println("err")
	}
}

func main() {
	// TestWrap()
	TestIs()
}
