package gamification

import "testing"

func TestScore(t *testing.T) {
	final := score(true, 0.8762517, 0.98172321)
	t.Log(final)

	final = score(false, 0.8762517, 0.98172321)
	t.Log(final)

	final = score(true, 0.5117778, 1)
	t.Log(final)

	final = score(false, 0.5117778, 1)
	t.Log(final)

	final = score(true, 0.6412316, 1)
	t.Log(final)

	final = score(false, 0.6412316, 1)
	t.Log(final)
}
