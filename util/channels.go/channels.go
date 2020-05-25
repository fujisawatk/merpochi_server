package channels

// OK channelの判別処理
func OK(done <-chan bool) bool {
	select {
	case ok := <-done:
		if ok {
			return true
		}
	}
	return false
}
