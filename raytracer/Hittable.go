package raytracer

// Hittable interface
type Hittable interface {
	Hit(ray Ray, tMin float64, tMax float64, rec *HitRecord) bool
}
