package raytracer

import "math"

// Sphere (center, radius)
type Sphere struct {
	Center   Vec3
	Radius   float64
	Material Material
}

// NewSphere return a new Sphere
func NewSphere(center Vec3, radius float64, material Material) Sphere {
	return Sphere{center, radius, material}
}

func calculHitRecord(rec *HitRecord, s Sphere, ray Ray, temp float64) {
	rec.T = temp
	rec.P = ray.At(rec.T)
	outwardNormal := Div(Sub(rec.P, s.Center), s.Radius)
	rec.setFaceNormal(ray, outwardNormal)
	rec.Material = s.Material
}

// Hit return if sphere is hitten
func (s Sphere) Hit(ray Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	oc := Sub(ray.Origin, s.Center)
	a := ray.Direction.LengthSquared()
	halfB := Dot(oc, ray.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c

	if discriminant > 0 {
		root := math.Sqrt(discriminant)

		temp := (-halfB - root) / a
		if temp < tMax && temp > tMin {
			calculHitRecord(rec, s, ray, temp)
			return true
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			calculHitRecord(rec, s, ray, temp)
			return true
		}
	}
	return false
}
