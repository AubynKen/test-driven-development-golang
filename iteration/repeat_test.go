package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat(10)
	expect := "aaaaaaaaaa"
	if got != expect {
		t.Errorf("got %q expect %q", got, expect)
	}
}

func ExampleRepeat() {
	repeatedString := Repeat(10)
	fmt.Println(repeatedString)
	// Output: aaaaaaaaaa
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat(10)
	}
}
