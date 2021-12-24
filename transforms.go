package barnacle

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// Rotate : Rotate a point by angle radians.
// Best if translated to origin.
func Rotate(p Point, angle float64) Point {

	acos := math.Cos(angle)
	asin := math.Sin(angle)
	rm := mat.NewDense(2, 2, []float64{
		acos, -asin,
		asin, acos,
	})

	newpoint := Point{}

	tmpm := mat.NewDense(1, 2, []float64{p.X, p.Y})
	var w mat.Dense
	w.Mul(tmpm, rm)

	newpoint.X = w.At(0, 0)
	newpoint.Y = w.At(0, 1)

	return newpoint

}

// Scale : Scale a simple point by scale factor.
func Scale(p Point, scalefactor float64) Point {
	temppoint := Point{}

	temppoint.X = p.X * scalefactor
	temppoint.Y = p.Y * scalefactor

	return temppoint
}

// Translate : Translate a simple point by tx,ty
func Translate(p Point, tx float64, ty float64) Point {
	temppoint := Point{}

	tmpmatrix := mat.NewDense(1, 2, []float64{p.X, p.Y})

	f := make([]float64, 2)
	f[0] = -tx
	f[1] = -ty

	b := mat.NewDense(1, 2, f)

	var c mat.Dense
	c.Add(tmpmatrix, b)

	temppoint.X = c.At(0, 0)
	temppoint.Y = c.At(0, 1)

	return temppoint
}
