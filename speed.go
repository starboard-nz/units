package units

/**
 * Copyright (c) 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Speed represents a quantity of speed.
type Speed interface {
	Unit
	// Knots returns the speed in knots
	Knot() Knot
	// Kphs returns the speed in kilometers per hour.
	Kph() Kph
	// Mps returns the speed in metres per second
	Mps() Mps
	// Mph returns the speed in miles per hour
	Mph() Mph
}

type (
	Knot float64 // kn
	Kph  float64 // km/h
	Mps  float64 // m/s
	Mph  float64 // mi/h
)

// ParseSpeed parses a string that contains a speed and a commonly used unit abbreviation
// and returns it as a distance unit.
func ParseSpeed(sp string) (Speed, error) {
	tokens := strings.Fields(sp)
	if len(tokens) == 1 {
		tokens = []string{
			strings.TrimRightFunc(tokens[0], func(r rune) bool {
				return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == '/'
			}),
			strings.TrimLeftFunc(tokens[0], func(r rune) bool {
				return r >= '0' && r <= '9' || r == '-' || r == '.'
			}),
		}
	}

	if len(tokens) != 2 {
		return Mps(0), ErrParse
	}

	val, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return Mps(0), err
	}

	switch tokens[1] {
	case "m/s":
		return Mps(val), nil
	case "kn":
		return Knot(val), nil
	case "km/h":
		return Kph(val), nil
	case "mph", "mi/h":
		return Mph(val), nil
	default:
		return Mps(0), ErrUnknownUnit
	}
}

// Knots
func (k Knot) Name() string {
	return "knots"
}

func (k Knot) Short() string {
	return "kn"
}

func (k Knot) String() string {
	return fmt.Sprintf("%f kn", k)
}

func (k Knot) Valid() bool {
	return !math.IsNaN(float64(k))
}

func (k Knot) Kph() Kph {
	return Kph(float64(k) * float64(NauticalMileInMetres) / 1000)
}

func (k Knot) Knot() Knot {
	return k
}

func (k Knot) Mps() Mps {
	return Mps(float64(k) * float64(NauticalMileInMetres) / 3600)
}

func (k Knot) Mph() Mph {
	return Mph(float64(k) * float64(NauticalMileInMetres) / float64(MileInMetres))
}

// km/h
func (k Kph) Name() string {
	// FIXME: should return spelling according to locale
	return "Kilometres per hour"
}

func (k Kph) Short() string {
	return "km/h"
}

func (k Kph) String() string {
	return fmt.Sprintf("%f km/h", k)
}

func (k Kph) Valid() bool {
	return !math.IsNaN(float64(k))
}

func (k Kph) Kph() Kph {
	return k
}

func (k Kph) Knot() Knot {
	return Knot(float64(k) * 1000 / float64(NauticalMileInMetres))
}

func (k Kph) Mps() Mps {
	return Mps(k / 3.6)
}

func (k Kph) Mph() Mph {
	return Mph(float64(k) * 1000 / float64(MileInMetres))
}

// m/s
func (m Mps) Name() string {
	// FIXME: should return spelling according to locale
	return "metres per second"
}

func (m Mps) Short() string {
	return "m/s"
}

func (m Mps) String() string {
	return fmt.Sprintf("%f km/h", m)
}

func (m Mps) Valid() bool {
	return !math.IsNaN(float64(m))
}

func (m Mps) Kph() Kph {
	return Kph(m * 3.6)
}

func (m Mps) Knot() Knot {
	return Knot(float64(m) * 3600 / float64(NauticalMileInMetres))
}

func (m Mps) Mps() Mps {
	return m
}

func (m Mps) Mph() Mph {
	return Mph(float64(m) * 3600 / float64(MileInMetres))
}

// mph (mi/h)
func (m Mph) Name() string {
	return "miles per hour"
}

func (m Mph) Short() string {
	return "mph"
}

func (m Mph) String() string {
	return fmt.Sprintf("%f mph", m)
}

func (m Mph) Valid() bool {
	return !math.IsNaN(float64(m))
}

func (m Mph) Kph() Kph {
	return Kph(float64(m) * float64(MileInMetres) / 1000)
}

func (m Mph) Knot() Knot {
	return Knot(float64(m) * float64(MileInMetres) / float64(NauticalMileInMetres))
}

func (m Mph) Mps() Mps {
	return Mps(float64(m) * float64(MileInMetres) / 3600)
}

func (m Mph) Mph() Mph {
	return m
}
