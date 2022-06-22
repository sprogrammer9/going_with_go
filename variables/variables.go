package main

import ( //multiple imports
	"fmt"
	"reflect"
	"strconv"
)

var (
	name, course   = "Oyeyebi Ridwan", "Getting Started with Go" //double declaration
	module, clip   = 4, 2                                        // same here
	time           = "6"                                         // used for type conversion
	courseComplete = false                                       // you dont have to use a variable declared at package level

	ptr *string = &time // ptr makes string a pointer variable holding the memory location of time
)

func main() {

	// courseComplete := false... you must use a variable declared at function level
	fmt.Printf("Name is '%s' and Course is '%s' and name is of type", name, course)
	fmt.Printf("Module is %v and Clip is %v", module, clip)
	fmt.Println("Name is of type", reflect.TypeOf(name), "while course is of type", reflect.TypeOf(course))
	fmt.Println("Module is of type", reflect.TypeOf(module), "while clip is a ", reflect.TypeOf(clip))

	iModule, err := strconv.Atoi(time) // strconv converts a string to int using Atoi, Go passes an argument by value and not reference, it makes a copy of the variable and sends the value to the function so any changes the function makes affects the copy made and not the original data in the variable
	if err == nil {                    // this says if the error code is empty
		total := iModule + clip
		fmt.Println("The sum of the above is: ", total)
		fmt.Println("The memory address of 'total' is ", &total)                // & is used to print the memory address of a variable
		fmt.Println("Pointing to the address ", ptr, "and the value of ", *ptr) // dereferencing a pointer variable by adding * to it
		updateCourse(&course)                                                   // using pointer to change the content of a variable
		fmt.Print(course)                                                       // prints the updated content of the course
	}
}

func updateCourse(course *string) string {
	*course = "getting started with docker"
	fmt.Println(*course)
	return *course
}
