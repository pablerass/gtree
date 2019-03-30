package gtree

import (
    "fmt"
)

type BinaryTree struct {
    root *binaryNode
}

type Entry struct {
    key string
    data string
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

func (t *BinaryTree) InsertEntry(entry Entry) {
    t.Insert(entry.key, entry.data)
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

func traverse(t BinaryTree, traverseFunction func(binaryNode, chan Entry)) <-chan Entry {
    ch := make(chan Entry)
    go func() {
        if t.root != nil {
            traverseFunction(*t.root, ch)
        }
        close(ch)
    }()
    return ch
}

func (t BinaryTree) TraverseInorder() <-chan Entry {
    return traverse(t, binaryNode.traverseInorder)
}

func (t BinaryTree) TraversePreorder() <-chan Entry {
    return traverse(t, binaryNode.traversePreorder)
}

func (t BinaryTree) TraversePostorder() <-chan Entry {
    return traverse(t, binaryNode.traversePostorder)
}

func (n binaryNode) traverseInorder(ch chan Entry) {
    if n.left != nil {
        n.left.traverseInorder(ch)
    }
    ch <- Entry{key: n.key, data: n.data}
    if n.right != nil {
        n.right.traverseInorder(ch)
    }
}

func (n binaryNode) traversePreorder(ch chan Entry) {
    ch <- Entry{key: n.key, data: n.data}
    if n.left != nil {
        n.left.traversePreorder(ch)
    }
    if n.right != nil {
        n.right.traversePreorder(ch)
    }
}

func (n binaryNode) traversePostorder(ch chan Entry) {
    if n.left != nil {
        n.left.traversePostorder(ch)
    }
    if n.right != nil {
        n.right.traversePostorder(ch)
    }
    ch <- Entry{key: n.key, data: n.data}
}

func (t *BinaryTree) Pack() {
    if t.root != nil {
        var nt BinaryTree
        for entry := range t.TraversePreorder() {
            nt.InsertEntry(entry)
        }
        t.root = nt.root
    }
}

func (t BinaryTree) Print() {
    for entry := range t.TraverseInorder() {
        entry.Print()
    }
}

func (e Entry) Print() {
    fmt.Println(e.key, e.data)
}