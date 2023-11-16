package convpkg

func CtoF(temp Celsius) Fahrenheit {
	fTemp := (temp * 9 / 5) + 32
	return Fahrenheit(fTemp)
}

func FtoC(temp Fahrenheit) Celsius {
	cTemp := (temp - 32) * 5 / 9
	return Celsius(cTemp)
}
