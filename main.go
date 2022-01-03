package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", homePage)
	//r.HandleFunc("/user/{username}", routage.ProfilePage)
	r.HandleFunc("/count", countPage)
	r.HandleFunc("/image/{cycle}/{size}", imagePage)
	http.ListenAndServe(":3000", r)
}
func imagePage(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	vars := mux.Vars(r)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cycles, _ := strconv.Atoi(vars["cycle"])
		size, _ := strconv.Atoi(vars["size"])
		lissajous(w, float64(cycles), size)
	}()
	wg.Wait()
	mu.Unlock()
}

func countPage(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	w.Write([]byte(fmt.Sprint("Count: ", count)))
	mu.Unlock()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	tmpl := template.Must(template.ParseFiles("./pages/html/home.html"))
	tmpl.Execute(w, nil)
	mu.Unlock()
}

var palette = []color.Color{color.RGBA{R: 255, A: 80}, color.RGBA{G: 255, A: 80}, color.RGBA{B: 255, A: 80}, color.RGBA{R: 255, G: 255, B: 255, A: 80}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles float64, size int) {
	var (
		res     = 0.001
		nframes = 64
		delay   = 8
	)
	size = size
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size*1, 2*size*1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			//convert size form int to float64
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
