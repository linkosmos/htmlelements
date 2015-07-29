package htmlelements

import (
	"testing"

	"golang.org/x/net/html/atom"
)

var existTests = []struct {
	atomSlice     AtomSlice
	testedElement atom.Atom
	expected      bool
}{
	{AtomSlice{atom.H1, atom.A}, atom.H1, true},
	{AtomSlice{atom.P}, atom.Base, false},
	{AtomSlice{atom.Div, atom.Aside}, atom.P, false},
}

func TestExist(t *testing.T) {
	for _, test := range existTests {
		got := test.atomSlice.Exist(test.testedElement)
		if test.atomSlice.Exist(test.testedElement) != test.expected {
			t.Errorf("Expected %t got %t", test.expected, got)
		}
	}

}
