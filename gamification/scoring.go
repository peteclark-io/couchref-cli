package gamification

import "math"

func Score(vote bool, result float64, confidence float64) float64 {
	userVote := 1.00
	if !vote {
		userVote = 0.00
	}

	baseScore := math.Abs(userVote - result)

	confidentScore := baseScore * confidence

	splitScore := confidentScore - 0.500

	finalScore := -10.00 * splitScore
	if finalScore > 5 {
		return 5.00
	}

	return finalScore
}
