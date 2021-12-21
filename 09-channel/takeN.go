package main

func takeN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	// 创建输出流
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		// 只读取前num个数据
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- valueStream:
			}
		}
	}()
	return takeStream
}

func main() {

}
