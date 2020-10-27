package raytracer

import (
	"math"
	"math/rand"
)

// Dielectric (RefractionIndex)
type Dielectric struct {
	RefractionIndex float64
}

// NewDielectric return a new Dielectric
func NewDielectric(refractionIndex float64) Dielectric {
	return Dielectric{refractionIndex}
}

func (d Dielectric) scatter(ray Ray, rec HitRecord, attenuation *Vec3, scattered *Ray) bool {
	*attenuation = NewVec3(1, 1, 1)
	var refractionRatio float64
	if rec.FrontFace {
		refractionRatio = 1 / d.RefractionIndex
	} else {
		refractionRatio = d.RefractionIndex
	}

	unitDirection := ray.Direction.UnitVector()
	cosTheta := math.Min(Dot(unitDirection.Neg(), rec.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1
	var direction Vec3

	if cannotRefract || (reflectance(cosTheta, refractionRatio) > rand.Float64()) {
		direction = Reflect(unitDirection, rec.Normal)
	} else {
		direction = Refract(unitDirection, rec.Normal, refractionRatio)
	}

	*scattered = NewRay(rec.P, direction)
	return true
}

func reflectance(cosine, refIdx float64) float64 {
	// Use Schlick's approximation for reflectance
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
