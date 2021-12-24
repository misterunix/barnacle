package barnacle

type Triangle struct {
	Points [3]Point
	Center Point
}

func (t Triangle) CalcCenter() Point {

	p := Point{}

	tx := 0.0
	ty := 0.0
	tempcount := len(t.Points)
	for i := 0; i < tempcount; i++ {
		tx += t.Points[i].X
		ty += t.Points[i].Y
	}
	tx /= float64(tempcount)
	ty /= float64(tempcount)

	p.X = tx
	p.Y = ty
	return p
}

// Translate : Translate a triangle t by tx,ty referenced from the center.
// Center is assumed to be already calculated.
func (t Triangle) Translate(tx, ty float64) Triangle {

	temptriangle := Triangle{}
	count := len(t.Points)

	for i := 0; i < count; i++ {
		//temptriangle.Points[i] = Translate(t.Points[i], t.Center.X, t.Center.Y)
		temptriangle.Points[i] = Translate(t.Points[i], tx, ty)
	}

	return temptriangle
}

// Rotate : Rotate a triangle around its assumed centerpoint
// Brst if t has already been translated to the origin.
// Calculates a new center.
func (t Triangle) Rotate(angle float64) Triangle {

	temptriangle := Triangle{}
	for i := 0; i < 3; i++ {
		temppoint := Rotate(t.Points[i], angle)
		temptriangle.Points[i].X = temppoint.X
		temptriangle.Points[i].Y = temppoint.Y
	}
	temptriangle.Center = temptriangle.CalcCenter()
	return temptriangle
}

func (t Triangle) Scale(factor float64) Triangle {

	temptriangle := Triangle{}
	count := len(t.Points)
	for i := 0; i < count; i++ {
		temppoint := Scale(t.Points[i], factor)
		temptriangle.Points[i].X = temppoint.X
		temptriangle.Points[i].Y = temppoint.Y
	}
	return temptriangle
}
