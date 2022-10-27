package main

import "fmt"

func calc(a int) int {
	a = a + 1
	return a
}

func pcalc(a *int) int {
	*a = *a + 1
	return *a
}

var sum = func(a, b int) int {
	fmt.Printf("func sum call!\n")
	return a + b
}

func doinput(f func(int, int) int, a, b int) int {
	fmt.Printf("func doinput call!\n")
	return f(a, b)
}

func wrap(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a, b int) int {
			fmt.Printf("wrap.add call!\n")
			return a + b
		}
	case "sub":
		return func(a int, b int) int {
			fmt.Printf("wrap.sub call!\n")
			return a - b
		}
	default:
		return nil
	}

}

func fa(base int) (func(int) int, func(int) int) {
	println(&base, base)
	return func(i int) int {
			base += i
			println(&base, base)
			return base
		}, func(i int) int {
			base -= i
			println(&base, base)
			return base
		}
}
func main() {
	//var a float32 = 4 // 这里会自动转换
	////var b int = a // 强类型语言，这里语法错误
	//var b int = 4
	//fmt.Printf("%f %d Hello! world!", a, b)

	////test switch
	//switch i := "y"; i {
	//case "y", "Y":
	//	fmt.Printf("yes")
	//	fallthrough
	//case "n":
	//	fmt.Printf("no")
	//}

	////test param
	//var a int = 5
	//b := calc(a)
	//fmt.Printf("a, b=%d, %d\n", a, b)
	//c := pcalc(&a)
	//fmt.Printf("a, c=%d, %d\n", a, c)

	////test anonymous func
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Printf("no err\n")
	//	}
	//}()
	//
	//sum(1, 2)
	//doinput(func(a int, b int) int {
	//	fmt.Printf("doinput.f call!\n")
	//	return a * b
	//}, 2, 3)
	//
	//opFunc := wrap("add")
	//opFunc(4, 5)

	//test closure
	f, g := fa(0)
	s, k := fa(0)
	println(f(1), g(2))
	println(s(1), k(2))

}
