package main

import (
	"github.com/SebastienDS/Ray-Tracer/raytracer"
)

func main() {
	// Image
	aspectRatio := 16.0 / 9.0
	WIDTH := 400
	HEIGHT := int(float64(WIDTH) / aspectRatio)

	// World
	world := raytracer.NewHittableList()
	world.Add(raytracer.NewSphere(raytracer.NewVec3(0, 0, -1), 0.5))
	world.Add(raytracer.NewSphere(raytracer.NewVec3(0, -100.5, -1), 100))

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := raytracer.NewVec3(0, 0, 0)
	horizontal := raytracer.NewVec3(viewportWidth, 0, 0)
	vertical := raytracer.NewVec3(0, viewportHeight, 0)
	// lowerLeftCorner = origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length);
	lowerLeftCorner := raytracer.Neg(raytracer.Neg(raytracer.Neg(origin, raytracer.Div(horizontal, 2)), raytracer.Div(vertical, 2)), raytracer.NewVec3(0, 0, focalLength))

	// Render
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

			pixelColor := raytracer.RayColor(ray, world)
			buffer = append(buffer, pixelColor)
		}
	}

	raytracer.DrawImage(buffer, WIDTH, HEIGHT)
}
