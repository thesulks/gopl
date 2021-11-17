package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	png.Encode(os.Stdout, mandelbrotImage()) // NOTE: ignoring errors
}

func mandelbrotImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

// Is image.RGBA concurrency-safe?
func mandelbrotImageWaitGroup() *image.RGBA {
	var wg sync.WaitGroup

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)
			go func(y float64, px, py int) {
				defer wg.Done()

				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}(y, px, py)
		}
	}

	wg.Wait()
	return img
}

func mandelbrotImageWaitGroupWithToken(size int) *image.RGBA {
	var wg sync.WaitGroup
	tokens := make(chan struct{}, size)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			tokens <- struct{}{}
			wg.Add(1)
			go func(y float64, px, py int) {
				defer func() {
					<-tokens
					wg.Done()
				}()
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}(y, px, py)
		}
	}

	wg.Wait()
	return img
}

func mandelbrotImageLocality(size int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup

	stride := height / size
	for i := 0; i < size; i++ {
		start := i * stride
		end := (i + 1) * stride
		if end > height {
			end = height
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for py := start; py < end; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
		}()
	}

	wg.Wait()
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
