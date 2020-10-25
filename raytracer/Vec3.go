package raytracer

// Vec3 (x, y, z)
type Vec3 struct {
	x float32
	y float32
	z float32
}

// NewVec3 construct the Vec3
func NewVec3(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}
