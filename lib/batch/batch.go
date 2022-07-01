package batch

import (
	"fmt"
	"sync"
	"time"
)

type user struct {
	ID int64
}

var start = time.Now()

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	fmt.Println(time.Since(start).Milliseconds())
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	// c := make(chan int64, pool)
	var mu sync.Mutex
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		// c <- i
		go func(j int64) {
			// j := <-c
			item := getOne(j)
			mu.Lock()
			res = append(res, item)
			// defer func() {
			mu.Unlock()
			// }()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res
}
