package raytracer

// HitRecord represent an hit
type HitRecord struct {
	P         Vec3
	Normal    Vec3
	T         float64
	FrontFace bool
}

func (rec *HitRecord) setFaceNormal(ray Ray, outwardNormal Vec3) {
	rec.FrontFace = Dot(ray.Direction, outwardNormal) < 0
	if rec.FrontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.Neg()
	}
}
