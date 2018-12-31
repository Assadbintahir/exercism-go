// space package is an exercise from exercism go track
package space

// type Planet used in test cases
type Planet string

// constant values for calculations
const earthOrbitalPeriod         float64 = 31557600
// planet orbital period in earth years
var planetOrbitalInEarthYears = map[Planet] float64 {
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// function age returns age in years upon receiving age in seconds
func Age(inSeconds float64, planet Planet) (inYears float64) {
	return inSeconds / planetOrbitalPeriod(planetOrbitalInEarthYears[planet])
}

// converts orbital period of planets from earth years to respective planet years
func planetOrbitalPeriod(periodInEarthYears float64) (orbitalPeriodInPlanetYears float64) {
	return periodInEarthYears * earthOrbitalPeriod
}
