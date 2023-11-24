package usefulgo

import (
	"fmt"
	"os"
	"testing"
)

func Check(e error) {
	if e != nil {
		fmt.Printf("ERROR: %+v\n", e)
		os.Exit(1)
	}
}

func CheckTest(e error, t *testing.T) {
	if e != nil {
		t.Logf("ERROR: %+v\n", e)
		t.Fail()
	}
}

func CheckWithReason(e error, reason string) {
	if e != nil {
		fmt.Printf("%s: %+v\n", reason, e)
		os.Exit(1)
	}
}

func CheckTestWithReason(e error, t *testing.T, reason string) {
	if e != nil {
		t.Logf("%s: %+v\n", reason, e)
		t.Fail()
	}
}
