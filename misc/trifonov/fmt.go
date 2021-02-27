package main

import (
	"fmt"
)

func main() {
	myvar := "operation"
	selectStatement := `
            SELECT role FROM abc INNER JOIN xyz ON (abc.name = %s) 
        `
	interpolated := fmt.Sprintf(selectStatement, myvar)
	fmt.Println(interpolated)

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	car := struct {
		Speed  int
		Weight string
	}{6, "tom"}

	fmt.Println("car:", car)
}
