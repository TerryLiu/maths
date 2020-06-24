/*
maths computes accurate running mean, variance, and standard deviation

It is based on code by John D Cook: http://www.johndcook.com/blog/standard_deviation/
*/
package maths

import (
	"math"
)

type Swatch struct {
	// 元素个数
	N uint
	//最新的平均值
	NewM float64
	// 上一个平均值
	OldM float64
	// 最新的样本差平方和
	NewS float64
	// 上一个样本差平方和
	OldS float64
	// 标准差
	StdVal float64
}

func NewSwatch() *Swatch {
	return &Swatch{}
}
// 批量插入元素
func (this *Swatch) PutList(y []float64) *Swatch {
	if this == nil {
		return nil
	}
	for i, _ := range y {
		this.Push(y[i])
	}
	return this
}
//插入单个元素
func (this *Swatch) Push(x float64) *Swatch {
	if this == nil {
		return nil
	}
	this.N++

	// See Knuth TAOCP vol 2, 3rd edition, page 232
	if this.N == 1 {
		this.OldM = x
		this.NewM = x
		this.OldS = 0.0
	} else {
		this.NewM = this.OldM + (x-this.OldM)/float64(this.N)
		this.NewS = this.OldS + (x-this.OldM)*(x-this.NewM)

		// set up for next iteration
		this.OldM = this.NewM
		this.OldS = this.NewS
	}
	return this
}

// 样本数量
func (this *Swatch) NumOfData() uint {
	if this == nil {
		return 0
	}
	return this.N
}

// 算术平均值
func (this *Swatch) Ma() float64 {
	if this == nil {
		return 0
	}
	return this.NewM
}

//方差
func (this *Swatch) Variance() float64 {
	if this.N > 1 {
		return this.NewS / (float64(this.N) - 1)
	}

	return 0.0
}

// 标准偏差
func (this *Swatch) Std() float64 {
	if this == nil {
		return 0
	}
	this.StdVal=math.Sqrt(this.Variance())
	return this.StdVal
}

// 布林线上轨
func (this *Swatch) Upper() float64 {
	if this == nil {
		return 0
	}
	if this.StdVal == 0 {
		this.StdVal=math.Sqrt(this.Variance())
	}
	return this.NewM+this.StdVal*2
}
// 布林线下轨
func (this *Swatch) Lower() float64 {
	if this == nil {
		return 0
	}
	if this.StdVal == 0 {
		this.StdVal=math.Sqrt(this.Variance())
	}
	return this.NewM-this.StdVal*2
}