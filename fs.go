package main

import "github.com/spf13/afero"

// FS is the filesystem.
var FS = afero.NewOsFs()
