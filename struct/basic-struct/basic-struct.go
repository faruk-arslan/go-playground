package main

import "fmt"

type spec struct {
	mechanism string
}
type watch struct {
	brand     string
	specSheet spec
}

func main() {
	// Create a watch instance
	sheet := spec{"quartz"}
	myWatch := watch{"swatch", sheet}
	// myWatch instance also can be defined like below as single line:
	// myWatch := watch{"swatch", spec{"quartz"}}
	fmt.Println("watch before -->", myWatch)
	fmt.Println("sheet before -->", sheet)
	// Modify the fields on watch instance
	myWatch.brand = "longines"
	myWatch.specSheet.mechanism = "auto"
	/**
	The watch instance has been modified but the sheet instance stays the same
	because in the watch struct we did not use the memory address of the spec struct
	used the value of it instead.
	*/
	fmt.Println("watch after -->", myWatch)
	fmt.Println("sheet after -->", sheet) // --> quartz
}
