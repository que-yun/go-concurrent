package main

/**
使用 go run --race map_data_race.go  会出现data race问题
*/

func main() {
	var m = make(map[int]int, 10) // 初始化一个map
	go func() {
		for {
			m[1] = 1 //设置key
		}
	}()
	go func() {
		for {
			_ = m[2] //访问这个map
		}
	}()
	select {}
}
