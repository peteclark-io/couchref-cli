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

	bumpedScore := -20.00 * splitScore

	var finalScore float64
	if bumpedScore > 0 {
		finalScore = math.Log(bumpedScore) + 2
	} else {
		finalScore = bumpedScore
	}

	return finalScore
}