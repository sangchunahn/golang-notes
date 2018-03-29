// to execute the project we run "go run main.go" or the main file
// or run ./main after running the build code "go build main.go"

// package main package is a project or workspace
// purpose is to group together code with a similar purpose
// it must declare which package the file belongs to
// two packages in go executable file (main was created because it is executable) or reusable file (code dependancies, librarys or reuse code)
// how do we know? name of the package that determines executable
// main is executable. if we use something else its will not be executable
package main

// used to give our package access to our code
// ask fmt to access our code
// fmt is standard library of go
// golang.org/pkg lists all the packages in the library
import "fmt"

// we called our function main because we named our package main and package main requires a func main.
// func is short for function
func main() {
	fmt.Println("Hi there!")
}
