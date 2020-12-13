package main

import "fmt"

func main() {
	var number int
START:
	fmt.Println("input number : ")
	fmt.Scan(&number)
	if number >= 500 && number <= 9000 {
		for i := 3; i <= number; i = i + 2 {
			fmt.Println(i)
		}
	} else {
		fmt.Println("invalid number")
		goto START
	}
}

func b() {
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		}
	}
}

func c() {
	var f float64
	fmt.Scanf("%E", &f)
	fmt.Printf("%E", f)
}

func d() {
	var d int
	fmt.Scanf("%d", &d)
	fmt.Printf("%b", d)
}

func e() {
	var d int
	var f float64
	fmt.Scanf("%d %E", &d, &f)
	fmt.Printf("%b %E", d, f)
}

func f() {
	var d int
	var f float64
	var s string
	fmt.Scanf("%d %E %s", &d, &f, &s)
	fmt.Printf("%T %T %T", d, f, s)
}

func g() {
	x := []int{248, 428, 824}
	for _, v := range x {
		fmt.Printf("%d\n", v)
	}
}
