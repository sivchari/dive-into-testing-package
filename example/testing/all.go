package testing

func all() {
	// START OMIT
	runAction = &work.Action{
		Mode:       "test run",
		Func:       c.builderRunTest,
		Deps:       []*work.Action{buildAction},
		Package:    p,
		IgnoreFail: true, // run (prepare output) even if build failed
		TryCache:   c.tryCache,
		Objdir:     testDir,
	}
	cleanAction = &work.Action{
		Mode:       "test clean",
		Func:       builderCleanTest,
		Deps:       []*work.Action{runAction},
		Package:    p,
		IgnoreFail: true, // clean even if test failed
		Objdir:     testDir,
	}
	printAction = &work.Action{
		Mode:       "test print",
		Func:       builderPrintTest,
		Deps:       []*work.Action{cleanAction},
		Package:    p,
		IgnoreFail: true, // print even if test failed
	}
	// END OMIT
}
