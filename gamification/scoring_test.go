package gamification

import "testing"

func TestScore(t *testing.T) {
	final := Score(true, 0.8762517, 0.98172321)
	t.Log(true, 0.8762517, 0.98172321)
	t.Log(final)

	final = Score(false, 0.8762517, 0.98172321)
	t.Log(false, 0.8762517, 0.98172321)
	t.Log(final)

	final = Score(true, 0.50000000001, 1)
	t.Log(true, 0.50000000001, 1)
	t.Log(final)

	final = Score(false, 0.50000000001, 1)
	t.Log(false, 0.50000000001, 1)
	t.Log(final)

	final = Score(true, 0.6412316, 1)
	t.Log(true, 0.6412316, 1)
	t.Log(final)

	final = Score(false, 0.6412316, 1)
	t.Log(false, 0.6412316, 1)
	t.Log(final)

	final = Score(false, 1, 1)
	t.Log(false, 1, 1)
	t.Log(final)

	final = Score(true, 1, 1)
	t.Log(true, 1, 1)
	t.Log(final)

	final = Score(true, 0.5, 1)
	t.Log(true, 0.5, 1)
	t.Log(final)

	final = Score(false, 0.5, 1)
	t.Log(false, 0.5, 1)
	t.Log(final)
}
