package gtree

type BinarySearchTree struct {
    BinaryTree
}

func (t *BinarySearchTree) InsertEntry(entry Entry) {
    t.Insert(entry.key, entry.data)
}

func (t *BinarySearchTree) Insert(key string, data string) {
    if t.root == nil {
        t.root = newBinaryNode(key, data)
    } else {
        t.root.insertSearch(newBinaryNode(key, data))
    }
}

func (n *binaryNode) insertSearch(node *binaryNode) {
    if n.key == node.key {
        n.data = node.data
    } else if node.key < n.key {
        if n.left == nil {
            n.left = node
       } else {
            n.left.insertSearch(node)
        }
    } else {
        if n.right == nil {
            n.right = node
        } else {
            n.right.insertSearch(node)
        }
    }
}

func (t *BinarySearchTree) Delete(key string) {
    if t.root != nil {
        t.root = t.root.deleteSearch(key)
    }
}

func (n *binaryNode) deleteSearch(key string) *binaryNode {
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

func (t BinarySearchTree) Search(key string) (string, bool) {
    if t.root == nil {
        return "", false
    } else {
        return t.root.search(key)
    }
}

func (n binaryNode) search(key string) (string, bool) {
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