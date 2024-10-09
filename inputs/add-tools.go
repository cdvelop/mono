package inputs

func permutation(A, B []string, extras ...[]string) (out []string) {

	for i, a := range A {
		for _, b := range B {
			combination := a + " " + b
			for _, extra := range extras {
				combination += " " + extra[i]
			}
			out = append(out, combination)
		}
	}

	return
}
