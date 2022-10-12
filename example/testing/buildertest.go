package testing

func builderTest() {
	// Build Package structs describing:
	//	pmain - pkg.test binary
	//	ptest - package + test files
	//	pxtest - package of external test files
	pmain, ptest, pxtest, err := load.TestPackagesFor(ctx, pkgOpts, p, cover)
	if err != nil {
		return nil, nil, nil, err
	}
}
