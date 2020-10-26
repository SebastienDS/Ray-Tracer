package raytracer

// Metal (Albedo)
type Metal struct {
	Albedo Vec3
}

// NewMetal return a new Metal
func NewMetal(albedo Vec3) Metal {
	return Metal{albedo}
}

func (m Metal) scatter(ray Ray, rec HitRecord, attenuation *Vec3, scattered *Ray) bool {
	reflected := Reflect(ray.Direction.UnitVector(), rec.Normal)
	*scattered = NewRay(rec.P, reflected)
	*attenuation = m.Albedo
	return Dot(scattered.Direction, rec.Normal) > 0
}
