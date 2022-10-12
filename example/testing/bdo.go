package testing

import "context"

func (b *Builder) Do(ctx context.Context, root *Action) {
	// STEP1 OMIT
	all := actionList(root)
	for i, a := range all {
		a.priority = i
	}

	for _, a := range all {
		a.pending = len(a.Deps)
		if a.pending == 0 {
			b.ready.push(a)
			b.readySema <- true
		}
	}
	// STEP1END OMIT

	// STEP2 OMIT
	handle := func(ctx context.Context, a *Action) {
		var err error
		if a.Func != nil && (!a.Failed || a.IgnoreFail) {
			err = a.Func(b, ctx, a)
		}
	}
	// STEP2END OMIT

	// STEP3 OMIT
	par := cfg.BuildP
	if cfg.BuildN {
		par = 1
	}
	for i := 0; i < par; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case _, ok := <-b.readySema:
					if !ok {
						return
					}
					// Receiving a value from b.readySema entitles
					// us to take from the ready queue.
					b.exec.Lock()
					a := b.ready.pop()
					b.exec.Unlock()
					handle(ctx, a)
				}
			}
		}()
	}
	// STEP3END OMIT
}
