// Package galaxygen generates realistic-looking star galaxies.
//
// This algorithm was originally made in c++, by:
//
// Bryon Mau (www.cyberbrinedreams.com)
//
// And was then modified (in c++) by:
//
// Patrick Down (www.codemoon.com)
//
// This was done years ago and both sites now seems to be down.
//
// This go port was made from the salvaged parts from my shitty python port
// (done ages ago too), that was based on the original c++ source.
// It has now been cleaned up and made easier to read.
package galaxygen

import (
	"math"
	"math/rand"
)

const (
	deg2rad float64 = math.Pi / 180.0
)

// Star is nothing more but a point in 2D space.
type Star struct {
	X int
	Y int
}

// New generates a new galaxy filled with "stars".
// It lazely try to make x stars (but it's usually lower due to removed duplicates).
func New(stars, arms int, radius float64) []*Star {
	if arms < 1 {
		// Avoid division by zero errors.
		arms = 1
	}
	armangle := math.Mod(360.0/float64(arms), 360.0)
	angularspread := 250.0 / float64(arms)

	// The direction the galaxy spins in.
	direction := -1.0
	if rand.Intn(2) == 1 {
		direction = 1.0
	}

	var starmap []*Star
	for i := 0; i < stars; i++ {
		// Jitter makes the galaxy be a little more spread out, in a
		// "jittery" manner, and makes it be more "twirly".
		jitter := -1.0
		if rand.Intn(2) == 1 {
			jitter = 1.0
		}

		rad := hatRandom(radius)
		spread := hatRandom(angularspread) * jitter
		angle := float64(rand.Intn(arms)) * armangle

		// Mathematical!
		x := int(rad * math.Cos(deg2rad*(angle+rad*direction+spread)))
		y := int(rad * math.Sin(deg2rad*(angle+rad*direction+spread)))
		starmap = append(starmap, &Star{x, y})
	}

	// Filter out any duplicates.
	var uniques []*Star
	for _, s := range starmap {
		if hasStar(s, uniques) {
			continue
		}
		uniques = append(uniques, s)
	}
	return uniques
}

// hasStar checks if a slice of *Star already contains a star at X,Y.
func hasStar(val *Star, list []*Star) bool {
	for _, s := range list {
		if s.X == val.X && s.Y == val.Y {
			return true
		}
	}
	return false
}

// hatRandom makes... a new float, somehow, from another float.
// I have no idea why it's called hatRandom.
func hatRandom(rang float64) float64 {
	// couldn't use "range" for the variable because of "reserved go keyword"
	// or something silly, hence now it's called "rang"...
	area := 4 * math.Atan(6.0)
	p := area * rand.Float64()
	return math.Tan(p/4.0) * rang / 6.0
}
