package testing

// START OMIT
type common struct {
	output   []byte // Output generated by test or benchmark.
	ran      bool   // Test or benchmark (or one of its subtests) was executed.
	failed   bool   // Test or benchmark has failed.
	skipped  bool   // Test or benchmark has been skipped.
	done     bool   // Test is finished and all subtests have completed.
	finished bool   // Test function has completed.
	inFuzzFn bool   // Whether the fuzz target, if this is one, is running.
	bench    bool   // Whether the current test is a benchmark.
	runner   string // Function name of tRunner running the test.
	name     string // Name of test or benchmark.
}

// END OMIT