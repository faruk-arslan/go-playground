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
	// Creating with the :=
	sheet := spec{"quartz"}
	myWatch := watch{"swatch", sheet}
	fmt.Println(myWatch)

	/**
	Creating with the new keyword.
	Since we didn't give any value when creating the struct, default zero values are given initially.
	*/
	myWatch2 := new(watch)
	fmt.Println(myWatch2)  // Notice the "&" symbol in the log, it means we're holding a reference.
	fmt.Println(*myWatch2) // Now we see tha value without "&", because it's dereferenced with "*".
	myWatch2.brand = "casio"
	myWatch2.specSheet.mechanism = "quartz"
	fmt.Println(myWatch2)
}
