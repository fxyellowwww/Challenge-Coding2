package main

import "fmt"

func main() {
	y := 2482
	switch y {
	case 1946:
		fmt.Println("1946")
	case 9463:
		fmt.Println("9463")
	case 2482:
		fmt.Println("true")
	}
}

//รับค่า 500-9000 แล้วแสดงค่า 1-input โดยแสดงเฉพาะเลขคี่
func a() {
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

//แสดงเลข 1-1000 เลขที่หาร 3 ลงตัว ให้แสดงคำว่า fizz หาร 5 ลงตัว ให้แสดงคำว่า buzz
func b() {
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		}
	}
}

//รับค่าทศนิยม แล้วแสดงผลออกมาเป็นสัญกรณ์วิทยาศาสตร์
func c() {
	var f float64
	fmt.Scanf("%E", &f)
	fmt.Printf("%E", f)
}

//รับค่าจำนวนเต็ม แล้วแสดงผลออกมาเป็นเลขฐานสอง
func d() {
	var d int
	fmt.Scanf("%d", &d)
	fmt.Printf("%b", d)
}

//รับค่าจำนวนเต็มและทศนิยม แล้วแสดงผล โดยให้จำนวนเต็มแสดงเป็นเลขฐานสอง และทศนิยมแสดงเป็นสัญกรณ์วิทยาศาสตร์
func e() {
	var d int
	var f float64
	fmt.Scanf("%d %E", &d, &f)
	fmt.Printf("%b %E", d, f)
}

//รับค่า 3 ค่า (จำนวนเต็ม ทศนิยม ตัวอักษร) แล้วแสดงผลเป็นชนิดข้อมูลของค่าที่รับมา
func f() {
	var d int
	var f float64
	var s string
	fmt.Scanf("%d %E %s", &d, &f, &s)
	fmt.Printf("%T %T %T", d, f, s)
}

//คำสั่ง range
func g() {
	x := []int{248, 428, 824}
	for _, v := range x {
		fmt.Printf("%d\n", v)
	}
}

//คำสั่ง goto
func h() {
	var a int
ReadInput:
	fmt.Print("type number :")
	fmt.Scan(&a)
	if a < 320 {
		goto ReadInput
	}
	fmt.Println(a)
}

//คำสั่ง switch
func i() {
	y := 2482
	switch y {
	case 1946:
		fmt.Println("1946")
	case 9463:
		fmt.Println("9463")
	case 2482:
		fmt.Println("true")
	}
}
