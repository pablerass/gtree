package gtree

import (
    "fmt"
)

type BinaryTree struct {
    root *binaryNode
}

type binaryNode struct {
    key string
    data string
    left *binaryNode
    right *binaryNode
}

func newBinaryNode(key string, data string) *binaryNode {
    return &binaryNode{
        key: key,
        data: data,
        left: nil,
        right: nil,
    }
}

func (t *BinaryTree) Insert(key string, data string) {
    if t.root == nil {
        t.root = newBinaryNode(key, data)
    } else {
        t.root.insert(newBinaryNode(key, data))
    }
}

func (n *binaryNode) insert(node *binaryNode) {
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

func (t *BinaryTree) Delete(key string) {
    if t.root != nil {
        t.root = t.root.delete(key)
    }
}

func (n *binaryNode) delete(key string) *binaryNode {
    if key == n.key {
        if n.left == nil && n.right == nil {
            n = nil
        } else if n.left == nil {
            n = n.right
        } else if n.right == nil{
            n = n.left
        } else {
            old := n
            n = n.left
            n.insert(old.right)
        }
    } else if key < n.key {
        if n.left != nil {
            n.left = n.left.delete(key)
        }
    } else {
        if n.right != nil {
            n.right = n.right.delete(key)
        }
    }
    return n
}

func (t *BinaryTree) Search(key string) (string, bool) {
    if t.root == nil {
        return "", false
    } else {
        return t.root.search(key)
    }
}

func (n *binaryNode) search(key string) (string, bool) {
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

func (t *BinaryTree) Print() {
    if t.root != nil {
        t.root.print()
    }
}

func (n *binaryNode) print() {
    if n.left != nil {
        n.left.print()
    }
    fmt.Println(n.key, n.data)
    if n.right != nil {
        n.right.print()
    }
}