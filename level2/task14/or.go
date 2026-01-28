package task14

func or(ch ...<-chan interface{}) <-chan interface{} {
	if len(ch) == 1 {
		return ch[0]
	}

	done := make(chan interface{})

	if len(ch) == 0 {
		close(done)
		return done
	}

	go func() {
		select {
		case <-ch[0]:

		case <-or(ch[1:]...):
		}
		close(done)
	}()

	return done
}
