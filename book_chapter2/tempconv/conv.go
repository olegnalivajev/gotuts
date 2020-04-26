package tempconv

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celcius { return Celcius((f-32) * 5/9) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*9/5 + 32) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin((f-32) * 5/9) }
