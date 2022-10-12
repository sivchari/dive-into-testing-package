package testing

func ptest() {
	// START OMIT
	ptest = new(Package)
	*ptest = *p
	ptest.Error = ptestErr
	ptest.ForTest = p.ImportPath
	ptest.GoFiles = nil
	ptest.GoFiles = append(ptest.GoFiles, p.GoFiles...)
	ptest.GoFiles = append(ptest.GoFiles, p.TestGoFiles...)
	ptest.Target = ""
	// END OMIT
}
