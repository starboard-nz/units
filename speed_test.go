package units_test

/**
 * Copyright (c) 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starboard-nz/units"
)

func TestSpeed(t *testing.T) {
	t.Run("Knots", func(t *testing.T) {
		d := units.Mps(16)
		assert.InDelta(t, 31.101511879, float64(d.Knot()), δ)
	})

	t.Run("km/h", func(t *testing.T) {
		d := units.Mps(2)
		assert.InDelta(t, 7.2, float64(d.Kph()), δ)
	})

	t.Run("m/s", func(t *testing.T) {
		d := units.Knot(16)
		assert.InDelta(t, 8.23111111111, float64(d.Mps()), δ)
	})
}

func TestParseSpeed(t *testing.T) {
	var testData = map[string]units.Speed{
		"-10.4kn": units.Knot(-10.4),
		"32 m/s":  units.Mps(32),
		"120km/h": units.Kph(120),
		"23.6	mph": units.Mph(23.6),
		"10kn": units.Knot(10),
	}

	for str, exp := range testData {
		s, err := units.ParseSpeed(str)
		require.NoError(t, err, str)
		assert.InDelta(t, float64(exp.Mps()), float64(s.Mps()), δ)
		assert.InDelta(t, float64(exp.Kph()), float64(s.Kph()), δ)
		assert.InDelta(t, float64(exp.Mph()), float64(s.Mph()), δ)
		assert.InDelta(t, float64(exp.Knot()), float64(s.Knot()), δ)
		tExp := fmt.Sprintf("%T", exp)
		tUnit := fmt.Sprintf("%T", s)
		assert.Equal(t, tExp, tUnit)
	}

	var errTests = []string{"hello 6' world", "0.1.2m/s", "--123kn", "42"}

	for _, str := range errTests {
		_, err := units.ParseSpeed(str)
		assert.Error(t, err)
	}
}

func randomSpeedConversion(d units.Speed) units.Speed {
	u := rand.Intn(4)

	switch u {
	case 0:
		return d.Mps()
	case 1:
		return d.Knot()
	case 2:
		return d.Kph()
	case 3:
		return d.Mph()
	}

	return d.Mps()
}


func randomSpeed() units.Speed {
	val := rand.Float64()*200000 - 100000 // nolint:gosec

	u := rand.Intn(4)

	switch u {
	case 0:
		return units.Mps(val)
	case 1:
		return units.Knot(val)
	case 2:
		return units.Kph(val)
	case 3:
		return units.Mph(val)
	}

	return units.Mps(val)
}

func TestSpeedRandom(t *testing.T) {
	const N = 10000

	for i := 0; i < N; i++ {
		s0 := randomSpeed()
		s := s0

		// do 5 random conversions
		for j := 0; j < 5; j++ {
			s = randomSpeedConversion(s)
		}

		assert.InDelta(t, float64(s0.Mps()), float64(s.Mps()), δ)
		assert.InDelta(t, float64(s0.Kph()), float64(s.Kph()), δ)
		assert.InDelta(t, float64(s0.Mph()), float64(s.Mph()), δ)
		assert.InDelta(t, float64(s0.Knot()), float64(s.Knot()), δ)
	}
}
