package service
import (
	"math"
)

func CalculateBrierScore(probability float64, isCorrect bool) float64{
	var actual float64 = 0.0
	if isCorrect{
		actual = 1.0
	}
	diff := probability - actual
	return math.Pow(diff,2)
}

func CalculateAccuracy(correctPredictions, totalPredictions int) float64{
	if totalPredictions == 0{
		return 0.0
	}

	return (float64(correctPredictions) / float64(totalPredictions)) * 100.0
}