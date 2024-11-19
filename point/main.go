package main

import "github.com/01-edu/z01"

type point struct{
	x,y int
}

func setPoint(ptr *point) {
	
	ptr.x = 33
	ptr.y = 34
}

func main() {
	points := &point{}

	setPoint(points)
	z01.PrintRune()
	Printf("x = %d, y = %d\n", points.x, points.y)
}
