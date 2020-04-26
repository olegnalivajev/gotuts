package conv

// temperature conversions
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f-32) * 5/9) }
func KToF(k Kelvin) Fahrenheit  { return Fahrenheit(k*9/5 + 32) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin((f-32) * 5/9) }

// distance conversions

func CmToM(cm Centimeter) Meter { return Meter(cm/100) }
func CmToKm(cm Centimeter) Kilometer { return  Kilometer(cm/100000) }
func MToCm(m Meter) Centimeter { return Centimeter(m*100) }
func MToKm(m Meter) Kilometer { return Kilometer(m/1000) }
func KmToCm(km Kilometer) Centimeter { return Centimeter(km*100000) }
func KmToM(km Kilometer) Meter { return Meter(km*100) }
func KmToMile(km Kilometer) Mile { return Mile(km*1.6) }
func MileToKm(mi Mile) Kilometer { return Kilometer(mi/1.6) }

// weight conversions
func GToKg(g Gram) Kilogram { return Kilogram(g/1000) }
func KgToG(kg Kilogram) Gram { return Gram(kg*1000) }
func KgToTon(kg Kilogram) Ton { return Ton(kg/1000) }
func TonToKg(t Ton) Kilogram { return Kilogram(t*1000) }