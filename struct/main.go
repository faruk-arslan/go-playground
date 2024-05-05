package main

type spec struct {
	mechanism string
}
type watch struct {
	brand     string
	specSheet spec
}
type watchStructWithPointer struct {
	brand     string
	specSheet *spec
}
