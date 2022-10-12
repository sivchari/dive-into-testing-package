package testing

func TestPackagesFor() {
	// START OMIT
	pmain, ptest, pxtest = TestPackagesAndErrors(ctx, opts, p, cover)
	for _, p1 := range []*Package{ptest, pxtest, pmain} {
		if p1 == nil {
			// pxtest may be nil
			continue
		}
		// ErrorかDepsErrorがあるとbreak
	}
	if pmain.Error != nil || len(pmain.DepsErrors) > 0 {
		pmain = nil
	}
	if ptest.Error != nil || len(ptest.DepsErrors) > 0 {
		ptest = nil
	}
	if pxtest != nil && (pxtest.Error != nil || len(pxtest.DepsErrors) > 0) {
		pxtest = nil
	}
	return pmain, ptest, pxtest, err
	// END OMIT
}
