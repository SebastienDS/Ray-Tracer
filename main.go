package main

import (
	"github.com/SebastienDS/Ray-Tracer/raytracer"
)

func main() {
	// Image
	aspectRatio := 16.0 / 9.0
	WIDTH := 700
	HEIGHT := int(float64(WIDTH) / aspectRatio)
	samplesPerPixel := 100
	maxDepth := 50

	// World
	world := raytracer.NewHittableList()

	materialGround := raytracer.NewLambertian(raytracer.NewVec3(0.8, 0.8, 0))
	materialCenter := raytracer.NewLambertian(raytracer.NewVec3(0.7, 0.3, 0.3))
	materialLeft := raytracer.NewMetal(raytracer.NewVec3(0.8, 0.8, 0.8))
	materialRight := raytracer.NewMetal(raytracer.NewVec3(0.8, 0.6, 0.2))

	world.Add(raytracer.NewSphere(raytracer.NewVec3(0, -100.5, -1), 100, materialGround))
	world.Add(raytracer.NewSphere(raytracer.NewVec3(0, 0, -1), 0.5, materialCenter))
	world.Add(raytracer.NewSphere(raytracer.NewVec3(-1, 0, -1), 0.5, materialLeft))
	world.Add(raytracer.NewSphere(raytracer.NewVec3(1, 0, -1), 0.5, materialRight))

	// Camera
	camera := raytracer.NewCamera(aspectRatio)

	// Render
	var buffer []raytracer.Vec3

	for j := HEIGHT - 1; j >= 0; j-- {
		for i := 0; i < WIDTH; i++ {
			pixelColor := raytracer.NewVec3(0, 0, 0)

			for s := 0; s < samplesPerPixel; s++ {
				u := float64(i) / float64((WIDTH - 1))
				v := float64(j) / float64((HEIGHT - 1))

				ray := camera.GetRay(u, v)
				pixelColor.Add(raytracer.RayColor(ray, world, maxDepth))
			}
			raytracer.ConvertColor(&pixelColor, samplesPerPixel)
			buffer = append(buffer, pixelColor)
		}
	}

	raytracer.DrawImage(buffer, WIDTH, HEIGHT)
}
