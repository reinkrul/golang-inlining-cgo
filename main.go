package main

import "C"

func main() {
	InlinedOrProbablySkippedExample()
	InlinedExample()
	NotInlinedExample()
}

func InlinedOrProbablySkippedExample() {
	// Empty function
}

func InlinedExample() {
	println("MyFunc")
}

func NotInlinedExample() {
	println("MyFunc")
	C.CString("test")
}