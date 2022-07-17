package batch

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getButch(t *testing.T) {
	type args struct {
		n    int64
		pool int64
	}
	tests := []struct {
		args    args
		wantRes []user
	}{
		{args: args{n: 10, pool: 1}, wantRes: createRes(10)},
		{args: args{n: 10, pool: 2}, wantRes: createRes(10)},
		{args: args{n: 10, pool: 5}, wantRes: createRes(10)},
		{args: args{n: 20, pool: 4}, wantRes: createRes(20)},
		{args: args{n: 100, pool: 10}, wantRes: createRes(100)},
		{args: args{n: 15, pool: 5}, wantRes: createRes(15)},
		{args: args{n: 35, pool: 5}, wantRes: createRes(35)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			start := time.Now()
			const ms = int64(time.Millisecond)
			abs := func(inp int64) int64 {
				return int64(math.Abs(float64(inp)))
			}

			wanted := ms * tt.args.n / tt.args.pool * 100
			actualRes := getBatch(tt.args.n, tt.args.pool)
			actual := int64(time.Since(start))

			assert.LessOrEqual(t, abs(wanted-actual), ms*200)
			assert.ElementsMatch(t, tt.wantRes, actualRes)
		})
	}
}

func createRes(v int64) []user {
	res := make([]user, 0, v)
	for i := 0; i < int(v); i++ {
		res = append(res, user{ID: int64(i)})
	}
	return res
}
