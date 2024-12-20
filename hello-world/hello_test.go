package main 

import "testing"


func TestHello(t *testing.T){
	got := Hello("osama")

	want := "Hello, osama"

	if got != want {
		t.Errorf("got %q want %q ", got,want)
	}
}