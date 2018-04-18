package main

import (
	"fmt"
	"math"

	"github.com/lybecker/GolangPlayground/countries"

	"github.com/lybecker/GolangPlayground/languagelib"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

type Person struct {
	Name string
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	add := func(x, y int) int {
		return x + y
	}
	minus := func(x int, y int) (res int) {
		res = x - y
		return
	}

	var result = operator(add, 41, 1)
	fmt.Println(result)
	result = operator(minus, result, 1)
	fmt.Println(result)

	var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	anders := Person{"Anders"}
	fmt.Println(anders)
	nameChanger(&anders)
	fmt.Println(anders)

	languages := lib.GetAll()
	for _, v := range languages {
		fmt.Println(v)
	}

	countries := countrylist.GetAll()
	for _, v := range countries {
		fmt.Println(v)
	}
}

func operator(fn func(int, int) int, x int, y int) int {
	return fn(x, y)
}

func nameChanger(p *Person) {
	p.Name = "fisk"
}

func (p Person) String() string {
	return fmt.Sprintf("This is '%v'", p.Name)
}
