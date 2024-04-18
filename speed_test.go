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

func TestSpeed(t *testing.T) {
	t.Run("Knots", func(t *testing.T) {
		d := 16 * units.Mps
		assert.InEpsilon(t, 31.101511879, d.Kns(), ε)
	})

	t.Run("km/h", func(t *testing.T) {
		d := 2 * units.Mps
		assert.InEpsilon(t, 7.2, d.Kphs(), ε)
	})

	t.Run("m/s", func(t *testing.T) {
		d := 16 * units.Kn
		assert.InEpsilon(t, 8.23111111111, d.Mpss(), ε)
	})
}

func TestParseSpeed(t *testing.T) {
	var testData = map[string]units.Speed{
		"-10.4kn":   -10.4 * units.Knot,
		"32 m/s":  32 * units.Mps,
		"120km/h":    120 * units.Kph,
		"23.6	mph": 23.6 * units.Mph,
		"10kn": 10 * units.Kn, // non-breaking space
	}

	for str, exp := range testData {
		d, err := units.ParseSpeed(str)
		require.NoError(t, err, str)
		assert.InEpsilon(t, exp.Mpss(), d.Mpss(), ε)
	}

	var errTests = []string{"hello 6' world", "0.1.2m/s", "--123kn", "42"}

	for _, str := range errTests {
		_, err := units.ParseSpeed(str)
		assert.Error(t, err)
	}
}
