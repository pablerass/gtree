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
        t.root.insert(key, data)
    }
}

func (n *binaryNode) insert(key string, data string) {
    if key == n.key {
        n.data = data
    } else if key < n.key {
        if n.left == nil {
            n.left = newBinaryNode(key, data)
        } else {
            n.left.insert(key, data)
        }
    } else {
        if n.right == nil {
            n.right = newBinaryNode(key, data)
        } else {
            n.right.insert(key, data)
        }
    }
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