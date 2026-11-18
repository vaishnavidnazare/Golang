package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("hello world!")
	var sample string = "Hii"
	fmt.Printf("the string is %s", sample)
	fmt.Println(" ")
	var (
		name string
		age  int
	)
	fmt.Println("Enter your name and age: ")
	count, err := fmt.Scanf("%s %d", &name, &age)
	fmt.Println("Hii there ", name)
	fmt.Println("Okay so your age is ", age)
	fmt.Println("count: ", count)
	fmt.Println("err: ", err)
	fmt.Printf("the variable age = %v is of type %T \n", age, age)
	fmt.Printf("the variable age is of type %T \n", age)
	fmt.Printf("the variable name = %v is of type %v \n", name, reflect.TypeOf(name))

	var str string = "200"
	i, err := strconv.Atoi(str)
	fmt.Printf("%v, %v \n", i, err)

	var integer int = 2
	fmt.Printf("%v, %T \n", integer, strconv.Itoa(integer))
}

// package: used for organising and reusing code.
// import : to include code from other packages to use in our program.
// func main(): Building block of go program,
// fmt.Println: fmt is a package and Println is a function from it.The dot is a member access operator, used to access members of the package.

/*
muli line comment
*/
