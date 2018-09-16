package vsaur

import (
	"context"
	"fmt"
)

// Checker is a static analysis tool that given a module top level directory
// will report back a list of locations where issues have been found
type Checker interface {
	CheckName() string
	Check(context.Context) ([]Issue, error)
}

// Issue is a problem found in the project by a checker
type Issue struct {
	Description string
	Path        string // file system path to issue
	Line        uint   // line number within file (starting at 0)
	Column      uint   // column within line ( starting at 0)
}

// CheckerFactory constructs a checker from a module's top level directory path
type CheckerFactory func(moddir string) Checker

// allCheckers is a collection of all known checker factories keyed by the
// checker name
var allCheckers = map[string]CheckerFactory{}

// AllCheckers returns a list of all known checker names
func AllCheckers() []string {
	var checkers []string
	for k, _ := range allCheckers {
		checkers = append(checkers, k)
	}
	return checkers
}

// MustAddChecker will add a checker to the global list. If there is a name
// conflict, a panic will be raised.
func MustAddChecker(name string, c CheckerFactory) {
	_, ok := allCheckers[name]
	if ok {
		panic(fmt.Errorf("a  checker is already registered as %s", name))
	}

	allCheckers[name] = c
}
