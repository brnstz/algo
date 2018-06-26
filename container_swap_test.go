package algo

import "testing"

func TestContainerSwap(t *testing.T) {
	container := [][]int{
		{1, 1},
		{1, 1},
	}

	t.Fatal(ContainerSwap(container))

}
