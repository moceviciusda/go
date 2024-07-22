package space

type Planet string

type sPlanet struct {
	name      Planet
	orbPeriod float64
}

const secondsInEarthYear = 31557600

var planets = []sPlanet{
	{"Mercury", 0.2408467},
	{"Venus", 0.61519726},
	{"Earth", 1.0},
	{"Mars", 1.8808158},
	{"Jupiter", 11.862615},
	{"Saturn", 29.447498},
	{"Uranus", 84.016846},
	{"Neptune", 164.79132},
}

func Age(seconds float64, planet Planet) float64 {
	for _, p := range planets {
		if p.name == planet {
			return seconds / (p.orbPeriod * secondsInEarthYear)
		}
	}
	return -1
}
