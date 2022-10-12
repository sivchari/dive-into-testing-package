package testing

func (t *T) report() {
	if t.parent == nil {
		return
	}
	dstr := fmtDuration(t.duration)
	// START OMIT
	format := "--- %s: %s (%s)\n"
	if t.Failed() {
		t.flushToParent(t.name, format, "FAIL", t.name, dstr)
	} else if t.chatty != nil {
		if t.Skipped() {
			t.flushToParent(t.name, format, "SKIP", t.name, dstr)
		} else {
			t.flushToParent(t.name, format, "PASS", t.name, dstr)
		}
	}
	// END OMIT
}
