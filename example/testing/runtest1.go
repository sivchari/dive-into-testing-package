package testing

import "os"

func runTests() {
	// START OMIT
	ctx := newTestContext(*parallel, newMatcher(matchString, *match, "-test.run"))
	ctx.deadline = deadline
	t := &T{
		common: common{
			signal:  make(chan bool, 1),
			barrier: make(chan bool),
			w:       os.Stdout,
		},
		context: ctx,
	}
	// END OMIT
}
