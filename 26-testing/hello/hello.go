package hello

import (
	"fmt"
	"testing_test/intego"
)

func Hello() string {
	return "Hello, testing ground"
}
func Hello2(n string) string {
	if n == "" {
		n = "People"
	}
	return fmt.Sprintf("Hello, testing ground...%s", n)
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello2("fa"))
	fmt.Println(intego.Add(2, 3))

}
