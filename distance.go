package units

/**
 * Copyright (c) 2020, 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

import (
	"math"
	"strconv"
	"strings"
)

// Distance represents a quantity of length.
// Internally, it stores the distance in metres, as a float64.
// Create a new Distance by multiplying one of the constants:
// d1 := 100 * units.Metre
// d2 := 10 * units.NauticalMile
// Use methods like Metres() or Kilometres() to get the distance in the unit of your choice.
// US and UK spelling of consts and methods supported.

type Distance float64

const (
	Metre        Distance = 1
	Meter        Distance = 1
	Kilometer    Distance = 1000
	Kilometre    Distance = 1000
	Millimeter   Distance = 0.001
	Millimetre   Distance = 0.001
	Foot         Distance = 0.3048
	Yard         Distance = 0.9144
	Inch         Distance = 0.0254
	NauticalMile Distance = 1852
	Fathom       Distance = 1.8288
	Mile         Distance = 1609.344
)

// Valid returns true if the distance is valid. Invalid distances may be returned by
// functions when the result cannot be calculated.
func (d Distance) Valid() bool {
	return !math.IsNaN(float64(d))
}

// Metres returns the DistanceUnits d in metres
func (d Distance) Metres() float64 {
	return float64(d)
}

// Meters also returns the Distance d in meters 🇺🇸
func (d Distance) Meters() float64 {
	return float64(d)
}

// Kilometres returns the Distance d in kilometres
func (d Distance) Kilometres() float64 {
	return float64(d / Kilometre)
}

// Kilometers also returns the Distance d in kilometers 🇺🇸
func (d Distance) Kilometers() float64 {
	return float64(d / Kilometer)
}

// NauticalMiles returns the Distance d in nautical miles
func (d Distance) NauticalMiles() float64 {
	return float64(d / NauticalMile)
}

// Miles returns the Distance d in miles (land)
func (d Distance) Miles() float64 {
	return float64(d / Mile)
}

// Feet returns the Distance d in feet
func (d Distance) Feet() float64 {
	return float64(d / Foot)
}

func ParseDistance(dist string) (Distance, error) {
	tokens := strings.Fields(dist)
	if len(tokens) == 1 {
		tokens = []string{
			strings.TrimRightFunc(tokens[0], func(r rune) bool {
				return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == '"' || r == '\''
			}),
			strings.TrimLeftFunc(tokens[0], func(r rune) bool {
				return r >= '0' && r <= '9' || r == '-' || r == '.'
			}),
		}
	}

	if len(tokens) != 2 {
		return 0, ErrParse
	}

	val, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, err
	}

	switch tokens[1] {
	case "m":
		return Distance(val), nil
	case "NM":
		return Distance(val) * NauticalMile, nil
	case "mile", "miles":
		return Distance(val) * Mile, nil
	case "km":
		return Distance(val) * Kilometer, nil
	case "ft", "'":
		return Distance(val) * Foot, nil
	case "in", "inch", "\"":
		return Distance(val) * Inch, nil
	default:
		return 0, ErrUnknownUnit
	}
}
