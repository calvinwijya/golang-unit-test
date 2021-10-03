package helper

import "fmt"

func HelloWorld(name string) string {
	return "hello " + name
}

func Buzz(total int) {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
