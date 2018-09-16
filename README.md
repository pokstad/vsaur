# VSaur = Vendorsaurus
VSaur is shorthand for Vendorsaurus. Vendorsaurus is collection of utilities for maintaining versioned Go modules (a.k.a. vgo) usage in a project. It is designed to be used in CI/CD systems to automate best practices.

## Features
Planned features:

- Lint mod file (`go.mod`)
	- Flag local path usage - should only be allowed during local development and not in commits to CI
- Lint sum file (`go.sum`)
	- Flag mismatch of `go.sum` listing with what the actual project requires to compile
- Verify vendor folder matches `go.mod` and `go.sum`

## Install
Vendorsaurus requires Go 1.11 or higher. To install, go get it:

`go get -u github.com/pokstad/vsaur/cmd/vsaur`

This will install the `vsaur` executable at `$GOPATH/src/go/bin/vsaur`. If your `$GOPATH` environment variable is not set, by default it will be `$HOME/go`.

Or you can download a release copy from the GitHub page.

## Usage
Simply run the `vsaur` command in the top level directory of a Go module you wish to validate:

```
$ $GOPATH/bin/vsaur
```

For a current and complete list of command options, view the command's Go doc string:

```
$ go doc github.com/pokstad/vsaur/cmd/vsaur
Command vsaur provides static analysis tools for automating best practices
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
```