package main

import (
	"fmt"
	"strings"
)

// func runTest(cmd *Command, args []string) {
func runTest() {
	// START OMIT
	// Prepare build + run + print actions for all packages being tested.
	for _, p := range pkgs {
		buildTest, runTest, printTest, err := builderTest(&b, ctx, pkgOpts, p, allImports[p])
		if err != nil {
			str := err.Error()
			str = strings.TrimPrefix(str, "\n")
			if p.ImportPath != "" {
				base.Errorf("# %s\n%s", p.ImportPath, str)
			} else {
				base.Errorf("%s", str)
			}
			fmt.Printf("FAIL\t%s [setup failed]\n", p.ImportPath)
			continue
		}
	}
	// Ultimately the goal is to print the output.
	root := &work.Action{Mode: "go test", Func: printExitStatus, Deps: prints}
	b.Do(ctx, root)
	// END OMIT
}
