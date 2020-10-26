package main

import (
	"fmt"
	"math"
	"os"

	"github.com/SebastienDS/Ray-Tracer/raytracer"
)

// DrawImage create a ppm file
func DrawImage(buffer []raytracer.Vec3, WIDTH int, HEIGHT int) {
	f, _ := os.Create("raytracer.ppm")
	defer f.Close()

	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", WIDTH, HEIGHT))

	for _, color := range buffer {
		f.WriteString(fmt.Sprintf("%d %d %d\n", int(255*color.X), int(255*color.Y), int(255*color.Z)))
	}
}

// RayColor return the color of the background
func RayColor(ray raytracer.Ray) raytracer.Vec3 {
	t := HitSphere(raytracer.NewVec3(0, 0, -1), 0.5, ray)
	if t > 0 {
		n := raytracer.Neg(ray.At(t), raytracer.NewVec3(0, 0, -1)).UnitVector()
		return raytracer.MulF(raytracer.NewVec3(n.X+1, n.Y+1, n.Z+1), 0.5)
	}
	unitDirection := ray.Direction.UnitVector()
	t = 0.5 * (unitDirection.Y + 1.0)
	return raytracer.Add(
		raytracer.MulF(raytracer.NewVec3(1.0, 1.0, 1.0), (1.0-t)),
		raytracer.MulF(raytracer.NewVec3(0.5, 0.7, 1.0), t))
}

// HitSphere return if ray hit the sphere
func HitSphere(center raytracer.Vec3, radius float64, ray raytracer.Ray) float64 {
	oc := raytracer.Neg(ray.Origin, center)
	a := raytracer.Dot(ray.Direction, ray.Direction)
	b := raytracer.Dot(oc, ray.Direction) * 2
	c := raytracer.Dot(oc, oc) - radius*radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1
	}
	return (-b - math.Sqrt(discriminant)) / (2 * a)

}

func main() {
	// Image
	aspectRatio := 16.0 / 9.0
	WIDTH := 400
	HEIGHT := int(float64(WIDTH) / aspectRatio)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := raytracer.NewVec3(0, 0, 0)
	horizontal := raytracer.NewVec3(viewportWidth, 0, 0)
	vertical := raytracer.NewVec3(0, viewportHeight, 0)
	// lowerLeftCorner = origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length);
	lowerLeftCorner := raytracer.Neg(raytracer.Neg(raytracer.Neg(origin, raytracer.Div(horizontal, 2)), raytracer.Div(vertical, 2)), raytracer.NewVec3(0, 0, focalLength))

	var buffer []raytracer.Vec3

	for j := HEIGHT - 1; j >= 0; j-- {
		for i := 0; i < WIDTH; i++ {
			u := float64(i) / float64((WIDTH - 1))
			v := float64(j) / float64((HEIGHT - 1))

			// ray = (origin, lower_left_corner + u*horizontal + v*vertical - origin)
			ray := raytracer.NewRay(origin, raytracer.Neg(
				raytracer.Add(lowerLeftCorner,
					raytracer.Add(raytracer.MulF(horizontal, u), raytracer.MulF(vertical, v))),
				origin))

			pixelColor := RayColor(ray)
			buffer = append(buffer, pixelColor)
		}
	}

	DrawImage(buffer, WIDTH, HEIGHT)
}
