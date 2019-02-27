package gtree

import (
    "fmt"
)

type BinaryTree struct {
    key string
    data string
    left *BinaryTree
    right *BinaryTree
}

func NewBinaryTree(key string, data string) *BinaryTree {
    return &BinaryTree{
        key: key,
        data: data,
        left: nil,
        right: nil,
    }
}

func (t *BinaryTree) Insert(key string, data string) {
    if key == t.key {
        t.data = data
    } else if key < t.key {
        if t.left == nil {
            t.left = NewBinaryTree(key, data)
        } else {
            t.left.Insert(key, data)
        }
    } else {
        if t.right == nil {
            t.right = NewBinaryTree(key, data)
        } else {
            t.right.Insert(key, data)
        }
    }
}

func (t *BinaryTree) Search(key string) (string, bool) {
    if key == t.key {
        return t.data, true
    } else if key < t.key {
        if t.left == nil {
            return "", false
        } else {
            return t.left.Search(key)
        }
    } else {
        if t.right == nil {
            return "", false
        } else {
            return t.right.Search(key)
        }
    }
}

func (t *BinaryTree) Print() {
    if t.left != nil {
        t.left.Print()
    }
    fmt.Println(t.key, t.data)
    if t.right != nil {
        t.right.Print()
    }
}