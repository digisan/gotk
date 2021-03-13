package iter

import "testing"

func TestIter(t *testing.T) {

	for i := range N(10) {
		fPln(i)
	}
	fPln(" ------------ ")

	for i := range Iter(-10) {
		fPln(i)
	}
	fPln(" ------------ ")

	for i := range Iter(1, 10) {
		fPln(i)
	}
	fPln(" ------------ ")

	for i := range Iter(2, 2, 10) {
		fPln(i)
	}
	fPln(" ------------ ")

	func(slc ...int) {
		for _, a := range slc {
			fPln(a)
		}
	}(Iter2Slc(11, -4, 2)...)

	// ----------------------- //
	fPln(" ************************************* ")

	for i := range Iter(10, -2, -2) {
		fPln(i)
	}
	fPln(" ------------ ")

	for i := range Iter(10, 3) {
		fPln(i)
	}
}
