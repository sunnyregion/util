package cosinesimilarity

import "errors"

//GetInnerProduct 计算内积值
func GetInnerProduct(a, b []float64, l int) (ip float64, err error) {
	ip = 0

	for i := 0; i < l; i++ {
		ip += a[i] * b[i]
	}
	if ip == 0 {
		err = errors.New(`Inner product value error`)
	}
	return
}
