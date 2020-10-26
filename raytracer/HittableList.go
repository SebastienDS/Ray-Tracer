package raytracer

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
