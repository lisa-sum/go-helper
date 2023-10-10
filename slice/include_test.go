package slice

import "testing"

func TestInclude(t *testing.T) {
	bool := Include([]int{1, 2, 3}, 1)
	t.Log(bool)
}
