package main

import "time"

// WaitForSignal は done が d 以内に閉じられれば true を、
// d を過ぎても来なければ（タイムアウト）false を返す。
//
// time.After によるタイマーを含むため、通常のテストでは
// 「タイムアウトを確認する」のに実時間 d だけ待つ必要がある。
// STEP 02 では synctest の fake clock でこの待ちを消す。
func WaitForSignal(done <-chan struct{}, d time.Duration) bool {
	select {
	case <-done:
		return true
	case <-time.After(d):
		return false
	}
}

func main() {}
