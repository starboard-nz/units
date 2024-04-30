package units

/**
 * Copyright (c) 2024, Starboard Maritime Intelligence
 * All rights reserved. Use is subject to License terms.
 * See LICENSE in the root directory of this source tree.
 */

type Unit interface {
	// Returns abbreviation of the unit.
	Short() string
	// Returns full name of the unit.
	Name() string
	// Returns true if the value is valid.
	Valid() bool
	// String returns a formatted string of the quantity and unit, e.g. "10 km/h"
	String() string
}
