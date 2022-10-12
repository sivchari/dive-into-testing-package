package testing

import "path/filepath"

// START OMIT
func loadTestFuncs(ptest *Package) (*testFuncs, error) {
	t := &testFuncs{
		Package: ptest,
	}
	var err error
	for _, file := range ptest.TestGoFiles {
		if lerr := t.load(filepath.Join(ptest.Dir, file), "_test", &t.ImportTest, &t.NeedTest); lerr != nil && err == nil {
			err = lerr
		}
	}
	for _, file := range ptest.XTestGoFiles {
		if lerr := t.load(filepath.Join(ptest.Dir, file), "_xtest", &t.ImportXtest, &t.NeedXtest); lerr != nil && err == nil {
			err = lerr
		}
	}
	return t, err
}

// END OMIT
