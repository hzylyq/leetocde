package tries

import "testing"

func TestMultiStringSearch(t *testing.T) {
	MultiStringSearch("this is a big string", []string{"this", "yo", "is", "a", "bigger", "string", "kappa"})
}
