package barnacle

import "math"

const Version = "0.0.0a"

type Point struct {
	X, Y float64
}

// Add : Add two points and return a new point
func (a Point) Add(b Point) Point {
	v := Point{}
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	return v
}

// Sub : Subtract two points and return a new point
func (a Point) Sub(b Point) Point {
	v := Point{}
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	return v
}

// MultiplyByScalar : Mulyiply a point by a scalar and return a new point
func (a Point) MultiplyByScalar(s float64) Point {
	v := Point{}
	v.X = a.X * s
	v.Y = a.Y * s
	return v
}

// Dot : Dot product of two points and returning a new point
func (a Point) Dot(b Point) float64 {
	return a.X*b.X + a.Y*b.Y
}

// Length : Return the length of a point
func (a Point) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// Normalize :
func (a Point) Normalize() Point {
	return a.MultiplyByScalar(1.0 / a.Length())
}
