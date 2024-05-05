package main

import "fmt"

type spec struct {
	mechanism string
}
type watchWithPointer struct {
	brand     string
	specSheet *spec
}

func main() {
	sheet := spec{"quartz"}
	// The memory address of the sheet object is given when creating a myWatch instance via "&" symbol
	myWatch := watchWithPointer{"swatch", &sheet}
	fmt.Println("myWatch before -->", myWatch.brand, myWatch.specSheet.mechanism)
	fmt.Println("sheet before -->", sheet)
	// Change the fields
	myWatch.brand = "tissot"
	/**
	Below change has been reflected to the sheet instance itself, since we used its reference
	in the specSheet field of the watchWithPointer struct and changed it afterward.
	*/
	myWatch.specSheet.mechanism = "auto"
	fmt.Println("myWatch after -->", myWatch.brand, myWatch.specSheet.mechanism)
	fmt.Println("sheet after -->", sheet) // --> auto
}
