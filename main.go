package main

import (
	"fmt"
	"os"

	"github.com/SebastienDS/Ray-Tracer/raytracer"
)

// DrawImage create a ppm file
func DrawImage(buffer []raytracer.Color, WIDTH int, HEIGHT int) {
	f, _ := os.Create("image.ppm")
	defer f.Close()

	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", WIDTH, HEIGHT))

	for _, color := range buffer {
		f.WriteString(fmt.Sprintf("%d %d %d\n", color.R, color.G, color.B))
	}
}

func main() {
	WIDTH := 640
	HEIGHT := 480

	buffer := make([]raytracer.Color, WIDTH*HEIGHT)
	for j := 0; j < HEIGHT; j++ {
		buffer[j*WIDTH+42].R = 255
	}

	DrawImage(buffer, WIDTH, HEIGHT)

}
