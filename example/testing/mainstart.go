package testing

func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, fuzzTargets []InternalFuzzTarget, examples []InternalExample) *M {
	Init()
	return &M{
		deps:        deps,
		tests:       tests,
		benchmarks:  benchmarks,
		fuzzTargets: fuzzTargets,
		examples:    examples,
	}
}
