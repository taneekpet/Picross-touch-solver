package ptsolver

func (b *Board) countConsecutiveEmptyInRow(row int) (start, end int) {
	for i := 0; i < b.dimension; i++ {
		if b.solution[row][i] != EMPTY {
			break
		}
		start++
	}
	for i := b.dimension - 1; i >= 0; i-- {
		if b.solution[row][i] != EMPTY {
			break
		}
		end++
	}
	return
}

func (b *Board) fillFullRowWithHint(row, offset int) (resultChanged, resultConflicted bool) {
	col := offset
	for i, hint := range b.rowHint[row] {
		// fill between
		if i != 0 {
			changed, conflicted := b.fillPosition(row, col, EMPTY)
			if conflicted {
				return resultChanged, true
			}
			resultChanged = resultChanged || changed
			col++
		}
		// fill as hint
		for count := 0; count < hint; count++ {
			changed, conflicted := b.fillPosition(row, col, FILLED)
			if conflicted {
				return resultChanged, true
			}
			resultChanged = resultChanged || changed
			col++
		}
	}
	return
}

func (b *Board) sumToDimension() (resultChanged, resultConflicted bool) {
	for i := 0; i < b.dimension; i++ {
		emptyStart, emptyEnd := b.countConsecutiveEmptyInRow(i)
		if emptyStart+sumListWithSeparator(b.rowHint[i])+emptyEnd == b.dimension {
			changed, conflicted := b.fillFullRowWithHint(i, emptyStart)
			if conflicted {
				return resultChanged, true
			}
			resultChanged = resultChanged || changed
		}
	}
	return
}

// =======================================================================================

func (b *Board) checkStartOrEndOfRow(row int) (resultChanged, resultConflicted bool) {
	// start
	if b.solution[row][0] == FILLED {
		for col := 0; col < b.rowHint[row][0]; col++ {
			changed, conflicted := b.fillPosition(row, col, FILLED)
			if conflicted {
				return resultChanged, true
			}
			resultChanged = resultChanged || changed
		}
	}
	// end
	if b.solution[row][b.dimension-1] == FILLED {
		for col := b.dimension - 1; col >= 0; col-- {
			changed, conflicted := b.fillPosition(row, col, FILLED)
			if conflicted {
				return resultChanged, true
			}
			resultChanged = resultChanged || changed
		}
	}
	return
}

func (b *Board) checkStartOrEndOf() (resultChanged, resultConflicted bool) {
	for i := 0; i < b.dimension; i++ {
		changed, conflicted := b.checkStartOrEndOfRow(i)
		if conflicted {
			return resultChanged, true
		}
		resultChanged = resultChanged || changed
	}
	return
}

// =======================================================================================

// to add more solver
