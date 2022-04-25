package stringutilities_test

import (
	"charlitosf/tfm-server/pkg/stringutilities"
	"testing"
)

// Test the reverse split join function
func TestReverseSplitJoin(t *testing.T) {
	// Test 1
	s := "1.2.3"
	if stringutilities.ReverseSplitJoin(s) != "3.2.1" {
		t.Errorf("ReverseSplitJoin(%s) != %s", s, "3.2.1")
	}
	// Test 2
	s = "1.2"
	if stringutilities.ReverseSplitJoin(s) != "2.1" {
		t.Errorf("ReverseSplitJoin(%s) != %s", s, "2.1")
	}
	// Test 3
	s = "test"
	if stringutilities.ReverseSplitJoin(s) != "test" {
		t.Errorf("ReverseSplitJoin(%s) != %s", s, "test")
	}
}
