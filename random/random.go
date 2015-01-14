package random

import (
	"math/rand"
	"time"
        "fmt"
)

type RandByTime struct {
	r *rand.Rand
}

func NewRandByTime() *RandByTime {
	rt := &RandByTime{}
	rt.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return rt
}

func (this *RandByTime) RandIntArray(size, maxnum int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = this.r.Intn(maxnum)
	}
        return arr
}

func (this *RandByTime) RandIntString(size, maxnum int) string {
	s := ""
	for i := 0; i < size; i++ {
		s += fmt.Sprintf("%d", this.r.Intn(maxnum))
	}
	return s
}
