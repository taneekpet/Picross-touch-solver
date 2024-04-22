package ptsolver

func sumListWithSeparator(input []int) int {
	result := 0
	for _, hint := range input {
		result += hint + 1
	}
	return result - 1
}

func (b *Board) fillRowWithHint(row int) (changed bool, conflicted bool) {
	col := 0
	for i, hint := range b.rowHint[row] {
		if i != 0 {
			if b.solution[row][col] != EMPTY { // changed
				if b.solution[row][col] != UNSURE { // conflicted
					return true, true
				}
				b.solution[row][col] = EMPTY
				changed = true
			}
			col++
		}
		for count := 0; count < hint; count++ {
			if b.solution[row][col] != FILLED { // changed
				if b.solution[row][col] != UNSURE { // conflicted
					return true, true
				}
				b.solution[row][col] = FILLED
				changed = true
			}
			col++
		}
	}
	return false, false
}

func (b *Board) sumToDimension() (bool, bool) {
	for i := 0; i < b.dimension; i++ {
		if sumListWithSeparator(b.rowHint[i]) == b.dimension {
			changed, conflicted := b.fillRowWithHint(i)
			if conflicted {
				return changed, true
			}
		}
	}
	// to be implemented more
	return false, false
}

// to add more solver
