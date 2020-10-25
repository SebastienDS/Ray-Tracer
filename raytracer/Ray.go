package raytracer

// Ray (origin, direction)
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

// NewRay return a new Ray
func NewRay(origin, direction Vec3) Ray {
	return Ray{origin, direction}
}

// At return the Point at t
func (r Ray) At(t float64) Vec3 {
	return Add(r.Origin, MulF(r.Direction, t))
}
