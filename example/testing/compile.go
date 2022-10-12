package testing

func complile() {
	// START OMIT
	// Set compile objdir to testDir we've already created,
	// so that the default file path stripping applies to _testmain.go.
	b.CompileAction(work.ModeBuild, work.ModeBuild, pmain).Objdir = testDir

	a := b.LinkAction(work.ModeBuild, work.ModeBuild, pmain)
	// END OMIT
}
