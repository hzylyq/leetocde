package tree

import "testing"

func TestGenerateTrees(t *testing.T) {
	res := GenerateTrees(3)
	t.Log(res)
}
