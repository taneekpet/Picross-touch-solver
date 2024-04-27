package ptsolver

import (
	"fmt"
	"testing"
)

func Test_checkEmpty(_ *testing.T) {
	b := Init(
		5,
		[][]int{
			{1},
			{1},
			{},
			{1},
			{1},
		},
		[][]int{
			{0},
			{1},
			{1},
			{1},
			{0},
		},
	)
	checkEmpty(&b)
	b.rotateClockwise()
	checkEmpty(&b)
	b.rotateCounterClockwise()
	fmt.Printf("%v", b.PrettyPrint())
}
