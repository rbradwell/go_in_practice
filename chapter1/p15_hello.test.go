package main

// TestName : testing p14_hello.go
// this is meant to test p14_hello.go getName method by typing 'go test; at command line but doesn't work
// struct is passed in to help with testing
/*
func TestName(t *testing.T) {
	name := getName()
	if name != "World!" {
		t.Error("Response from getname is unexpected value")
	}
}
*/
// as of GO 1.5 code coverage is also a core tool
// go tst -cover
// see http://blog.golang.org/cover
