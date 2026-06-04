package main

import "time"

// Classify は parallel/steps と同じく、わざと 1 秒 Sleep する。
// 通常のテストではこの Sleep の分だけ実時間がかかるが、
// STEP 01 では synctest を使ってこの待ち時間を「消す」。
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

func main() {}
