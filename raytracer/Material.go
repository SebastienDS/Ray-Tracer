package raytracer

// Material interface
type Material interface {
	scatter(ray Ray, rec HitRecord, attenuation *Vec3, scattered *Ray) bool
}
