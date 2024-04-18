package units

/**
 * Copyright (c) 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

import (
	"math"
)

// Speed represents a quantity of speed.
// Internally, it stores the speed in m/s, as a float64.
// Create a new Speed by multiplying one of the constants:
// d1 := 100 * units.Knot
// d2 := 10 * units.Kph
// Use methods like Knots() or Ki() to get the distance in the unit of your choice.
// US and UK spelling of consts and methods supported.
// Use Metres() or Kilometres() to get the distance in the unit of your choice.
// If you prefer imperial units, use NauticalMiles() Miles() or Feet().
// US and UK spelling supported.
// Create a new Distance by multiplying one of the constants:
// d := 10 * units.Metre

type Speed float64

const (
	Mps  Speed = 1
	Kn   Speed = Speed(float64(NauticalMile) / 3600)
	Kph  Speed = Speed(float64(Kilometre) / 3600)
	Mph  Speed = Speed(float64(Mile) / 3600)
)

// Valid returns true if the speed is valid. Invalid distances may be returned by
// functions when the result cannot be calculated.
func (s Speed) Valid() bool {
	return !math.IsNaN(float64(s))
}

// Kphs returns the speed in kilometres per hour.
func (s Speed) Kphs() float64 {
	return float64(s) * 3600 / float64(Kilometre)
}

// Kns returns the speed in Knots.
func (s Speed) Kns() float64 {
	return float64(s) * 3600 / float64(NauticalMile)
}

// Mpss return the speed in m/s.
func (s Speed) Mpss() float64 {
	return float64(s)
}
