package units_test

/**
 * Copyright (c) 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starboard-nz/units"
)

const ε = 0.000000001

func TestDistance(t *testing.T) {

	t.Run("NauticalMiles", func(t *testing.T) {
		d := 10000 * units.Metre
		assert.InEpsilon(t, 5.399568034557236, d.NauticalMiles(), ε)
	})

	t.Run("Miles", func(t *testing.T) {
		d := 5.399568034557236 * units.NauticalMile
		assert.InEpsilon(t, 6.2137119223733395, d.Miles(), ε)
	})

	t.Run("Miles", func(t *testing.T) {
		d := 3 * units.Foot
		assert.InEpsilon(t, 0.9144, d.Metres(), ε)
	})
}

func TestParseDistance(t *testing.T) {
	var testData = map[string]units.Distance{
		"-10.4m":   -10.4 * units.Metre,
		"32 inch":  32 * units.Inch,
		"120km":    120 * units.Kilometre,
		"23.6mile": 23.6 * units.Mile,
		"-17.3\"":  -17.3 * units.Inch,
		".5'":      0.5 * units.Foot,
		"79.3NM":   79.3 * units.NauticalMile,
	}

	for str, exp := range testData {
		d, err := units.ParseDistance(str)
		require.NoError(t, err)
		assert.InEpsilon(t, exp.Metres(), d.Metres(), ε)
	}

	var errTests = []string{"hello 6' world", "0.1.2m", "--123NM", "42"}

	for _, str := range errTests {
		_, err := units.ParseDistance(str)
		assert.Error(t, err)
	}
}
