// Package triangle from go track in exercism
package triangle

import "math"

// Kind type tells the type of triangle.
type Kind int

// Constant types consumed by test cases
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides function takes three sides of triangle and return triangle kind, or if it is invalid.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !isTriangle(a, b, c) {
		k = NaT
	} else if isEquilateralTriangle(a, b, c) {
		k = Equ
	} else if isIsoscelesTriangle(a, b, c) {
		k = Iso
	} else {
		k = Sca
	}
	return k
}

// isIsoscelesTriangle function takes three sides of a triangle and returns a boolean if its Isosceles or not
func isIsoscelesTriangle(a, b, c float64) bool {
	return a == b || b == c || c == a
}

// isEquilateralTriangle function takes three sides of a triangle and returns a boolean if its Equilateral or not
func isEquilateralTriangle(a, b, c float64) bool {
	return a == b && b == c
}

// isTriangle function takes three lengths and returns boolean if its a triangle or not
func isTriangle(a, b, c float64) bool {
	InvalidSides := invalidLength(a) || invalidLength(b) || invalidLength(c)
	IsGreaterThanZero := a > 0 && b > 0 && c > 0

	return IsGreaterThanZero && triangleInqualityProp(a, b, c) && !InvalidSides
}

// triangleInqualityProp function takes three sides of a triangle and returns boolean if it satisfies inquality prop or not
func triangleInqualityProp(a, b, c float64) bool {
	return (a+b) >= c && (b+c) >= a && (a+c) >= b
}

// invalidLength function takes a leangth and return boolean if its valid or not
func invalidLength(a float64) bool {
	return math.IsNaN(a) || math.IsInf(a, 0)
}
