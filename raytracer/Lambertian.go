package raytracer

// Lambertian (albedo)
type Lambertian struct {
	Albedo Vec3
}

// NewLambertian return a new Lambertian
func NewLambertian(albedo Vec3) Lambertian {
	return Lambertian{albedo}
}

func (l Lambertian) scatter(ray Ray, rec HitRecord, attenuation *Vec3, scattered *Ray) bool {
	scatterDirection := Add(rec.Normal, RandomUnitVector())
	*scattered = NewRay(rec.P, scatterDirection)
	*attenuation = l.Albedo
	return true
}
