package main

import "reflect"

//Reverse reverses a slice.
func Reverse(slice interface{}) {
	p := reflect.TypeOf(slice)
	if p.Kind() != reflect.Ptr {
		println("Expected pointed to slice. Found " + p.Kind().String())
		return
	}

	if p.Elem().Kind() != reflect.Slice {
		println("Can only reverse slices. Found type " + p.Elem().Kind().String())
		return
	}

	// create empty slice
	values := reflect.ValueOf(slice)
	tmp := reflect.MakeSlice(values.Elem().Type(), values.Elem().Len(), values.Elem().Cap())

	// write slice values to empty slice in reverse
	for i := 0; i < values.Elem().Len(); i++ {
		tmp.Index(values.Elem().Len()-1-i).Set(values.Elem().Index(i))
	}

	// write reverse values back to slice
	for i := 0; i < values.Elem().Len(); i++ {
		values.Elem().Index(i).Set(tmp.Index(i))
	}
}

func main() {
	println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it applly to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
}
