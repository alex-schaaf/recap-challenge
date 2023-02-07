package utils

import (
	"fmt"
	"testing"
)

func TestMaxDrawdownZero(t *testing.T) {
	result := MaximumDrawdown([]float64{2, 2, 2, 2})
	if result != 0 {
		t.Errorf("Drawdown of identical values should be zero")
	}
}
func TestMaxDrawdown(t *testing.T) {
	result := MaximumDrawdown([]float64{2, 1, 1, 1, 2, 0, 2, 2, 2})
	fmt.Println(result)
	if result != -100 {
		t.Errorf("Failed to identify second drawdown")
	}
}
func TestMaxDrawdownNope(t *testing.T) {
	result := MaximumDrawdown([]float64{1, 2, 3, 4, 5})
	fmt.Println(result)
	if result != 50 {
		t.Errorf("MaximumDrawdown of slice without actual drawdown should be positive")
	}
}
func TestMaxDrawdownSingleValue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MaximumDrawdown should panic if only single value in data slice")
		}
	}()

	MaximumDrawdown([]float64{2})
}

func TestCumMin(t *testing.T) {
	result := CumMin([]float64{10})
	if result != 10 {
		t.Errorf("CumMin of single element slice should be that element")
	}

	result = CumMin([]float64{10, 20, 5})
	if result != 5 {
		t.Errorf("CumMin failed to find minimum value in slice")
	}

	result = CumMin([]float64{10, 20, -1, 200})
	if result != -1 {
		t.Errorf("CumMin failed to find negative minimum value in slice")
	}
}
