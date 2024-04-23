package ptsolver

func sumListWithSeparator(input []int) int {
	result := 0
	for _, hint := range input {
		result += hint + 1
	}
	return result - 1
}

func (b *Board) fillPosition(row, col int, symbol pointType) (bool, bool) {
	if b.solution[row][col] == symbol {
		return false, false
	}
	if b.solution[row][col] != UNSURE {
		b.solution[row][col] = CONFLICTED
		return true, true // conflicted
	}
	b.solution[row][col] = symbol
	return true, false
}

func (b *Board) rotateCounterClockwise() {
	// easy, not efficient
	b.rotateClockwise()
	b.rotateClockwise()
	b.rotateClockwise()
}

func (b *Board) rotateClockwise() {
	newRowHint := make([][]int, b.dimension)
	newColHint := make([][]int, b.dimension)

	newSolution := make([][]pointType, b.dimension)
	for i := 0; i < b.dimension; i++ {
		newSolution[i] = make([]pointType, b.dimension)
	}

	// transform
	for i := 0; i < b.dimension; i++ {
		newRowHint[i] = b.colHint[i]
		newColHint[i] = b.rowHint[(b.dimension-1)-i]
		for j := 0; j < b.dimension; j++ {
			newSolution[i][j] = b.solution[(b.dimension-1)-j][i]
		}
	}

	b.rowHint = newRowHint
	b.colHint = newColHint
	b.solution = newSolution
	b.rotatedDegree = (b.rotatedDegree + 90) % 360
}
