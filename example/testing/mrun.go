package testing

import (
	"fmt"
	"os"
)

func mrun() {
	// START OMIT
	if !*isFuzzWorker {
		deadline := m.startAlarm()
		haveExamples = len(m.examples) > 0
		testRan, testOk := runTests(m.deps.MatchString, m.tests, deadline)
		fuzzTargetsRan, fuzzTargetsOk := runFuzzTests(m.deps, m.fuzzTargets, deadline)
		exampleRan, exampleOk := runExamples(m.deps.MatchString, m.examples)
		m.stopAlarm()
		if !testRan && !exampleRan && !fuzzTargetsRan && *matchBenchmarks == "" && *matchFuzz == "" {
			fmt.Fprintln(os.Stderr, "testing: warning: no tests to run")
		}
		if !testOk || !exampleOk || !fuzzTargetsOk || !runBenchmarks(m.deps.ImportPath(), m.deps.MatchString, m.benchmarks) || race.Errors() > 0 {
			fmt.Println("FAIL")
			m.exitCode = 1
			return
		}
	}
	// END OMIT
}
