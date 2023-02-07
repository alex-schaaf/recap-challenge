package utils

// SimpleReturn calculates the ratio between the end market value and beginning market value.
// Assumes that the given data is sorted chronologically from newest to oldest
func SimpleReturn(data []float64) float64 {
	bgm := data[len(data)-1]
	emv := data[0]
	return (emv/bgm - 1) * 100
}

// MaximumDrawdown calculates the largest difference from a maximum to a later minimum in percent.
func MaximumDrawdown(data []float64) float64 {
	var drawdowns []float64
	for i := 0; i < len(data)-2; i++ {
		cm := CumMin(data[i+1 : len(data)-1])
		// drawdown := cm/data[i] - 1
		drawdown := (cm - data[i]) / data[i]
		drawdowns = append(drawdowns, drawdown)
	}

	if len(drawdowns) == 0 {
		panic("Can't calculate drawdown of slice with single value")
	}

	maxDrawdown := drawdowns[0]
	for _, dd := range drawdowns[1 : len(drawdowns)-1] {
		if dd < maxDrawdown {
			maxDrawdown = dd
		}
	}

	return maxDrawdown * 100
}

// CumMin calculates the cumulative minimum value.
func CumMin(data []float64) float64 {
	if len(data) == 0 {
		panic("Can't calculate cumulative minimum of empty slice")
	} else if len(data) == 1 {
		return data[0]
	}
	smallest := data[0]
	for _, value := range data[1 : len(data)-1] {
		if value < smallest {
			smallest = value
		}
	}
	return smallest
}
