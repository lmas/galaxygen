package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/lmas/galaxygen"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Seed the RNG so we get new galaxies on each run.
	rand.Seed(time.Now().UnixNano())

	// Make a new galaxy now and show some simple stats.
	start := time.Now()
	starmap := galaxygen.New(10000, 5, 100.0)
	dur := time.Since(start)
	stars := len(starmap)
	fmt.Println("Generation time:", dur)
	fmt.Println("Stars:", stars)

	// Add the stars to a scatter plot, so we can visualize the galaxy.
	pts := make(plotter.XYs, stars)
	for i := 0; i < stars; i++ {
		pts[i].X = float64(starmap[i].X)
		pts[i].Y = float64(starmap[i].Y)
	}
	s, e := plotter.NewScatter(pts)
	checkErr(e)
	s.Radius = 1

	// Make the plot and save it.
	p, e := plot.New()
	checkErr(e)
	p.Add(s)
	p.Title.Text = fmt.Sprintf("Galaxy with %d stars", stars)
	p.Legend.Add("Star", s)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	e = p.Save(512, 512, "galaxy.png")
	checkErr(e)
}
