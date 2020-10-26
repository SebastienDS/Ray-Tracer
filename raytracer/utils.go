package raytracer

import (
	"fmt"
	"math"
	"math/rand"
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

// ConvertColor _
func ConvertColor(pixelColor *Vec3, samplesPerPixel int) {
	scale := 1.0 / float64(samplesPerPixel)
	pixelColor.X = Clamp(pixelColor.X*scale, 0, 1)
	pixelColor.Y = Clamp(pixelColor.Y*scale, 0, 1)
	pixelColor.Z = Clamp(pixelColor.Z*scale, 0, 1)
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

// RandomFloat return a random number in [min, max)
func RandomFloat(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

// Clamp the value x to the range [min, max]
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
