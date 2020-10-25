package main

import (
	"fmt"
	"os"

	"github.com/SebastienDS/Ray-Tracer/raytracer"
)

// DrawImage create a ppm file
func DrawImage(buffer []raytracer.Vec3, WIDTH int, HEIGHT int) {
	f, _ := os.Create("raytracer.ppm")
	defer f.Close()

	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", WIDTH, HEIGHT))

	for _, color := range buffer {
		f.WriteString(fmt.Sprintf("%d %d %d\n", int(color.X), int(color.Y), int(color.Z)))
	}
}

func main() {
	WIDTH := 640
	HEIGHT := 480

	var buffer []raytracer.Vec3

	for j := HEIGHT - 1; j >= 0; j-- {
		for i := 0; i < WIDTH; i++ {
			color := raytracer.Vec3{}
			r := float64(i) / float64((WIDTH - 1))
			g := float64(j) / float64((HEIGHT - 1))
			b := float64(0.25)

			color.X = 255 * r
			color.Y = 255 * g
			color.Z = 255 * b
			buffer = append(buffer, color)
		}
	}

	DrawImage(buffer, WIDTH, HEIGHT)
}
