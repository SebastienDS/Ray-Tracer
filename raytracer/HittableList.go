package raytracer

import "math/rand"

// HittableList contains Hittables objects
type HittableList struct {
	Objects []Hittable
}

// NewHittableList return a new HittableList
func NewHittableList() HittableList {
	return HittableList{}
}

// Add an hittable object to HittableList
func (h *HittableList) Add(object Hittable) {
	h.Objects = append(h.Objects, object)
}

// Clear Objects
func (h *HittableList) Clear() {
	h.Objects = nil
}

// Hit return if an Hittable Object is hitten
func (h HittableList) Hit(ray Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	tempRec := HitRecord{}
	hitAnything := false
	closestSoFar := tMax

	for _, object := range h.Objects {
		if object.Hit(ray, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}
	return hitAnything
}

// RandomScene create a random Scene
func RandomScene() HittableList {
	world := NewHittableList()

	groundMaterial := NewLambertian(NewVec3(0.5, 0.5, 0.5))
	world.Add(NewSphere(NewVec3(0, -1000, 0), 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := NewVec3(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if Sub(center, NewVec3(4, 0.2, 0)).Length() > 0.9 {
				var sphereMaterial Material

				if chooseMat < 0.8 {
					// diffuse
					albedo := Mul(NewVec3Random(), NewVec3Random())
					sphereMaterial = NewLambertian(albedo)
				} else if chooseMat < 0.95 {
					// Metal
					albedo := NewVec3Rand(0.5, 1)
					fuzz := RandomFloat(0, 0.5)
					sphereMaterial = NewMetal(albedo, fuzz)
				} else {
					// glass
					sphereMaterial = NewDielectric(1.50)
				}
				world.Add(NewSphere(center, 0.2, sphereMaterial))

			}
		}
	}

	material1 := NewDielectric(1.5)
	world.Add(NewSphere(NewVec3(0, 1, 0), 1, material1))

	material2 := NewLambertian(NewVec3(0.4, 0.2, 0.1))
	world.Add(NewSphere(NewVec3(-4, 1, 0), 1, material2))

	material3 := NewMetal(NewVec3(0.7, 0.6, 0.5), 0)
	world.Add(NewSphere(NewVec3(4, 1, 0), 1, material3))

	return world
}
