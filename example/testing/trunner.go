package testing

func tRunner() {
	// START OMIT
	tRunner(t, func(t *T) {
		for _, test := range tests {
			t.Run(test.Name, test.F)
		}
	})
	// END OMIT
}
