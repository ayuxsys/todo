package must

func Eval[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
