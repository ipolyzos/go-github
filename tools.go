// +build tools

package github

import (
	// used by continuous integration scripts
	_ "github.com/google/go-github/v24"
	_ "github.com/mattn/goveralls"
	_ "golang.org/x/tools/cmd/cover"
)
