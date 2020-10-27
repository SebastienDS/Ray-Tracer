package raytracer

import "math"

// Camera (origin, lowerLeftCorner, Horizontal, Vertical)
type Camera struct {
	Origin          Vec3
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
	U               Vec3
	V               Vec3
	W               Vec3
	LensRadius      float64
}

// NewCamera return a new Camera
func NewCamera(lookFrom Vec3, lookAt Vec3, vup Vec3, vfov float64, aspectRatio float64, aperture float64, focusDist float64) Camera {
	theta := DegreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h
	viewportWidth := aspectRatio * viewportHeight

	camera := Camera{}

	camera.W = Sub(lookFrom, lookAt).UnitVector()
	camera.U = Cross(vup, camera.W).UnitVector()
	camera.V = Cross(camera.W, camera.U)
	camera.Origin = lookFrom
	camera.Horizontal = MulF(MulF(camera.U, viewportWidth), focusDist)
	camera.Vertical = MulF(MulF(camera.V, viewportHeight), focusDist)
	// lowerLeftCorner = origin - horizontal/2 - vertical/2 - focusDist*w
	camera.LowerLeftCorner = Sub(Sub(Sub(camera.Origin, Div(camera.Horizontal, 2)), Div(camera.Vertical, 2)), MulF(camera.W, focusDist))
	camera.LensRadius = aperture / 2

	return camera
}

// GetRay return the ray (u, v)
func (c Camera) GetRay(u, v float64) Ray {
	rd := MulF(RandomInUnitDisk(), c.LensRadius)
	offset := Add(MulF(c.U, rd.X), MulF(c.V, rd.Y))

	// ray = (origin + offset, lower_left_corner + u*horizontal + v*vertical - origin - offset)
	return NewRay(Add(c.Origin, offset), Sub(Sub(Add(c.LowerLeftCorner, Add(MulF(c.Horizontal, u), MulF(c.Vertical, v))), c.Origin), offset))
}
