package batch

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	errG, _ := errgroup.WithContext(context.Background())
	errG.SetLimit(int(pool))
	chUsers := make(chan user, n)

	for i := int64(0); i < n; i++ {
		k := i

		errG.Go(func() error {
			func(j int64) {

				item := getOne(j)
				chUsers <- item
			}(k) // i -> k -> j
			return nil
		})
	}

	errG.Wait()
	close(chUsers)

	for item := range chUsers {
		res = append(res, item)
	}

	return
}
