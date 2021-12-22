package main

func asStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	// 创建一个unbuffeered的channel
	s := make(chan interface{})
	go func() {
		defer close(s)
		for _, v := range values {
			select {
			case <-done:
				return
			case s <- v:
			}
		}
	}()
	return s
}

func main() {

}
