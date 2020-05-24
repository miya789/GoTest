package main

import (
	"encoding/json"
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

func main() {
	testInterface()
	testJSON()
	testInterfaceClass()
}