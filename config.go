package main

import (
	"io"
	"os"

	"github.com/rickb777/typewriter"
)

type config struct {
	out        io.Writer
	customName string
	*typewriter.Config
}

var defaultConfig = config{
	out:        os.Stdout,
	customName: "_gen.go",
	Config:     &typewriter.Config{},
}

// keep in sync with imports.go
var stdImports = typewriter.NewImportSpecSet(
	typewriter.ImportSpec{Name: "_", Path: "github.com/rickb777/slice"},
	typewriter.ImportSpec{Name: "_", Path: "github.com/rickb777/stringer"},
)
