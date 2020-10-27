package raytracer

import "math"

// Camera (origin, lowerLeftCorner, Horizontal, Vertical)
type Camera struct {
	Origin          Vec3
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
}

// NewCamera return a new Camera
func NewCamera(lookFrom Vec3, lookAt Vec3, vup Vec3, vfov float64, aspectRatio float64) Camera {
	theta := DegreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h
	viewportWidth := aspectRatio * viewportHeight

	w := Sub(lookFrom, lookAt).UnitVector()
	u := Cross(vup, w).UnitVector()
	v := Cross(w, u)

	camera := Camera{}
	camera.Origin = lookFrom
	camera.Horizontal = MulF(u, viewportWidth)
	camera.Vertical = MulF(v, viewportHeight)
	// lowerLeftCorner = origin - horizontal/2 - vertical/2 - w
	camera.LowerLeftCorner = Sub(Sub(Sub(camera.Origin, Div(camera.Horizontal, 2)), Div(camera.Vertical, 2)), w)

	return camera
}

// GetRay return the ray (u, v)
func (c Camera) GetRay(u, v float64) Ray {
	// ray = (origin, lower_left_corner + u*horizontal + v*vertical - origin)
	return NewRay(c.Origin, Sub(Add(c.LowerLeftCorner, Add(MulF(c.Horizontal, u), MulF(c.Vertical, v))), c.Origin))
}
