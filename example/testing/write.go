package testing

import "os"

func write() {
	// START OMIT
	if !cfg.BuildN {
		// writeTestmain writes _testmain.go,
		// using the test description gathered in t.
		if err := os.WriteFile(testDir+"_testmain.go", *pmain.Internal.TestmainGo, 0666); err != nil {
			return nil, nil, nil, err
		}
	}
	// END OMIT
}
