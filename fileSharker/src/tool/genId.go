package tool

import (
	"math/rand"
	"sync"
	"time"
)

var (
	mu sync.Mutex
	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func CreateId() int64 {

	mu.Lock()
	defer mu.Unlock()

	var n int64
	dt := time.Now()

	y, m, d := dt.Date()
	y = getY(dt.Year())

	n = int64(y)                                    // 年：7位
	n = (n << 5) | int64(m)                         // 月：5位
	n = (n << 6) | int64(d)                         // 月：6位
	n = (n << 6) | int64(dt.Hour())                 // 月：6位
	n = (n << 7) | int64(dt.Minute())               // 月：7位
	n = (n << 7) | int64(dt.Second())               // 月：7位
	n = (n << 11) | int64(rand.Int31()&((1<<11)-1)) // 随机11位

	//yy, mm1, dd, hh, mm2, ss, rr := DecodeId(n)
	//
	//fmt.Printf("y:%d:%d\n", y, yy)
	//fmt.Printf("M:%d:%d\n", m, mm1)
	//fmt.Printf("d:%d:%d\n", d, dd)
	//fmt.Printf("h:%d:%d\n", time.Now().Hour(), hh)
	//fmt.Printf("m:%d:%d\n", time.Now().Minute(), mm2)
	//fmt.Printf("s:%d:%d\n", time.Now().Second(), ss)
	//fmt.Printf("r:%d:%d\n", r, rr)

	return n
}

func DecodeId(id int64) (y, m1, d, h, m2, s, r int) {
	var n = int(id)

	r = n & ((1 << 11) - 1)
	n = n >> 11

	s = n & ((1 << 7) - 1)
	n = n >> 7

	m2 = n & ((1 << 7) - 1)
	n = n >> 7

	h = n & ((1 << 6) - 1)
	n = n >> 6

	d = n & ((1 << 6) - 1)
	n = n >> 6

	m1 = n & ((1 << 5) - 1)
	n = n >> 5

	y = n & ((1 << 7) - 1)

	return
}

func getY(yyyy int) int {
	return yyyy - (yyyy/100)*100
}
