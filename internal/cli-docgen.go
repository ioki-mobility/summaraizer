package main

import (
	"fmt"
	"os"

	"github.com/ioki-mobility/summaraizer/internal/cli"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := cli.NewRootCmd()
	cmd.DisableAutoGenTag = true
	if err := doc.GenMarkdownTree(cmd, "./internal/cli/docs"); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
