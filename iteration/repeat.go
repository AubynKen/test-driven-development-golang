package iteration

func Repeat(nbIter int) (repeated string) {
	for i := 0; i < nbIter; i++ {
		repeated += "a"
	}
	return
}
