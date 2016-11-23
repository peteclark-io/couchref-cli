package gamification

import "math"

func Score(vote bool, result float64, confidence float64) float64 {
	userVote := 1.00
	if !vote {
		userVote = 0.00
	}

	baseScore := math.Abs(userVote - result)

	splitScore := baseScore - 0.500

	bumpedScore := -30.00 * splitScore

	var finalScore float64
	if bumpedScore > 0 {
		finalScore = math.Log2(bumpedScore + 2)
	} else {
		finalScore = bumpedScore
	}

	return finalScore * confidence
}
