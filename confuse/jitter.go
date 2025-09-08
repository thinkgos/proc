package confuse

import "math/rand/v2"

// Jitter Returns a random value from the following cardinal.
// cardinal: the cardinal value.
// randomizationFactor: the randomization factor, It should be in the range (0, 1].
// random: the random value, It should be in the range (0, 1]. always use rand.Float64()
//
//	value rand in [cardinal - randomizationFactor * cardinal, cardinal + randomizationFactor * cardinal].
func Jitter(cardinal int, randomizationFactor, random float64) int {
	if randomizationFactor <= 0 || randomizationFactor > 1 {
		return cardinal
	}
	var delta = randomizationFactor * float64(cardinal)
	var minCardinal = float64(cardinal) - delta
	var maxCardinal = float64(cardinal) + delta

	// Get a random value from the range [minCardinal, maxCardinal].
	// The formula used below has a +1 because if the minCardinal is 1 and the maxCardinal is 3 then
	// we want a 33% chance for selecting either 1, 2 or 3.
	return int(minCardinal + (random * (maxCardinal - minCardinal + 1)))
}

// Jitter2 Returns a random value from the following cardinal.
// see [Jitter]
func Jitter2(cardinal int, randomizationFactor float64) int {
	return Jitter(cardinal, randomizationFactor, rand.Float64())
}
