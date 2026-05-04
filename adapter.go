package main

import (
	"fmt"
	"math"
)

type HasRadius interface {
	getRadius() float64
}

type RoundHole struct {
	radius float64
}

type RoundPeg struct {
	radius float64
}

func (r RoundHole) getRadius () float64 {
	return r.radius
}

func (r RoundHole) fits (p HasRadius) bool {
	return r.radius >= p.getRadius()
}

func NewRoundHole (r float64) *RoundHole {
	return &RoundHole {
		radius: r,
	}
}

func (p RoundPeg) getRadius () float64 {
	return p.radius
}

func NewRoundPeg (r float64) RoundPeg {
	return RoundPeg {
		radius: r,
	}
}

type SquarePeg struct {
	width float64
}

func NewSquarePeg (w float64) SquarePeg {
	return SquarePeg {
		width: w,
	}
}

func (p SquarePeg) getWidth () float64 {
	return p.width
}

type SquarePegAdapter struct {
	width float64
}

func (a SquarePegAdapter) getRadius () float64 {
	return math.Sqrt(2) * a.width / 2
}

func NewSquarePegAdapter (w float64) SquarePegAdapter {
	return SquarePegAdapter {
		width: w,
	}
}


func main () {
	hole := NewRoundHole(10)
	// peg := NewRoundPeg(11)
	squarePeg := NewSquarePegAdapter(12)

	fmt.Printf("hole radius %f\n", hole.getRadius())
	fmt.Printf("hole fits peg %v\n", hole.fits(squarePeg))
}