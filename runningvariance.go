/*
maths computes accurate running mean, variance, and standard deviation

It is based on code by John D Cook: http://www.johndcook.com/blog/standard_deviation/
*/
package main

import (
	"math"
)

type RunningStat struct {
	// 元素个数
	N    uint
	//最新的平均值
	NewM float64
	// 上一个平均值
	OldM float64
	// 最新的
	NewS float64
	OldS float64
}

func NewRunningStat() *RunningStat {
	return &RunningStat{}
}

func (r *RunningStat) Push(x float64) {
	r.N++

	// See Knuth TAOCP vol 2, 3rd edition, page 232
	if r.N == 1 {
		r.OldM = x
		r.NewM = x
		r.OldS = 0.0
	} else {
		r.NewM = r.OldM + (x-r.OldM)/float64(r.N)
		r.NewS = r.OldS + (x-r.OldM)*(x-r.NewM)

		// set up for next iteration
		r.OldM = r.NewM
		r.OldS = r.NewS
	}
}
// 样本数量
func (r *RunningStat) NumDataValues() uint {
	return r.N
}
// 均值
func (r *RunningStat) Mean() float64 {
	return r.NewM
}
//方差
func (r *RunningStat) Variance() float64 {
	if r.N > 1 {
		return r.NewS / (float64(r.N) - 1)
	}

	return 0.0
}
// 标准偏差
func (r *RunningStat) StandardDeviation() float64 {
	return math.Sqrt(r.Variance())
}
