package gtree

import (
    "fmt"
)

type Tree struct {
    value int
    left *Tree
    right *Tree
}

func New(value int) Tree {
    return Tree{
        value: value,
    }
}

func (t Tree) Insert(value int) {
    if value < t.value {
        if t.left == nil {
            newTree := New(value)
            t.left = &newTree
        } else {
            t.left.Insert(value)
        }
    } else {
        if t.right == nil {
            newTree := New(value)
            t.right = &newTree
        } else {
            t.right.Insert(value)
        }
    }
}

func (t Tree) Print() {
    if t.left != nil {
        t.left.Print()
    }
    fmt.Println(t.value)
    if t.right != nil {
        t.right.Print()
    }
}