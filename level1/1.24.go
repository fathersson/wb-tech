package main

import (
	"fmt"
	"math"
)

// Point - структура, представляющая точку на плоскости
// Поля x и y пишутся со строчной буквы, чтобы быть
// "приватными" (инкапсулированными) для этого пакета
type Point struct {
	x float64
	y float64
}

// NewPoint - конструктор для создания новой точки
// Так как поля x и y приватные, это единственный способ
// создать точку из другого пакета
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// DistanceTo рассчитывает расстояние от текущей точки до другой точки p2
// Используется теорема Пифагора
func (p *Point) DistanceTo(p2 *Point) float64 {
	// Формула: sqrt( (x2-x1)^2 + (y2-y1)^2 )
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}

func main() {
	p1 := NewPoint(2.1, 9.6)
	p2 := NewPoint(8.7, 5.9)

	distance := p1.DistanceTo(p2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
