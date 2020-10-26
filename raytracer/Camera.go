package raytracer

// Camera (origin, lowerLeftCorner, Horizontal, Vertical)
type Camera struct {
	Origin          Vec3
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
}

// NewCamera return a new Camera
func NewCamera() Camera {
	aspectRatio := 16.0 / 9.0
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	camera := Camera{}
	camera.Origin = NewVec3(0, 0, 0)
	camera.Horizontal = NewVec3(viewportWidth, 0, 0)
	camera.Vertical = NewVec3(0, viewportHeight, 0)
	// lowerLeftCorner = origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length);
	camera.LowerLeftCorner = Neg(Neg(Neg(camera.Origin, Div(camera.Horizontal, 2)), Div(camera.Vertical, 2)), NewVec3(0, 0, focalLength))

	return camera
}

// GetRay return the ray (u, v)
func (c Camera) GetRay(u, v float64) Ray {
	// ray = (origin, lower_left_corner + u*horizontal + v*vertical - origin)
	return NewRay(c.Origin, Neg(Add(c.LowerLeftCorner, Add(MulF(c.Horizontal, u), MulF(c.Vertical, v))), c.Origin))
}
