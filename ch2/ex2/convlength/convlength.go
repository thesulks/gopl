package convlength

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func MToFt(m Meter) Feet {
	return Feet(m * 3.281)
}

func FtToM(f Feet) Meter {
	return Meter(f / 3.281)
}
