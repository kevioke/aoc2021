package main

func getAllAdjacent(row, col, rowLen, colLen int) [][]int {
	adjacent := [][]int{}
	if row+1 < rowLen {
		adjacent = append(adjacent, []int{row + 1, col})
		if col+1 < colLen {
			adjacent = append(adjacent, []int{row + 1, col + 1})
		}
		if col-1 >= 0 {
			adjacent = append(adjacent, []int{row + 1, col - 1})
		}
	}
	if row-1 >= 0 {
		adjacent = append(adjacent, []int{row - 1, col})
		if col+1 < colLen {
			adjacent = append(adjacent, []int{row - 1, col + 1})
		}
		if col-1 >= 0 {
			adjacent = append(adjacent, []int{row - 1, col - 1})
		}
	}
	if col+1 < colLen {
		adjacent = append(adjacent, []int{row, col + 1})
	}
	if col-1 >= 0 {
		adjacent = append(adjacent, []int{row, col - 1})
	}
	return adjacent
}
