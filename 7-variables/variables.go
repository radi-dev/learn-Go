package main

import "fmt"

var declared_var_1 int
var declared_var_2, declared_var_3 int
var initialized_var_w_type int = 3
var initialized_var_no_type = 3
var (
	bracketed_var_1 = 5
	bracketed_var_2 = 6
)

//---short forms only work inside functions
// bad_short_form_initialized_var := 2
// bad_short_form_initialized_vars_1, bad_short_form_initialized_vars_2 := 2

func main() {
	declared_var_1 = 2
	declared_var_2 = 4

	short_form_initialized_var := 2
	short_form_initialized_vars_1, short_form_initialized_vars_2 := 2, 8

	fmt.Println(
		short_form_initialized_var, '\n',
		short_form_initialized_vars_1, '\n',
		short_form_initialized_vars_2, '\n', bracketed_var_2)

}
