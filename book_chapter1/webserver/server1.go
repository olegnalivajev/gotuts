package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main()  {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/count", counter)
	http.HandleFunc("/liss/", liss)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	incCount()
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	_, _ = fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	_, _ = fmt.Fprintf(w, "Host: %q\n", r.Host)
	_, _ = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); r != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func incCount() {
	mu.Lock()
	count++
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func liss(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); r != nil {
		log.Print(err)
	}
	cycles, size, delay := 5, 100, 8
	for k, v := range r.Form {
		switch k {
		case "cycles": cycles, _ = strconv.Atoi(v[0])
		case "size": size, _ = strconv.Atoi(v[0])
		case "delay": delay, _ = strconv.Atoi(v[0])
		}
	}
	lissajous(w, cycles, size, delay)
}

var palette = []color.Color{color.White, color.RGBA{0, 128, 0, 1},
	color.RGBA{255,0,0,1}, color.RGBA{255,255,0,1}}

func lissajous(out io.Writer, c, s, d int) {
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

