/*Command vsaur provides static analysis tools for automating best practices
related to versioned Go modules.

Usage:

To analyze everything in a project, simply run vsaur while in a project
directory. This will run all possible checks.

	$ vsaur
	TODO: sample output of vsaur running all checks

To only run a subset of checks, provides the -checks flag with a comma
delimited list of check names:

	$ vsaur -checks mod-local-replaces,sum-overflow
	TODO: sample output of vsaur running selective checks

By default, vsaur will look for a file in the current directory named
".vsaur.json" to load a JSON formatted configuration instead of relying on
flags. Or specify a config file path of your choosing with the flag -config:

	$ vsaur -config .vsaur-release.json

*/
package main

import (
	"flag"
	"strings"

	"github.com/pokstad/vsaur"
)

const (
	defaultCfgPath = ".vsaur.json"
)

var (
	checkFlags = flag.String("checks", "", "comma delimited list of checks")
	configPath = flag.String("config", defaultCfgPath, "path to config file")
)

func getChecks() []vsaur.Check {
	checks := strings.Split
	return nil
}

func main() {
	flag.Parse()

	checks := getChecks()
}
