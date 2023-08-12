package utils

// GetOrPanic returns the first parameter if err is nil.
// Otherwise, it will panic
func GetOrPanic[T any](input T, err error) T {
	if err != nil {
		panic(err)
	}
	return input
}

func StringifyErr(input error) string {
	return input.Error()
}
