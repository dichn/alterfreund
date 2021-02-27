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

}

