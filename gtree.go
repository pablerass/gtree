package gtree

import (
    "fmt"
)

type Tree struct {
    value int
    child1 *Tree
    child2 *Tree
}

func New(value int) *Tree {
    return &Tree{
        value: value,
        child1: nil,
        child2: nil,
    }
}

func (t *Tree) Insert(value int) {
    if value < t.value {
        if t.child1 == nil {
            t.child1 = New(value)
        } else {
            t.child1.Insert(value)
        }
    } else {
        if t.child2 == nil {
            t.child2 = New(value)
        } else {
            t.child2.Insert(value)
        }
    }
}

func (t Tree) Print() {
    if t.child1 != nil {
        t.child1.Print()
    }
    fmt.Println(t.value)
    if t.child2 != nil {
        t.child2.Print()
    }
}