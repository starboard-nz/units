package units_test

/**
 * Copyright (c) 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

import (
	"fmt"
	"testing"
	"math/rand"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starboard-nz/units"
)

const δ = 0.000001

func TestDistance(t *testing.T) {

	t.Run("NauticalMiles", func(t *testing.T) {
		d := units.Metre(10000)
		assert.InDelta(t, 5.399568034557236, float64(d.NM()), δ)
	})

	t.Run("Miles", func(t *testing.T) {
		d := units.NM(5.399568034557236)
		assert.InDelta(t, 6.2137119223733395, float64(d.Mile()), δ)
	})

	t.Run("Kms", func(t *testing.T) {
		d := units.Km(235)
		assert.InDelta(t, 146.0222301757, float64(d.Mile()), δ)
	})
}

func TestParseDistance(t *testing.T) {
	var testData = map[string]units.Distance{
		"-10.4m":   units.Metre(-10.4),
		"32 inch":  units.Metre(32 * units.InchInMetres),
		"120km":    units.Km(120),
		"23.6mile": units.Mile(23.6),
		"-17.3\"":  units.Metre(-17.3 * units.InchInMetres),
		".5'":      units.Metre(0.5 * units.FootInMetres),
		"79.3NM":   units.NM(79.3),
	}

	for str, exp := range testData {
		d, err := units.ParseDistance(str)
		require.NoError(t, err)
		assert.InDelta(t, float64(exp.Metre()), float64(d.Metre()), δ, str)
		assert.InDelta(t, float64(exp.Meter()), float64(d.Meter()), δ, str)
		assert.InDelta(t, float64(exp.Km()), float64(d.Km()), δ, str)
		assert.InDelta(t, float64(exp.NM()), float64(d.NM()), δ, str)
		assert.InDelta(t, float64(exp.Mile()), float64(d.Mile()), δ, str)
		tExp := fmt.Sprintf("%T", exp)
		tUnit := fmt.Sprintf("%T", d)
		assert.Equal(t, tExp, tUnit)
	}

	var errTests = []string{"hello 6' world", "0.1.2m", "--123NM", "42"}

	for _, str := range errTests {
		_, err := units.ParseDistance(str)
		assert.Error(t, err)
	}
}

func randomDistanceConversion(d units.Distance) units.Distance {
	u := rand.Intn(5)

	switch u {
	case 0:
		return d.Metre()
	case 1:
		return d.Meter()
	case 2:
		return d.Km()
	case 3:
		return d.NM()
	case 4:
		return d.Mile()
	}

	return d.Metre()
}


func randomDistance() units.Distance {
	val := rand.Float64()*200000 - 100000 // nolint:gosec

	u := rand.Intn(5)

	switch u {
	case 0:
		return units.Metre(val)	
	case 1:
		return units.Meter(val)
	case 2:
		return units.Km(val)
	case 3:
		return units.NM(val)
	case 4:
		return units.Mile(val)
	}

	return units.Metre(0)
}

func TestDistanceRandom(t *testing.T) {
	const N = 10000

	for i := 0; i < N; i++ {
		d0 := randomDistance()
		d := d0

		// do 5 random conversions
		for j := 0; j < 5; j++ {
			d = randomDistanceConversion(d)
		}

		assert.InDelta(t, float64(d0.Metre()), float64(d.Metre()), δ)
		assert.InDelta(t, float64(d0.Meter()), float64(d.Meter()), δ)
		assert.InDelta(t, float64(d0.Km()), float64(d.Km()), δ)
		assert.InDelta(t, float64(d0.NM()), float64(d.NM()), δ)
		assert.InDelta(t, float64(d0.Mile()), float64(d.Mile()), δ)		
	}
}
