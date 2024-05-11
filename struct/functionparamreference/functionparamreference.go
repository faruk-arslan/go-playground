package main

import "fmt"

type spec struct {
	mechanism string
}
type watchStructWithPointer struct {
	brand     string
	specSheet *spec
}

func main() {
	sheet := spec{mechanism: "auto"}
	watch := watchStructWithPointer{brand: "new brand", specSheet: &sheet}

	// With below changes, the sheet is changed.
	// We did not send the reference of the watch struct itself but in specSheet we hold the reference of the specStruct.
	// So the changes that have been made in the function is reflected both on the sheet instance itself and the specSheet of the watch struct.
	fmt.Println("sheet before --> ", sheet)
	fmt.Println("watch sheet before --> ", watch.specSheet.mechanism)
	changeMechanism("changed mechanism", watch)
	fmt.Println("sheet after -->", sheet)
	fmt.Println("watch sheet after --> ", watch.specSheet.mechanism)
	fmt.Println("----------")
	// This has the same behaviour with the above example
	fmt.Println("sheet before -->", sheet)
	fmt.Println("watch sheet before --> ", watch.specSheet.mechanism)
	changeMechanismWithPointer("pointer mechanism", &watch)
	fmt.Println("sheet after -->", sheet)
	fmt.Println("watch sheet after --> ", watch.specSheet.mechanism)

}

func changeMechanism(mechanism string, watch watchStructWithPointer) {
	watch.specSheet.mechanism = mechanism
	fmt.Println("inside changeMechanism: ", watch.specSheet.mechanism)
}

func changeMechanismWithPointer(mechanism string, watch *watchStructWithPointer) {
	watch.specSheet.mechanism = mechanism
	fmt.Println("inside changeMechanismWithPointer: ", watch.specSheet.mechanism)
}
