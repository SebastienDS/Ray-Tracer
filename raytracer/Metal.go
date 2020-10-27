package raytracer

// Metal (Albedo)
type Metal struct {
	Albedo Vec3
	Fuzz   float64
}

// NewMetal return a new Metal
func NewMetal(albedo Vec3, fuzz float64) Metal {
	if fuzz < 1 {
		return Metal{albedo, fuzz}
	}
	return Metal{albedo, 1}
}

func (m Metal) scatter(ray Ray, rec HitRecord, attenuation *Vec3, scattered *Ray) bool {
	reflected := Reflect(ray.Direction.UnitVector(), rec.Normal)
	*scattered = NewRay(rec.P, Add(reflected, MulF(RandomInUnitSphere(), m.Fuzz)))
	*attenuation = m.Albedo
	return Dot(scattered.Direction, rec.Normal) > 0
}
