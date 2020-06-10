package test_algorithm

// 二叉树排序
type tree struct {
	value int
	left  *tree
	right *tree
}

func TreeSort(values []int) []int {
	var root *tree
	for _, value := range values {
		root = addTree(root, value)
	}
	values = appendValues(values[:0], root)
	return values
}

func addTree(t *tree, value int) *tree {
	if t == nil {
		t = &tree{}
		t.value = value
		return t
	}
	if value < t.value {
		t.left = addTree(t.left, value)
	} else {
		t.right = addTree(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
