package ptsolver

import "fmt"

type pointType int

const UNSURE pointType = 0
const EMPTY pointType = 1
const FILLED pointType = 2
const CONFLICTED pointType = 3

type Board struct {
	dimension int
	rowHint   [][]int
	colHint   [][]int

	unsureSymbol     string
	emptySymbol      string
	filledSymbol     string
	conflictedSymbol string

	solution   [][]pointType
	solved     bool
	conflicted bool

	solverFunctionList []func() (bool, bool)
}

func (b *Board) PrettyPrint() error {
	if len(b.solution) != b.dimension {
		return fmt.Errorf("number of rows not match dimension")
	}
	for i := 0; i < b.dimension; i++ {
		if len(b.solution[i]) != b.dimension {
			return fmt.Errorf("number of column at row: %d not match dimension", i)
		}
	}
	if len(b.unsureSymbol) != 1 {
		return fmt.Errorf("unsure symbol should be 1 charactor")
	}
	if len(b.emptySymbol) != 1 {
		return fmt.Errorf("empty symbol should be 1 charactor")
	}
	if len(b.filledSymbol) != 1 {
		return fmt.Errorf("filled symbol should be 1 charactor")
	}
	if len(b.conflictedSymbol) != 1 {
		return fmt.Errorf("conflicted symbol should be 1 charactor")
	}

	printer := ""
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.solution[i][j] == FILLED {
				printer += b.filledSymbol
			} else if b.solution[i][j] == EMPTY {
				printer += b.emptySymbol
			} else if b.solution[i][j] == CONFLICTED {
				printer += b.conflictedSymbol
			} else {
				printer += b.unsureSymbol
			}
		}
		printer += "\n"
	}
	fmt.Printf("%s", printer)
	return nil
}

func (b *Board) SetCharactor(unsure, empty, filled, conflicted string) {
	b.unsureSymbol = unsure
	b.emptySymbol = empty
	b.filledSymbol = filled
	b.conflictedSymbol = conflicted
}

func (b *Board) IsConflicted() bool {
	return b.conflicted
}

func (b *Board) IsSolved() bool {
	if b.conflicted {
		return false
	}
	if b.solved {
		return true
	}
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.solution[i][j] == UNSURE {
				return false
			}
		}
	}
	b.solved = true
	return true
}

func (b *Board) Solve() bool {
	if b.solved {
		return true
	}
	changed := true
	for changed {
		changed = false
		for _, solver := range b.solverFunctionList {
			tmpChanged, conflicted := solver()
			changed = changed || tmpChanged
			if conflicted {
				b.conflicted = true
				break
			}
		}
	}
	return b.IsSolved()
}

func Init(dimension int, rowHint, colHint [][]int) Board {
	sol := make([][]pointType, dimension)
	for i := 0; i < dimension; i++ {
		sol[i] = make([]pointType, dimension)
	}
	b := Board{
		dimension: dimension,
		rowHint:   rowHint,
		colHint:   colHint,

		unsureSymbol: "?",
		emptySymbol:  " ",
		filledSymbol: "X",

		solution:   sol,
		solved:     false,
		conflicted: false,
	}
	b.solverFunctionList = []func() (bool, bool){
		b.sumToDimension,
		// to be added
	}
	return b
}
