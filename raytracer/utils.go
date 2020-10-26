package raytracer

import (
	"fmt"
	"math"
	"os"
)

// DrawImage create a ppm file
func DrawImage(buffer []Vec3, WIDTH int, HEIGHT int) {
	f, _ := os.Create("raytracer.ppm")
	defer f.Close()

	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", WIDTH, HEIGHT))

	for _, color := range buffer {
		f.WriteString(fmt.Sprintf("%d %d %d\n", int(255*color.X), int(255*color.Y), int(255*color.Z)))
	}
}

// RayColor return the color of the background
func RayColor(ray Ray, world Hittable) Vec3 {
	rec := HitRecord{}

	if world.Hit(ray, 0, math.MaxFloat64, &rec) {
		return MulF(Add(rec.Normal, NewVec3(1, 1, 1)), 0.5)
	}

	unitDirection := ray.Direction.UnitVector()
	t := (unitDirection.Y + 1.0) * 0.5
	return Add(
		MulF(NewVec3(1.0, 1.0, 1.0), (1.0-t)),
		MulF(NewVec3(0.5, 0.7, 1.0), t))
}

// DegreesToRadians convert degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return (degrees * math.Pi) / 180.0
}
