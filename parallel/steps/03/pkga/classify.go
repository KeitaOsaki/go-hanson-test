package pkga

import "time"

func Classify(n int) string {
	time.Sleep(1000 * time.Millisecond)
	switch {
	case n < 0:
		return "neg"
	case n == 0:
		return "zero"
	default:
		return "pos"
	}
}
