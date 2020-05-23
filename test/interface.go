package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func main() {
	testInterface()
	testJSON()
}
