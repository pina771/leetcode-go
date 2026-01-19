package main

import (
	"fmt"
)

func main() {
	fmt.Println(computeArea(
		-2, -2, 2, 2, -3, -3, 3, -1,
	))
}

type Point struct {
	x, y int
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	// if they are not intersected
	if ax2 <= bx1 || bx2 <= ax1 || ay1 >= by2 || by1 >= ay2 {
		return (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	}
	rectA := (ax2 - ax1) * (ay2 - ay1)
	rectB := (bx2 - bx1) * (by2 - by1)

	botLeft := Point{
		min(ax1, bx1),
		min(ay1, by1),
	}
	topRight := Point{
		max(ax2, bx2),
		max(ay2, by2),
	}
	// if B is inside A
	if botLeft.x == ax1 && botLeft.y == ay1 && topRight.x == ax2 && topRight.y == ay2 {
		return rectA
	}
	// if A is inside B
	if botLeft.x == bx1 && botLeft.y == by1 && topRight.x == bx2 && topRight.y == by2 {
		return rectB
	}

	retval := rectA + rectB

	iBL := Point{max(bx1, ax1), max(by1, ay1)}
	iTR := Point{min(ax2, bx2), min(by2, ay2)}
	intersect := (iTR.y - iBL.y) * (iTR.x - iBL.x)
	fmt.Printf("intersection: bottomLeft = %+v \t topRight = %+v \t area = %d\n", iBL, iTR, intersect)
	return retval - intersect
}
