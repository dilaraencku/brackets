package test

import (
	"TGACase"
	"testing"
)

func TestBracketsSymmetryControlReturnYes(t *testing.T) {
	response := TGACase.IsBalanced("{([])}()")

	if response == "YES"{
		t.Error("Everything is okay")
	}
}

func TestBracketsSymmetryControlReturnNo(t *testing.T) {
	response := TGACase.IsBalanced("{([])}()")

	if response == "NO" {
		t.Error("Everything is okay")
	}
}


func TestBracketsSymmetryControlInvalidParam(t *testing.T) {
	response := TGACase.IsBalanced("(y)[]")

	if response != "NO" {
		t.Error("Everything is okay")
	}
}
