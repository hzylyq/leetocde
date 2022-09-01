package backtrack

import (
	"log"
	"testing"
)

func TestPermute(t *testing.T) {
	permute([]int{1, 2, 3})
}

func TestCombine(t *testing.T) {
	res := combine(4, 2)
	log.Print(res)
}
