package gtree

type BinarySearchTree struct {
    BinaryTree
}

type binarySearchNode struct {
    binaryNode
}

func newSearchBinaryNode(key string, data string) *binarySearchNode {
    binaryNode := binaryNode{
        key: key,
        data: data,
        left: nil,
        right: nil,
    }
    return &binarySearchNode {
        binaryNode,
    }
}

func (t *BinarySearchTree) Insert(key string, data string) {
    if t.root == nil {
        t.root = &newSearchBinaryNode(key, data).binaryNode
    } else {
        (*binarySearchNode).insert(&t.root.(binarySearchNode), newSearchBinaryNode(key, data))
    }
}

func (n *binarySearchNode) insert(node *binarySearchNode) {
    if n.key == node.key {
        n.data = node.data
    } else if node.key < n.key {
        if n.left == nil {
            n.left = node
        } else {
            n.left.insert(node)
        }
    } else {
        if n.right == nil {
            n.right = node
        } else {
            n.right.insert(node)
        }
    }
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
