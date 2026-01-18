package main

import "fmt"

var FirstName = "T V S S"

func main() {
	// var age int = 19
	// var name string = "Harsha"
	// var name1 string = "harsha16x"
	// middleName := "T V S S"
	count := 10
	username := "harsha16x"
	fmt.Println(FirstName)
	fmt.Println(count)
	fmt.Println(username)
	
	// Gofer := is a go special declaration
	// var or int is not req while using gofer
	// DEFAULT VALUES
	// numeric types :0
	// boolean types : false
	// string type : ""
	// pointers,slices,maps,funcs,structs : nil

}

func printName() {
	RealName := "Harsha"
	fmt.Println(RealName)
	// Only can be used in this function
	// better memory efficiency
}
