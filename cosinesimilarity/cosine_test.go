package cosinesimilarity

import (
	"fmt"
	"testing"
)

func TestCosine(t *testing.T) {

	a := []float64{1, 1, 0, 0, 1}
	b := []float64{1, 1, 0, 0, 1}
	cos, err := GetInnerProduct(a, b, 5)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(`=====>`, cos)
	if cos < 0.99 {
		t.Error("Expected similarity of 1, got instead ", cos)
	}

	a = []float64{2, 2, 2}
	b = []float64{2, 2, 2}
	cos, err = GetInnerProduct(a, b, 3)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(`=====>`, cos)
	if cos < 0.49999 || cos > 0.5 {
		t.Error("Expected similarity of 0.5, got instead ", cos)
	}
}
