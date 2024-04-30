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

d := units.NM(10)
kms := d.Km()

```

Using the parser:

```
d, err := units.ParseDistance("3000NM")
if err != nil {
       return err
}
```

Casting a float to a distance in some unit:

```
d0 := float64(15.3) // let's say this is nautical miles
d := units.NM(d0)
```

Print a distance:

d := units.Km(25)
fmt.Printf("%v", d) // prints "25.00 km"
fmt.Printf("%v", d.Metre()) // prints "25000.00 km"
fmt.Printf("%.0f %s", d, d.Short()) // prints "25 km"

### Speed:

```
import (
        "github.com/starboard-nz/units"
)

...

s := units.Knot(15)
kphs := s.Kph()

```

Casting a float to a Speed unit:

```
s0 := float64(15.3) // assuming this is in knots
s := units.Knot(d0)
```

Using the parser:

```
s, err := units.ParseSpeed("10 km/h")
if err != nil {
       return err
}
```

Print a speed:

s := units.Mps(25.1234)
fmt.Printf("%v", s) // prints "25.12 m/s"
fmt.Printf("%.0f %s", s, s.Short()) // prints "25 m/s"
