package leetcode

import (
	"fmt"
	"testing"
)

type question16 struct {
	para16
	ans16
}

// para 是参数
// one 代表第一个参数
type para16 struct {
	a      []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans16 struct {
	one int
}

func Test_Problem16(t *testing.T) {

	qs := []question16{


		{
			para16{[]int{1, 1, 1, 2}, 1},
			ans16{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 16------------------------\n")

	for _, q := range qs {
		_, p := q.ans16, q.para16
		fmt.Printf("【input】:%v       【output】:%v\n", p, threeSumClosest(p.a, p.target))
	}
	fmt.Printf("\n\n\n")
}
