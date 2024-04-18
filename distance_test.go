package units_test

/**
 * Copyright (c) 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/starboard-nz/units"
)

func TestDistance(t *testing.T) {
	const ε = 0.000000001

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
