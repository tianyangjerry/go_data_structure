package binaryTree

// node 二叉搜索树 size表示以该节点为根的子树的节点个数 count 当前节点重复数量
type node struct {
	value int
	left  *node
	right *node

	size  int
	count int
}

type BST struct {
	root *node
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) Insert(value int) {
	b.root = insert(b.root, value)
}

func insert(nowNode *node, value int) *node {
	if nowNode == nil {
		return &node{
			value: value,
			size:  1,
			count: 1,
		}
	}

	if nowNode.value == value {
		nowNode.count++
		return nowNode
	}

	if nowNode.value < value {
		nowNode.right = insert(nowNode.right, value)
	} else {
		nowNode.left = insert(nowNode.left, value)
	}

	nowNode.size++
	return nowNode
}

func (b *BST) Search(value int) int {
	return search(b.root, value)
}

func search(nowNode *node, value int) int {
	if nowNode == nil {
		return 0
	}

	if nowNode.value == value {
		return nowNode.count
	}

	if nowNode.value < value {
		return search(nowNode.right, value)
	} else {
		return search(nowNode.left, value)
	}
}

func (b *BST) InorderTraversal() {
	inorderTraversal(b.root)
}

func inorderTraversal(nowNode *node) {
	if nowNode != nil {
		inorderTraversal(nowNode.left)
		println(nowNode.value)
		inorderTraversal(nowNode.right)
	}
}

func (b *BST) FindMin() (int, bool) {
	m, t := findMin(b.root)

	if t == false {
		return -1, false
	}

	return m, true
}

func findMin(nowNode *node) (int, bool) {
	if nowNode == nil {
		return -1, false
	}

	for nowNode.left != nil {
		nowNode = nowNode.left
	}

	return nowNode.value, true
}

func (b *BST) FindMax() (int, bool) {
	m, t := findMax(b.root)

	if t == false {
		return -1, false
	}

	return m, true
}

func findMax(nowNode *node) (int, bool) {
	if nowNode == nil {
		return -1, false
	}

	for nowNode.right != nil {
		nowNode = nowNode.right
	}

	return nowNode.value, true
}
