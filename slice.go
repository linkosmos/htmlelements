package htmlelements

import "golang.org/x/net/html/atom"

// AtomSlice - convenience type to host Atom's & iterate through
type AtomSlice []atom.Atom

// Exist - check if given Atom element exist
func (atomslice *AtomSlice) Exist(at atom.Atom) bool {
	for _, a := range *atomslice {
		if a == at {
			return true
		}
	}
	return false
}
