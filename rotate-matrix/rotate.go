package main

import "fmt"

func algo1(n int, matrix [][]uint32) [][]uint32 {
	result := make([][]uint32, n)
	for i := 0; i < n; i++ {
		result[i] = make([]uint32, n)
	}

	fmt.Println(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[j][n-1-i] = matrix[i][j]
		}
	}

	return result
}

type Point struct {
	X uint32
	Y uint32
}

func algo2(n int, matrix [][]uint32) [][]uint32 {
	start := Point{0, 0}
	end := Point{0, uint32(n) - 1}

	for start.X < end.Y {
		swap(matrix, &start, &end)
		start.X += 1
		start.Y += 1
		end.X += 1
		end.Y -= 1
	}
	return matrix
}

func swap(matrix [][]uint32, start *Point, end *Point) {
	for i := 0; i < int(end.Y-start.Y); i++ {
		var tmp uint32
		s := Point{start.X + uint32(i), start.Y}
		e := Point{s.Y, uint32(len(matrix)-1) - s.X}
		last := matrix[s.X][s.Y]

		for i := 0; i < 4; i++ {
			tmp = matrix[e.X][e.Y]
			matrix[e.X][e.Y] = last
			last = tmp
			s.X = e.X
			s.Y = e.Y
			t := e.Y
			e.Y = uint32(len(matrix)-1) - e.X
			e.X = t
		}
	}
}

func main() {
	n := 4
	matrix := [][]uint32{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	result := algo1(n, matrix)
	fmt.Println("algorithm1:", result)

	result = algo2(n, matrix)
	fmt.Println("algorithm2:", result)
}
