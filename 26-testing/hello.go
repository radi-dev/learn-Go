package testing_test

import "fmt"

func Hello() string {
	return "Hello, testing ground"
}
func Hello2(n string) string {
	return fmt.Sprintf("Hello, testing ground...%s", n)
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello2("fa"))

}
