package test

import (
	"TGACase"
	"testing"
)

func TestBracketsSymmetryControlReturnYes(t *testing.T) {
	response := TGACase.IsBalanced("{([])}()")

	if response != "YES"{
		t.Error("There is a error, this is not excepted")
	}
}

func TestBracketsSymmetryControlReturnNo(t *testing.T) {
	response := TGACase.IsBalanced("{([])}()")

	if response != "YES"{
		t.Error("There is a error, this is not excepted")
	}
}



func TestBracketsSymmetryControlInvalidParam(t *testing.T) {
	response := TGACase.IsBalanced("(y)[]")

	if response != "N"{
		t.Error("There is a error, this is not excepted")
	}
}
