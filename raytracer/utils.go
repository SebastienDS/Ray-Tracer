package raytracer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

// RenderPPM create a ppm file
func RenderPPM(buffer []Vec3, WIDTH int, HEIGHT int) {
	f, _ := os.Create("raytracer.ppm")
	defer f.Close()

	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", WIDTH, HEIGHT))

	for _, color := range buffer {
		f.WriteString(fmt.Sprintf("%d %d %d\n", int(255*color.X), int(255*color.Y), int(255*color.Z)))
	}
}

// RenderPNG create a ppm file
func RenderPNG(buffer []Vec3, WIDTH int, HEIGHT int) {
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	for index, value := range buffer {
		i := index % WIDTH
		j := index / WIDTH

		img.Set(i, j, color.RGBA{uint8(255 * value.X), uint8(255 * value.Y), uint8(255 * value.Z), 0xff})
	}

	f, _ := os.OpenFile("raytracer.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

// ConvertColor in [0, 1]
func ConvertColor(pixelColor *Vec3, samplesPerPixel int) {
	scale := 1.0 / float64(samplesPerPixel)
	pixelColor.X = Clamp(math.Sqrt(pixelColor.X*scale), 0, 1)
	pixelColor.Y = Clamp(math.Sqrt(pixelColor.Y*scale), 0, 1)
	pixelColor.Z = Clamp(math.Sqrt(pixelColor.Z*scale), 0, 1)
}

// RayColor return the color of the background
func RayColor(ray Ray, world Hittable, depth int) Vec3 {
	rec := HitRecord{}

	if depth <= 0 {
		return NewVec3(0, 0, 0)
	}

	if world.Hit(ray, 0.001, math.MaxFloat64, &rec) {
		var scattered Ray
		var attenuation Vec3
		if rec.Material.scatter(ray, rec, &attenuation, &scattered) {
			return Mul(attenuation, RayColor(scattered, world, depth-1))
		}
		return NewVec3(0, 0, 0)
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
