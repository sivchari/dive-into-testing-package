package testing

import "time"

func tRunner() {
	// START OMIT
	tRunner(t, func(t *T) {
		for _, test := range tests {
			t.Run(test.Name, test.F)
		}
	})
	// END OMIT
}

func tRunner2() {
	// START2 OMIT
	if len(t.sub) > 0 {
		// Run parallel subtests.
		// Decrease the running count for this test.
		t.context.release()
		// Release the parallel subtests.
		close(t.barrier)
		// Wait for subtests to complete.
		for _, sub := range t.sub {
			<-sub.signal
		}
		cleanupStart := time.Now()
		err := t.runCleanup(recoverAndReturnPanic)
		t.duration += time.Since(cleanupStart)
		if err != nil {
			doPanic(err)
		}
		if !t.isParallel {
			// Reacquire the count for sequential tests. See comment in Run.
			t.context.waitParallel()
		}
	} else if t.isParallel {
		// Only release the count for this test if it was run as a parallel test. See comment in Run method.
		t.context.release()
	}
	// END2 OMIT
}
