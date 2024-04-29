package units

/**
 * Copyright (c) 2020, 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Distance represents a quantity of speed.
type Distance interface {
	Unit
	// Metre returns the distance in metres
	Metre() Metre
	// Meter returns the distance in metres
	Meter() Meter
	// Km returns the distance in Kilometres
	Km() Km
	// NM returns the distance in Nautical Miles
	NM() NM
	// Mile returns the distance in Miles
	Mile() Mile
}

// Weird Imperial distance units
const (
	FootInMetres         = Metre(0.3048)
	YardInMetres         = Metre(0.9144)
	InchInMetres         = Metre(0.0254)
	NauticalMileInMetres = Metre(1852)
	MileInMetres         = Metre(1609.344)
)

type (
	Metre float64
	Meter float64
	Km    float64
	NM    float64
	Mile  float64
)

// ParseDistance parses a string that contains a distance and a commonly used unit abbreviation
// and returns it as a distance unit.
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
		return Metre(0), ErrParse
	}

	val, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return Metre(0), err
	}

	switch tokens[1] {
	case "m":
		return Metre(val), nil
	case "NM", "nmi":
		return NM(val), nil
	case "mile", "miles":
		return Mile(val), nil
	case "km":
		return Km(val), nil
	case "ft", "'":
		return Metre(val) * FootInMetres, nil
	case "in", "inch", "\"":
		return Metre(val) * InchInMetres, nil
	default:
		return Metre(0), ErrUnknownUnit
	}
}


// Metres
func (m Metre) Name() string {
	return "metres"
}

func (m Metre) Short() string {
	return "m"
}

func (m Metre) String() string {
	return fmt.Sprintf("%f m", m)
}

func (m Metre) Valid() bool {
	return !math.IsNaN(float64(m))
}

func (m Metre) Metre() Metre {
	return m
}

func (m Metre) Meter() Meter {
	return Meter(m)
}

func (m Metre) Km() Km {
	return Km(m / 1000)
}

func (m Metre) NM() NM {
	return NM(m / NauticalMileInMetres)
}

func (m Metre) Mile() Mile {
	return Mile(m / MileInMetres)
}


// Meters (same with US spelling)
func (m Meter) Name() string {
	return "meters"
}

func (m Meter) Short() string {
	return "m"
}

func (m Meter) String() string {
	return fmt.Sprintf("%f m", m)
}

func (m Meter) Valid() bool {
	return !math.IsNaN(float64(m))
}

func (m Meter) Metre() Metre {
	return Metre(m)
}

func (m Meter) Meter() Meter {
	return m
}

func (m Meter) Km() Km {
	return Km(m / 1000)
}

func (m Meter) NM() NM {
	return NM(Metre(m) / NauticalMileInMetres)
}

func (m Meter) Mile() Mile {
	return Mile(Metre(m) / MileInMetres)
}


// Kilometres
func (k Km) Name() string {
	return "kilometres"
}

func (k Km) Short() string {
	return "km"
}

func (k Km) String() string {
	return fmt.Sprintf("%f km", k)
}

func (k Km) Valid() bool {
	return !math.IsNaN(float64(k))
}

func (k Km) Metre() Metre {
	return Metre(k * 1000)
}

func (k Km) Meter() Meter {
	return Meter(k * 1000)
}

func (k Km) Km() Km {
	return k
}

func (k Km) NM() NM {
	return NM(Metre(k) * 1000 / NauticalMileInMetres)
}

func (k Km) Mile() Mile {
	return Mile(Metre(k) * 1000 / MileInMetres)
}


// Nautical Miles
func (m NM) Name() string {
	return "nautical miles"
}

func (m NM) Short() string {
	return "NM"
}

func (m NM) String() string {
	return fmt.Sprintf("%f NM", m)
}

func (m NM) Valid() bool {
	return !math.IsNaN(float64(m))
}

func (m NM) Metre() Metre {
	return Metre(Metre(m) * NauticalMileInMetres)
}

func (m NM) Meter() Meter {
	return Meter(Metre(m) * NauticalMileInMetres)
}

func (m NM) Km() Km {
	return Km(Metre(m) * NauticalMileInMetres / 1000)
}

func (m NM) NM() NM {
	return m
}

func (m NM) Mile() Mile {
	return Mile(Metre(m) * NauticalMileInMetres / MileInMetres)
}


// Miles
func (m Mile) Name() string {
	return "miles"
}

func (m Mile) Short() string {
	return "mi"
}

func (m Mile) String() string {
	return fmt.Sprintf("%f mi", m)
}

func (m Mile) Valid() bool {
	return !math.IsNaN(float64(m))
}

func (m Mile) Metre() Metre {
	return Metre(Metre(m) * MileInMetres)
}

func (m Mile) Meter() Meter {
	return Meter(Metre(m) * MileInMetres)
}

func (m Mile) Km() Km {
	return Km(Metre(m) * MileInMetres / 1000)
}

func (m Mile) NM() NM {
	return NM(Metre(m) * MileInMetres / NauticalMileInMetres)
}

func (m Mile) Mile() Mile {
	return m
}
