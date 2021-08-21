package main

import (
	"math"
	"math/rand"
	"testing"
)

/**
 * @Author: yirufeng
 * @Date: 2021/8/21 5:53 下午
 * @Desc:
 **/

//unit test
func TestAdd(t *testing.T) {
	var a, b int
	for i := 0; i < 100; i++ {
		a, b = rand.Intn(math.MaxInt32>>20), rand.Intn(math.MaxInt32>>20)
		ret := Add(a, b)
		if ret != a+b {
			t.Fatalf("expected %d, but got %d", a+b, ret)
		}
	}
}
