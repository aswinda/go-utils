package main

import (
	"fmt"
	"strconv"
	"github.com/aswinda/go-utils/common/gomath"
)
func main() {
	x := []float64{ 98, 93, 77, 82, 83 }
	fmt.Printf(strconv.FormatFloat(gomath.Sum(x), 'f', 6, 64))
}
