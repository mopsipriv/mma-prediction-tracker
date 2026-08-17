package service

import (
	"math"
	"testing"
)


const tolerance = 0.0001

func TestCalculateBrierScore(t *testing.T) {
	
	tests := []struct {
		name        string  
		probability float64 
		isCorrect   bool    
		expected    float64 
	}{
		{
			name:        "High confidence correct prediction",
			probability: 0.80,
			isCorrect:   true,
			expected:    0.04, // (0.80 - 1.0)^2 = 0.04
		},
		{
			name:        "High confidence wrong prediction",
			probability: 0.80,
			isCorrect:   false,
			expected:    0.64, // (0.80 - 0.0)^2 = 0.64
		},
		{
			name:        "Neutral prediction correct",
			probability: 0.50,
			isCorrect:   true,
			expected:    0.25, // (0.50 - 1.0)^2 = 0.25
		},
		{
			name:        "Perfect prediction",
			probability: 1.0,
			isCorrect:   true,
			expected:    0.0, // (1.0 - 1.0)^2 = 0.0
		},
	}

	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateBrierScore(tt.probability, tt.isCorrect)

			
			if math.Abs(got-tt.expected) > tolerance {
				t.Errorf("CalculateBrierScore(%v, %v) = %v; expected %v",
					tt.probability, tt.isCorrect, got, tt.expected)
			}
		})
	}
}

func TestCalculateAccuracy(t *testing.T) {
	tests := []struct {
		name               string
		correctPredictions int
		totalPredictions   int
		expected           float64
	}{
		{
			name:               "Zero total predictions (division by zero protection)",
			correctPredictions: 0,
			totalPredictions:   0,
			expected:           0.0,
		},
		{
			name:               "50 percent accuracy",
			correctPredictions: 5,
			totalPredictions:   10,
			expected:           50.0,
		},
		{
			name:               "100 percent accuracy",
			correctPredictions: 4,
			totalPredictions:   4,
			expected:           100.0,
		},
		{
			name:               "0 percent accuracy",
			correctPredictions: 0,
			totalPredictions:   5,
			expected:           0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAccuracy(tt.correctPredictions, tt.totalPredictions)

			if math.Abs(got-tt.expected) > tolerance {
				t.Errorf("CalculateAccuracy(%d, %d) = %v; expected %v",
					tt.correctPredictions, tt.totalPredictions, got, tt.expected)
			}
		})
	}
}