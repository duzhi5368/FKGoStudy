package FKLib

import (
	"reflect"
	"testing"
)

func Test_TrimString(t *testing.T) {
	in1 := "Hello"
	in2 := "He"
	out := "llo"

	r := TrimString(in1, in2)

	if r != out {
		t.Errorf("Test_TrimString Failed: Hope = %s, Dest = %s", out, r)
	}
}

func Test_SplitString(t *testing.T) {
	in1 := "Hello World"
	in2 := "l"
	out := []string{"He", "", "o Wor", "d"}

	r := SplitString(in1, in2)

	if !reflect.DeepEqual(r, out) {
		t.Errorf("Test_TrimString Failed: Hope = %s, Dest = %s", out, r)
	}
}
