package main

import (
	"fmt"
	"math"
)

type (
	// Feet is ft
	Feet float64
	// Centimeter is cm
	Centimeter float64
	// Minute is min
	Minute float64
	// Second is sec
	Second float64
	// Millisecond is ms
	Millisecond float64
	// Fahrenheit is F
	Fahrenheit float64
	// Celsius is C
	Celsius float64
	// Radian is rad
	Radian float64
	// Degree is degree
	Degree float64
	// Kilogram is kg
	Kilogram float64
	// Pound is lbs
	Pound float64
)

// Converter is a bi-directional struct
type Converter struct {
	centimeter  Centimeter
	feet        Feet
	second      Second
	minute      Minute
	millisecond Millisecond
	fahrenheit  Fahrenheit
	celsius     Celsius
	degree      Degree
	radian      Radian
	pounds      Pound
	kilogram    Kilogram
}

// SecondToMillisecond converts seconds to minute
func (cvr Converter) SecondToMillisecond(s Second) Millisecond {
	return Millisecond(s * 1000)
}

// MinuteToSecond converts minutes to seconds
func (cvr Converter) MinuteToSecond(m Minute) Second {
	return Second(m * 60)
}

// CentimeterToFeet converts centimeter to feet
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(c / 30.48)
}

// FeetToCentimeter converts feet to centimeter
func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	return Centimeter(f * 30.48)
}

// KilogramToPound converts kilogram to pounds
func (cvr Converter) KilogramToPound(k Kilogram) Pound {
	return Pound(k * 2.205)
}

// PoundToKilogram converts pounds to kilogram
func (cvr Converter) PoundToKilogram(p Pound) Kilogram {
	return Kilogram(p / 2.205)
}

// FahrenheitToCelsius converts fahrenheit to celsius
func (cvr Converter) FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 0.5556)
}

// CelsiusToFahrenheit converts celsius to fahrenheit
func (cvr Converter) CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((c * 1.8) + 32)
}

// DegreeToRadian converts degree to radian
func (cvr Converter) DegreeToRadian(d Degree) Radian {
	return Radian(d * (math.Pi / 180))
}

// RadianToDegree converts radian to degree
func (cvr Converter) RadianToDegree(r Radian) Degree {
	return Degree((r * 180) / math.Pi)
}

func main() {
	cvr := Converter{}
	fmt.Println(cvr.SecondToMillisecond(60), "millisecond(s)")
	fmt.Println(cvr.MinuteToSecond(60), "second(s)")
	fmt.Println(cvr.CentimeterToFeet(60), "ft")
	fmt.Println(cvr.FeetToCentimeter(60), "cm")
	fmt.Println(cvr.KilogramToPound(60), "lbs")
	fmt.Println(cvr.PoundToKilogram(60), "kg")
	fmt.Println(cvr.FahrenheitToCelsius(60), "C")
	fmt.Println(cvr.CelsiusToFahrenheit(60), "F")
	fmt.Println(cvr.DegreeToRadian(60), "radian")
	fmt.Println(cvr.RadianToDegree(60), "degree")
}
