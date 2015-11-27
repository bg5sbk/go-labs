package labs31

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var benchx int

func init() {
	rand.Seed(time.Now().UnixNano())
	benchx = rand.Intn(1 << 20)
}

func Test_All(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		n := rand.Intn(1<<20) + 2
		a := Normal(n)
		b := Switch(n)
		c := IF1(n)
		d := IF2(n)
		e := Search(n)
		f := IF3(n)
		if f > 21 {
			f = 21
		}
		if a != b || b != c || c != d || d != e || e != f {
			t.Log(n, a, b, c, d, e, f)
			t.Fail()
		}
	}
}

func Benchmark_Normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Normal(benchx)
	}
}

func Benchmark_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Search(benchx)
	}
}

func Benchmark_Switch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Switch(benchx)
	}
}

func Benchmark_IF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IF1(benchx)
	}
}

func Benchmark_IF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IF2(benchx)
	}
}

func Benchmark_IF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IF3(benchx)
	}
}

func Normal(n int) uint {
	var i = uint(0)
	for ; n > (1<<i) && i <= 20; i++ {
	}
	return i
}

var classes = []int{
	1 << 0, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7, 1 << 8, 1 << 9, 1 << 10,
	1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15, 1 << 16, 1 << 17, 1 << 18, 1 << 19, 1 << 20,
}

func Search(n int) uint {
	return uint(sort.SearchInts(classes, n))
}

func Switch(n int) uint {
	switch {
	case n == 1:
		return 0
	case n == 2:
		return 1
	case n > 1<<1 && n <= 1<<2:
		return 2
	case n > 1<<2 && n <= 1<<3:
		return 3
	case n > 1<<3 && n <= 1<<4:
		return 4
	case n > 1<<4 && n <= 1<<5:
		return 5
	case n > 1<<5 && n <= 1<<6:
		return 6
	case n > 1<<6 && n <= 1<<7:
		return 7
	case n > 1<<7 && n <= 1<<8:
		return 8
	case n > 1<<8 && n <= 1<<9:
		return 9
	case n > 1<<9 && n <= 1<<10:
		return 10
	case n > 1<<10 && n <= 1<<11:
		return 11
	case n > 1<<11 && n <= 1<<12:
		return 12
	case n > 1<<12 && n <= 1<<13:
		return 13
	case n > 1<<13 && n <= 1<<14:
		return 14
	case n > 1<<14 && n <= 1<<15:
		return 15
	case n > 1<<15 && n <= 1<<16:
		return 16
	case n > 1<<16 && n <= 1<<17:
		return 17
	case n > 1<<17 && n <= 1<<18:
		return 18
	case n > 1<<18 && n <= 1<<19:
		return 19
	case n > 1<<19 && n <= 1<<20:
		return 20
	}
	return 21
}

func IF1(n int) uint {
	if n <= 1<<0 {
		return 0
	}
	if n <= 1<<1 {
		return 1
	}
	if n <= 1<<2 {
		return 2
	}
	if n <= 1<<3 {
		return 3
	}
	if n <= 1<<4 {
		return 4
	}
	if n <= 1<<5 {
		return 5
	}
	if n <= 1<<6 {
		return 6
	}
	if n <= 1<<7 {
		return 7
	}
	if n <= 1<<8 {
		return 8
	}
	if n <= 1<<9 {
		return 9
	}
	if n <= 1<<10 {
		return 10
	}
	if n <= 1<<11 {
		return 11
	}
	if n <= 1<<12 {
		return 12
	}
	if n <= 1<<13 {
		return 13
	}
	if n <= 1<<14 {
		return 14
	}
	if n <= 1<<15 {
		return 15
	}
	if n <= 1<<16 {
		return 16
	}
	if n <= 1<<17 {
		return 17
	}
	if n <= 1<<18 {
		return 18
	}
	if n <= 1<<19 {
		return 19
	}
	if n <= 1<<20 {
		return 20
	}
	return 21
}

func IF2(n int) uint {
	if n <= 1<<10 {
		if n <= 1<<5 {
			if n <= 1<<2 {
				if n <= 1<<1 {
					return 1
				}
				return 2
			} else {
				if n <= 1<<4 {
					if n <= 1<<3 {
						return 3
					}
					return 4
				}
				return 5
			}
		} else {
			if n <= 1<<7 {
				if n <= 1<<6 {
					return 6
				}
				return 7
			} else {
				if n <= 1<<9 {
					if n <= 1<<8 {
						return 8
					}
					return 9
				}
				return 10
			}
		}
	} else {
		if n <= 1<<15 {
			if n <= 1<<12 {
				if n <= 1<<11 {
					return 11
				}
				return 12
			} else {
				if n <= 1<<14 {
					if n <= 1<<13 {
						return 13
					}
					return 14
				}
				return 15
			}
		} else {
			if n <= 1<<17 {
				if n <= 1<<16 {
					return 16
				}
				return 17
			} else {
				if n <= 1<<19 {
					if n <= 1<<18 {
						return 18
					}
					return 19
				}
				if n <= 1<<20 {
					return 20
				}
			}
		}
	}
	return 21
}

func IF3(n int) uint {
	n--
	c := uint(0)
	if (n & 0xffff0000) != 0 {
		c += 16
		n >>= 16
	}
	if (n & 0xff00) != 0 {
		c += 8
		n >>= 8
	}
	if (n & 0xf0) != 0 {
		c += 4
		n >>= 4
	}
	if (n & 0xc) != 0 {
		c += 2
		n >>= 2
	}
	if (n & 0x2) != 0 {
		c++
		n >>= 1
	}
	if n != 0 {
		c++
	}
	return c
}
