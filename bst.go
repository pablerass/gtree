package gtree

type BinarySearchTree struct {
    BinaryTree
}

type binarySearchNode struct {
    binaryNode
}

func (t BinarySearchTree) Search(key string) (string, bool) {
    if t.root == nil {
        return "", false
    } else {
        return t.root.search(key)
    }
}

func (n binarySearchNode) search(key string) (string, bool) {
    if key == n.key {
        return n.data, true
    } else if key < n.key {
        if n.left == nil {
            return "", false
        } else {
            return n.left.search(key)
        }
    } else {
        if n.right == nil {
            return "", false
        } else {
            return n.right.search(key)
        }
    }
}
