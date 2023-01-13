package main

import "testing"

func TestHello2(t *testing.T) {
	res := GetHello2()
	if res != "hello world 3" {
		t.Fail()
	}
}
