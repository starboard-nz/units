Standard units for Go.

## Supported units

Currently supported:
 * Distance
 * Speed

## How to use

### Distance:

```
import (
        "github.com/starboard-nz/units"
)

...

d := 10 * units.NauticalMile
kms := d.Kilometres()

```

Using the parser:

```
d, err := units.ParseDistance("3000NM")
if err != nil {
       return err
}
```

Using a variable that stores a distance in a specific unit:

```
d0 := float64(15.3) // nautical miles
d := units.Distance(d0) * units.NauticalMile
```

### Speed:

```
import (
        "github.com/starboard-nz/units"
)

...

s := 15 * units.Kn
kms := d.Kphs()

```

Using a variable that stores a speed in a specific unit:

```
s0 := float64(15.3) // knots
s := units.Speed(d0) * units.Knot
```

Using the parser:

```
s, err := units.ParseSpeed("10 km/h")
if err != nil {
       return err
}
```