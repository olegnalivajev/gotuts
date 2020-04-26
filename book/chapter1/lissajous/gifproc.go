package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.RGBA{0, 128, 0, 1},
	color.RGBA{255,0,0,1}, color.RGBA{255,255,0,1}}

func Lissajous(out io.Writer, c, s, d int) {
	cycles := c // number of complete x oscillator revolutions
	size := s // image canvas covers [-size..+size]
	delay := d // delay between frames in 10ms units
	const (
		res = 0.001 // angular resolution
		nframes = 64 // number of animation frames
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			z := rand.Intn(3)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(z+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
