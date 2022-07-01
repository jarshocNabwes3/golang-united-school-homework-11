package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	sem := make(chan bool, pool)
	chUsers := make(chan user, n)

	for i := int64(0); i < n; i++ {
		wg.Add(1)
		sem <- true

		go func(j int64) {
			defer func() { <-sem }()
			defer wg.Done()

			item := getOne(j)
			chUsers <- item
		}(i)
	}

	wg.Wait()
	close(chUsers)

	for item := range chUsers {
		res = append(res, item)
	}

	return
}
