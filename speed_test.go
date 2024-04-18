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

func TestSpeed(t *testing.T) {
	const ε = 0.000000001

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
