package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"

	"github.com/SebastienDS/Ray-Tracer/raytracer"
)

func main() {
	// Image
	aspectRatio := 3.0 / 2.0
	WIDTH := 1200
	HEIGHT := int(float64(WIDTH) / aspectRatio)
	samplesPerPixel := 500
	maxDepth := 50

	// World
	world := raytracer.RandomScene()

	materialGround := raytracer.NewLambertian(raytracer.NewVec3(0.8, 0.8, 0))
	materialCenter := raytracer.NewLambertian(raytracer.NewVec3(0.1, 0.2, 0.5))
	materialLeft := raytracer.NewDielectric(1.5)
	materialRight := raytracer.NewMetal(raytracer.NewVec3(0.8, 0.6, 0.2), 0)

	world.Add(raytracer.NewSphere(raytracer.NewVec3(0, -100.5, -1), 100, materialGround))
	world.Add(raytracer.NewSphere(raytracer.NewVec3(0, 0, -1), 0.5, materialCenter))
	world.Add(raytracer.NewSphere(raytracer.NewVec3(-1, 0, -1), 0.5, materialLeft))
	world.Add(raytracer.NewSphere(raytracer.NewVec3(-1, 0, -1), -0.45, materialLeft))
	world.Add(raytracer.NewSphere(raytracer.NewVec3(1, 0, -1), 0.5, materialRight))

	// Camera
	lookFrom := raytracer.NewVec3(13, 2, 3)
	lookAt := raytracer.NewVec3(0, 0, 0)
	vup := raytracer.NewVec3(0, 1, 0)
	distToFocus := 10.0
	aperture := 0.1

	camera := raytracer.NewCamera(lookFrom, lookAt, vup, 20, aspectRatio, aperture, distToFocus)

	// Render
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	var wg sync.WaitGroup

	for j := 0; j < HEIGHT; j++ {
		wg.Add(1)

		go func(j int) {
			for i := 0; i < WIDTH; i++ {
				pixelColor := raytracer.NewVec3(0, 0, 0)

				for s := 0; s < samplesPerPixel; s++ {
					u := float64(i) / float64(WIDTH-1)
					v := float64(j) / float64(HEIGHT-1)

					ray := camera.GetRay(u, v)
					pixelColor.Add(raytracer.RayColor(ray, world, maxDepth))
				}
				raytracer.ConvertColor(&pixelColor, samplesPerPixel)
				img.Set(i, HEIGHT-j, color.RGBA{uint8(255 * pixelColor.X), uint8(255 * pixelColor.Y), uint8(255 * pixelColor.Z), 0xff})

				wg.Done()
			}
		}(j)
	}

	wg.Wait()

	f, _ := os.OpenFile("raytracer2.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
