package conv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Centimeter float64
type Meter float64
type Kilometer float64
type Mile float64

type Gram float64
type Kilogram float64
type Ton float64

const (
	AbsoluteZeroC Celsius = -273.15
	AbsoluteZeroK Kelvin  = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

func (cm Centimeter) String() string { return fmt.Sprintf("%gcm", cm) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (km Kilometer) String() string { return fmt.Sprintf("%gkm", km) }
func (mi Mile) String() string { return fmt.Sprintf("%gmi", mi) }

func (g Gram) String() string { return fmt.Sprintf("%gg", g) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
func (t Ton) String() string { return fmt.Sprintf("%gt", t) }

