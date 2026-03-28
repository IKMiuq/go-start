package main

import (
	"fmt"
	"os"
	"strings"
)

/* this is a comment */
// комментарий

func main() {
	lesson4()
}

func lesson1() {
	//var name = os.Args[1]
	var name1 string
	fmt.Scan(&name1)
	//fmt.Println("Hello, my name is " + name)
	fmt.Printf("Hello, my name is %s\n", name1)
}

func lesson2() {
	var arg1 int
	var arg2 int
	var operation string
	//var ErrInvalid = errors.New("Операция не найдена")
	if _, err := fmt.Scan(&arg1); err != nil {
		fmt.Fprintln(os.Stderr, "Не удалось прочитать первое число:", err)
		os.Exit(1)
	}
	fmt.Printf("Введите символ операции *, /, +, - : ")
	if _, err := fmt.Scan(&operation); err != nil {
		fmt.Fprintln(os.Stderr, "Не удалось прочитать операцию:", err)
		os.Exit(1)
	}
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		fmt.Fprintln(os.Stderr, "Операция не найдена")
		os.Exit(1)
	}
	if _, err := fmt.Scan(&arg2); err != nil {
		fmt.Fprintln(os.Stderr, "Не удалось прочитать второе число:", err)
		os.Exit(1)
	}
	switch operation {
	case "+":
		fmt.Printf("%d + %d = %d\n", arg1, arg2, arg1+arg2)
	case "-":
		fmt.Printf("%d - %d = %d\n", arg1, arg2, arg1-arg2)
	case "*":
		fmt.Printf("%d * %d = %d\n", arg1, arg2, arg1*arg2)
	case "/":
		if arg2 == 0 {
			fmt.Fprintln(os.Stderr, "На ноль делить нельзя")
			os.Exit(1)
		}
		fmt.Printf("%d / %d = %.2f\n", arg1, arg2, float64(arg1)/float64(arg2))
	}
}

func lesson3() {
	var far, funt int
	fmt.Scan(&far)
	fmt.Scan(&funt)
	fmt.Printf("%d фарингейт - это %.5f цельсия\n\r", far, (float64(far)-32)*5/9)
	fmt.Printf("%d фут - это %.5f метров\n\r", funt, float64(funt)*0.3048)
}

func lesson4() {
	var code string
	fmt.Scan(&code)
	code = strings.ToUpper(code[:1]) + code[1:]
	fmt.Print(code)
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println()
	fmt.Print(x[2:5])
	y := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := y[0]
	//size := count(y)
	//var size int
	for size := len(y) - 1; size >= 0; size-- {
		if y[size] < min {
			min = y[size]
		}
	}
	fmt.Println()
	fmt.Printf("Минимальное заначение в массиве: %d\n", min)
}
