package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func main() {
	// создаем экземпляры структур с координатами
	a := NewPoint(4, 8)
	b := NewPoint(1, 10)

	// вызываем функцию расчета расстояния между точками
	distance := CalcDistance(a, b)

	fmt.Println(distance)
}

func CalcDistance(a, b Point) float64 {
	// находим расстояние между точками по формуле
	return math.Sqrt(math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2))
}
