package vsaur

import (
	"context"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/pokstad/vsaur/internal/modfile"
)

// ModFileName is the standard file name for a Go module file
const modFileName = "go.mod"

// localReplaceChecker checks that a mod file does not contain replace
// statements that reference a local file path
type localReplaceChecker struct {
	modfile string // path to modfile
}

func newLocalReplaceChecker(moddir string) Checker {
	return localReplaceChecker{
		modfile: filepath.Join(moddir, modFileName),
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

	data, err := ioutil.ReadFile(lrc.modfile)
	if err != nil {
		return nil, err
	}

	mf, err := modfile.Parse(lrc.modfile, data, nil)
	if err != nil {
		return nil, err
	}

	// TODO: replace following with actual check
	log.Printf("mf: %#v", mf)

	return issues, nil
}

func init() {
	MustAddChecker(localReplaceChecker{}.CheckName(), newLocalReplaceChecker)
}
