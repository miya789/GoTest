package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
)

func testInterface() {
	var x, y interface{}

	fmt.Printf("%#v\n", x)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)

	x = 1
	y = []int{1, 2, 3}
	fmt.Printf("%#v\n", x)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)

	x = 2.1
	y = "hello"
	fmt.Printf("%#v\n", x)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)

	receive := 3
	x = receive

	if xi, ok := x.(int); ok {
		fmt.Println(xi * xi)
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other")
	}
}

func testJSON() {
	type Person struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Birthday string `json:"birthday"`
	}

	bytes, err := ioutil.ReadFile("test/sample.json")
	if err != nil {
		log.Fatal(err)
	}

	var persons []Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		log.Fatal(err)
	}

	for _, p := range persons {
		fmt.Printf("%d : %s %s\n", p.ID, p.Name, p.Birthday)
	}
}

type Car interface {
	run(int) string
	stop()
}

type MyCar struct {
	name  string
	speed int
}

func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}

func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

func testInterfaceClass() {
	myCar := &MyCar{
		name:  "aaa",
		speed: 101,
	}

	// var objCar Car = myCar
	objCar := myCar
	fmt.Printf("%T\n", objCar)
	fmt.Println(reflect.TypeOf(new(Car)))
	fmt.Println(reflect.TypeOf(new(MyCar)))
	fmt.Println(reflect.TypeOf(objCar))
	fmt.Println(objCar.run(50))
	objCar.stop()
}

// InvalidTestError is the struct for test.
type InvalidTestError struct {
	name  string
	count int
}

func (e *InvalidTestError) Error() string {
	return fmt.Sprintf("%d: test sample of InvalidTestError", e.count)
}

func testError() {
	// 1. errors.New(text string) error
	err1 := errors.New("1: test sample of errors.New()")
	{ // The difference of Println in built-in v.s. fmt
		fmt.Println(err1)
		fmt.Println(err1.Error())
		println(err1) // it returns address like "(0x50c540,0xc000082240)".
		println(err1.Error())
	}
	// 2. fmt.Errorf(fromat string, a ...interface{}) error
	//   予め定義したerrorで処理の分岐ができれば十分な場合用
	//   状態が持てないのが微妙そう
	//   (pkgとして取り出して使えるかは微妙？)
	err2 := fmt.Errorf("%d: %s", 2, "test sample of fmt.Errorf")
	{ // The difference of Println in built-in v.s. fmt
		fmt.Println(err2)
		fmt.Println(err2.Error())
		println(err2) // it returns address like "(0x50c540,0xc000082260)".
		println(err2.Error())
	}
	// 3. interface
	err3 := InvalidTestError{count: 3}
	{ // The difference of Println in built-in v.s. fmt
		fmt.Println(err3)
		fmt.Println(&err3) // it returns "{ 3}"
		fmt.Println(err3.Error())
		// println(err3) // syntax error
		println(err3.Error())
	}
	err4 := &InvalidTestError{count: 4}
	{ // The difference of Println in built-in v.s. fmt
		fmt.Println(err4)
		fmt.Println(err4.Error())
		println(err4) // it returns address.
		println(err4.Error())
	}

	var newErrorInterface interface{}
	invalidTestError := &InvalidTestError{name: "test"} // 適当なError
	newErrorInterface = invalidTestError

	{ // Check by type assertion
		r1, ok1 := err1.(error)
		fmt.Println(r1, ok1)
		r2, ok2 := err2.(error)
		fmt.Println(r2, ok2)
		// r3, ok3 := err3.(error)
		// fmt.Println(r3, ok3)
		// r4, ok4 := err4.(error)
		// fmt.Println(r4, ok4)

		if v, ok := newErrorInterface.(json.InvalidUnmarshalError); !ok || ok {
			fmt.Println(v, ok)
		}
		if v, ok := newErrorInterface.(*json.InvalidUnmarshalError); !ok || ok {
			fmt.Println(v, ok)
		}
		if v, ok := newErrorInterface.(InvalidTestError); !ok || ok {
			fmt.Println(v, ok)
		}
		if v, ok := newErrorInterface.(*InvalidTestError); !ok || ok {
			println(v, ok)
			fmt.Println(v, ok)
		}
	}

	{ // Print types
		fmt.Println(reflect.ValueOf(err1).Type())
		fmt.Println(reflect.ValueOf(invalidTestError).Type())
		fmt.Println(reflect.TypeOf(invalidTestError))
		fmt.Println(reflect.ValueOf(newErrorInterface).Type()) // interfaceとはバラさない
		fmt.Println(reflect.TypeOf(newErrorInterface))         // interfaceとはバラさない
	}
}

func testPointer() {
	// Instance without pointer
	var ptr MyCar
	fmt.Println(ptr)
	fmt.Println(&ptr)

	// Instance with null pointer
	var ptr2 *MyCar
	fmt.Println(ptr2)
	fmt.Println(&ptr2)

	// Empty instance with pointer
	ptr3 := &MyCar{}
	fmt.Println(ptr3)
	fmt.Println(&ptr3)

	// Not empty instance with pointer
	ptr4 := &MyCar{name: "1"}
	fmt.Println(ptr4)
	fmt.Println(&ptr4)
}

func main() {
	// testInterface()
	// testJSON()
	// testInterfaceClass()
	testError()
	testPointer()
}
