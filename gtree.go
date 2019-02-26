package gtree

import (
    "fmt"
)

type Tree struct {
    key string
    data string
    left *Tree
    right *Tree
}

func New(key string, data string) *Tree {
    return &Tree{
        key: key,
        data: data,
        left: nil,
        right: nil,
    }
}

func (t *Tree) Insert(key string, data string) {
    if key == t.key {
        t.data = data
    } else if key < t.key {
        if t.left == nil {
            t.left = New(key, data)
        } else {
            t.left.Insert(key, data)
        }
    } else {
        if t.right == nil {
            t.right = New(key, data)
        } else {
            t.right.Insert(key, data)
        }
    }
}

func (t *Tree) Search(key string) (string, bool) {
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

func (t *Tree) Print() {
    if t.left != nil {
        t.left.Print()
    }
    fmt.Println(t.key)
    if t.right != nil {
        t.right.Print()
    }
}