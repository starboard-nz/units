Standard units for Go.

## Supported unit

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


### Speed:

```
import (
        "github.com/starboard-nz/units"
)

...

s := 15 * units.Kn
kms := d.Kphs()

```
