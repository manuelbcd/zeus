package types

// IsAnyEmpty returns true if any of args has a zero value. It evaluates only
// string, int types.
func IsAnyEmpty(args ...interface{}) bool {

	for _, arg := range args {
		switch arg.(type) {

		case string:
			if arg == "" {
				return true
			}
		case int:
			if arg == 0 {
				return true
			}
		}
	}

	return false
}
