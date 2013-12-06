package collect

type OrderedMap struct {
	root   *omap_cell
	less   func(interface{}, interface{}) bool
	length int
}

type omap_cell struct {
	key, value  interface{}
	red         bool
	left, right *omap_cell
}

func NewOrderedMap(less func(interface{}, interface{}) bool) *OrderedMap {
	return &OrderedMap{less: less}
}

func (m *OrderedMap) Insert(key, value interface{}) (inserted bool) {
	m.root, inserted = m.insert(m.root, key, value)
	m.root.red = false
	if inserted {
		m.length++
	}
	return inserted
}

func (m *OrderedMap) insert(root *omap_cell, key, value interface{}) (*omap_cell, bool) {
	inserted := false
	if root == nil {
		return &omap_cell{key: key, value: value, red: true}, true
	}
	if isRed(root.left) && isRed(root.right) {
		colorFlip(root)
	}
	if m.less(key, root.key) {
		root.left, inserted = m.insert(root.left, key, value)
	} else if m.less(root.key, key) {
		root.right, inserted = m.insert(root.right, key, value)
	} else {
		root.value = value
	}

	if isRed(root.right) && !isRed(root.left) {
		root = rotateLeft(root)
	}
	if isRed(root.left) && isRed(root.left.left) {
		root = rotateRight(root)
	}
	return root, inserted
}

func isRed(root *omap_cell) bool { return root != nil && root.red }

func colorFlip(root *omap_cell) {
	root.red = !root.red
	if root.left != nil {
		root.left.red = !root.left.red
	}
	if root.right != nil {
		root.right.red = !root.right.red
	}
}

func rotateLeft(root *omap_cell) *omap_cell {
	x := root.right
	root.right = x.left
	x.left = root
	x.red = root.red
	root.red = true
	return x
}

func rotateRight(root *omap_cell) *omap_cell {
	x := root.left
	root.left = x.right
	x.right = root
	x.red = root.red
	root.red = true
	return x
}

func (m *OrderedMap) Find(key interface{}) (value interface{}, found bool) {
	root := m.root
	for root != nil {
		if m.less(key, root.key) {
			root = root.left
		} else if m.less(root.key, key) {
			root = root.right
		} else {
			return root.value, true
		}
	}
	return nil, false
}

func (m *OrderedMap) Do(function func(interface{}, interface{})) {
	do(m.root, function)
}

func do(root *omap_cell, fun func(interface{}, interface{})) {
	if root != nil {
		do(root.left, fun)
		fun(root.key, root.value)
		do(root.right, fun)
	}
}

func (m *OrderedMap) Len() int {
	return m.length
}
