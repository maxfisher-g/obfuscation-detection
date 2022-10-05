package stats

import (
	"math"
	"obfuscation-detection/utils"
	"testing"
)

func flatten(s SampleStatistics) (data [8]float64) {
	data[0] = s.Mean
	data[1] = s.Variance
	data[2] = s.Skewness
	for i := 0; i < 5; i++ {
		data[i+3] = s.Quartiles[i]
	}
	return
}

func areSummariesEqual(s1, s2 SampleStatistics, absTol float64) bool {
	if s1.Size != s2.Size {
		return false
	}
	s1Data := flatten(s1)
	s2Data := flatten(s2)
	for i := 0; i < len(s1Data); i++ {
		if !utils.FloatEquals(s1Data[i], s2Data[i], absTol) {
			return false
		}
	}
	return true
}

func TestSummary(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual := CalculateSampleStatistics(data)
	expected := SampleStatistics{
		Size:      9,
		Mean:      5,
		Variance:  7.5,
		Skewness:  0,
		Quartiles: [5]float64{1, 3, 5, 7, 9},
	}
	if !areSummariesEqual(actual, expected, 1e-4) {
		t.Errorf("Expected summary: %v\nactual summary %v\n", expected, actual)
	}

}

func TestSummary2(t *testing.T) {
	data := []int{36, 7, 40, 41, 6, 42, 43, 47, 49, 15, 39}
	actual := CalculateSampleStatistics(data)
	expected := SampleStatistics{
		Size:      11,
		Mean:      33.18181818181818,
		Variance:  251.9636363636363,
		Skewness:  -1.0634150819204964,
		Quartiles: [5]float64{6, 15, 40, 43, 49},
	}
	if !areSummariesEqual(actual, expected, 1e-4) {
		t.Errorf("Expected summary: %v\nactual summary %v\n", expected, actual)
	}
}

func TestSummary3(t *testing.T) {
	data := []int{36, 40, 7, 39, 15, 41}
	actual := CalculateSampleStatistics(data)
	expected := SampleStatistics{
		Size:      6,
		Mean:      29.666666666666668,
		Variance:  218.26666666666665,
		Skewness:  -1.039599522561593,
		Quartiles: [5]float64{7, 15, 37.5, 40, 41},
	}
	if !areSummariesEqual(actual, expected, 1e-4) {
		t.Errorf("Expected summary: %v\nactual summary: %v\n", expected, actual)
	}
}

func TestSummary4(t *testing.T) {
	var data []int
	actual := CalculateSampleStatistics(data)
	nan := math.NaN()
	expected := SampleStatistics{
		Size:      0,
		Mean:      nan,
		Variance:  nan,
		Skewness:  nan,
		Quartiles: [5]float64{nan, nan, nan, nan, nan},
	}
	if !areSummariesEqual(actual, expected, 1e-4) {
		t.Errorf("Expected summary: %v\nactual summary %v\n", expected, actual)
	}
}

func TestSummary5(t *testing.T) {
	data := []float64{1.5}
	actual := CalculateSampleStatistics(data)
	nan := math.NaN()
	expected := SampleStatistics{
		Size:      1,
		Mean:      1.5,
		Variance:  nan,
		Skewness:  nan,
		Quartiles: [5]float64{1.5, 1.5, 1.5, 1.5, 1.5},
	}
	if !areSummariesEqual(actual, expected, 1e-4) {
		t.Errorf("Expected summary: %v\nactual summary %v\n", expected, actual)
	}
}

func TestSummary6(t *testing.T) {
	data := []float64{1.5, 2.5}
	actual := CalculateSampleStatistics(data)
	nan := math.NaN()
	expected := SampleStatistics{
		Size:      2,
		Mean:      2.0,
		Variance:  0.5,
		Skewness:  nan,
		Quartiles: [5]float64{1.5, 1.5, 2.0, 2.5, 2.5},
	}
	if !areSummariesEqual(actual, expected, 1e-4) {
		t.Errorf("Expected summary: %v\nactual summary %v\n", expected, actual)
	}
}

func TestSummary7(t *testing.T) {
	data := []float64{-12.5, 0, 12.5}
	actual := CalculateSampleStatistics(data)
	expected := SampleStatistics{
		Size:      3,
		Mean:      0.0,
		Variance:  156.25,
		Skewness:  0,
		Quartiles: [5]float64{-12.5, -12.5, 0.0, 12.5, 12.5},
	}
	if !areSummariesEqual(actual, expected, 1e-4) {
		t.Errorf("Expected summary: %v\nactual summary %v\n", expected, actual)
	}
}
