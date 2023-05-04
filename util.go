package datagen

import (
	"math"
	"math/rand"
	"time"
)

var grand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Seek(seed int64) {
	grand.Seed(seed)
}

// [min,max] not [min,max)
func randInt(min, max int) int {
	n := max - min
	if n <= 0 {
		return min
	}
	return min + grand.Intn(n+1)
}

// [min,max)
func randFloat64(min, max float64) float64 {
	n := max - min
	if n == 0 {
		return min
	}
	return min + grand.Float64()*n
}

func randRune(start, end rune) rune {
	return rune(randInt(int(start), int(end)))
}

func randDigit() rune {
	return rune(byte(grand.Intn(10)) + '0')
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Floor(num*output)) / output
}

func pick[T any](slices []T) T {
	i := grand.Intn(len(slices))
	return slices[i]
}

// n = [min,max]
// len(n)==0
func randCount(n ...int) int {
	if len(n) == 0 {
		return -1
	}
	return randInt(n[0], n[len(n)-1])
}
