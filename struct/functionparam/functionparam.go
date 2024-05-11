package main

import "fmt"

type spec struct {
	mechanism string
}
type watchStruct struct {
	brand     string
	specSheet spec
}

func main() {
	sheet := spec{"quartz"}
	watch := watchStruct{"casio", sheet}
	fmt.Println("initial state --> ", watch)

	// ----- Change the brand of the watch instance -----
	// Sending as value
	changeBrand("new brand", watch)
	fmt.Println("after changeBrand --> ", watch) // brand has not changed on watch instance
	// Sending as reference
	changeBrandWithPointer("pointer brand", &watch)
	fmt.Println("after changeBrandWithPointer --> ", watch) // brand has been changed on watch instance
	fmt.Println("----------")

	// ----- Change the specSheet.mechanism of the watch instance -----
	// Sending as value (inner struct change)
	fmt.Println("initial mechanism --> ", watch.specSheet.mechanism)
	changeMechanism("changed mechanism", watch)
	fmt.Println("after changeMechanism --> ", watch.specSheet.mechanism) // mechanism has not changed
	// Sending as reference (inner struct change)
	changeMechanismWithPointer("changed inner mechanism", &watch)
	fmt.Println("after changeMechanismWithPointer -->", watch.specSheet.mechanism) // mechanism has been changed on watch instance
	// The mechanism of the watch instance has been changed, but the sheet instance itself stays the same
	// because spec struct is not a pointer in watchStruct.
	fmt.Println("sheet: ", sheet)
}

func changeBrand(brand string, watch watchStruct) {
	watch.brand = brand
	fmt.Println("inside changeBrand: ", watch)
}

func changeBrandWithPointer(brand string, watch *watchStruct) {
	watch.brand = brand // or (*watch).brand = brand
	fmt.Println("inside changeBrandWithPointer: ", watch)
}

func changeMechanism(mechanism string, watch watchStruct) {
	watch.specSheet.mechanism = mechanism
	fmt.Println("inside changeMechanism: ", watch.specSheet.mechanism)
}

func changeMechanismWithPointer(mechanism string, watch *watchStruct) {
	watch.specSheet.mechanism = mechanism
	fmt.Println("inside changeMechanismWithPointer: ", watch.specSheet.mechanism)
}
