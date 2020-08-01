package statistic

import (
	"testing"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}

	}
}

var testsSum = []testpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestAverageSum(t *testing.T) {
	for _, pairSum := range testsSum {
		v := AverageSum(pairSum.values)
		if v != pairSum.average {
			t.Error(
				"For", pairSum.values,
				"expected", pairSum.average,
				"got", v,
			)
		}

	}

}
