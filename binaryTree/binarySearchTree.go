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

func (b *BST) Insert(value int) {
	b.root = insert(b.root, value)
}

func insert(nowNode *node, value int) *node {
	if nowNode == nil {
		return &node{value: value}
	}

	if value < nowNode.value {
		nowNode.left = insert(nowNode.left, value)
	} else if value > nowNode.value {
		nowNode.right = insert(nowNode.right, value)
	} else {
		nowNode.count++
	}

	nowNode.size = nowNode.count + nowNode.left.size + nowNode.right.size

	return nowNode
}

func (b *BST) Search(value int) bool {
	return search(b.root, value)
}

func search(nowNode *node, value int) bool {
	if nowNode == nil {
		return false
	}
	if nowNode.value == value {
		return true
	} else if nowNode.value < value {
		return search(nowNode.left, value)
	} else {
		return search(nowNode.right, value)
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
