package channels

func SendChanMsg(msg string, ch chan<- string) {
	go func() {
		ch <- msg
	}()
}



