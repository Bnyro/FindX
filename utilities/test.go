package utilities

// A single test suite.
type TestSuite[A, W any] struct {
	Args A
	Want W
}

func Test[A, W comparable](tests []TestSuite[A, W], toTest func(A) W, log func(W, W)) {
	for _, test := range tests {
		if got := toTest(test.Args); got != test.Want {
			log(got, test.Want)
		}
	}
}
