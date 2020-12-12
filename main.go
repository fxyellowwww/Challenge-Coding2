package main

import "fmt"

func main() {
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		}
	}
}

func a() {
	var f float64
	fmt.Scanf("%E", &f)
	fmt.Printf("%E", f)
}

func b() {
	var d int
	fmt.Scanf("%d", &d)
	fmt.Printf("%b", d)
}

func c() {
	var d int
	var f float64
	fmt.Scanf("%d %E", &d, &f)
	fmt.Printf("%b %E", d, f)
}

func d() {
	var d int
	var f float64
	var s string
	fmt.Scanf("%d %E %s", &d, &f, &s)
	fmt.Printf("%T %T %T", d, f, s)
}
