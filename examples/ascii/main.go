package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/lmas/galaxygen"
)

// SIZE is the total size of the new galaxy to make (from -40 to +40).
const SIZE int = 80

func main() {
	// Set a "random" seed and make a new galaxy.
	rand.Seed(time.Now().UnixNano())
	galaxy := galaxygen.New(5000, 5, float64(SIZE)/2)

	// Make a ASCII map (it's about 30 rows and 80 columns wide, with the
	// central position at 0,0)
	m := makeMap(galaxy, 30, 0, 0)

	// Show the final map
	fmt.Println("+-----------------------------------------------------------------------------------+")
	for _, s := range strings.Split(m, "\n") {
		fmt.Printf("I %s I\n", s)
	}
	fmt.Println("+-----------------------------------------------------------------------------------+")
	fmt.Println("stars", len(galaxy))
}

func makeMap(gal []*galaxygen.Star, viewSize, startX, startY int) string {
	m := [SIZE + 1][SIZE + 1]string{}
	// Whipe it clean
	for y := 0; y < SIZE+1; y++ {
		for x := 0; x < SIZE+1; x++ {
			m[y][x] = " "
		}
	}

	// Add stars
	for _, s := range gal {
		x, y := s.X+SIZE/2, s.Y+SIZE/2
		m[y][x] = "."
	}

	// Mark the start position
	m[startY+SIZE/2][startX+SIZE/2] = "X"

	// Start "cutting" out a view area from the actual map.
	var ret string
	viewSize = viewSize / 2
	for y := viewSize - 1; y > -viewSize; y-- {
		y := y + startY + SIZE/2
		if y > SIZE {
			// We're at the border, can't show any more beyond it.
			// Add empty space as filling.
			ret += fmt.Sprint("                                                                                 \n")
			continue
		}
		if y < 0 {
			ret += fmt.Sprint("                                                                                 \n")
			continue
		}
		for x := 0; x < SIZE+1; x++ {
			ret += fmt.Sprint(m[y][x])
		}
		ret += fmt.Sprint("\n")
	}
	// Remove extra newline at the end.
	ret = ret[:len(ret)-1]
	return ret
}
