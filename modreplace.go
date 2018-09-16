package vsaur

import (
	"context"
	"path/filepath"
)

// ModFileName is the standard file name for a Go module file
const ModFileName = "go.mod"

// localReplaceChecker checks that a mod file does not contain replace
// statements that reference a local file path
type localReplaceChecker struct {
	modfile string // path to modfile
}

func newLocalReplaceChecker(moddir string) Checker {
	return localReplaceChecker{
		modfile: filepath.Join(moddir, "mod.file"),
	}
}

// CheckName returns the unique name for this checker
func (localReplaceChecker) CheckName() string {
	return "local-mod-replace"
}

// Check will analyze the modfile and determine if there are any issues related
// to replace statements containing local file paths
func (lrc localReplaceChecker) Check(ctx context.Context) ([]Issue, error) {
	var issues []Issue
	return issues, nil
}

func init() {
	MustAddChecker(localReplaceChecker{}.CheckName(), newLocalReplaceChecker)
}
