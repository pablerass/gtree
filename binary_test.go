package gtree

import (
    "fmt"
    "testing"
)

//      (b, 2)
//      /    \
// (a, 1)    (c, 3)
//                \
//                (n, 3)
//                /
//           (d, 2)
func ExampleInsert() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Insert("d", "2")
    bt.Print()
    // Output:
    // a 1
    // b 2
    // c 3
    // d 2
    // n 3
}

// (a, 1)
//      \
//      (c, 3)
//           \
//           (n, 3)
//           /
//      (d, 2)
func ExampleDelete() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Delete("d")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Delete("b")
    bt.Insert("d", "2")
    bt.Print()
    // Output:
    // a 1
    // c 3
    // d 2
    // n 3
}

func ExampleTraverseEmpty() {
    var bte BinaryTree
    bte.Print()
    // Output:
}

//      (b, 2)
//      /    \
// (a, 1)    (c, 3)
//                \
//                (n, 3)
//                /
//           (d, 2)
func ExamplePreorder() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Insert("d", "2")
    for entry := range bt.TraversePreorder() {
        entry.Print()
    }
    // Output:
    // b 2
    // a 1
    // c 3
    // n 3
    // d 2
}

//      (b, 2)
//      /    \
// (a, 1)    (c, 3)
//                \
//                (n, 3)
//                /
//           (d, 2)
func ExamplePostorder() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Insert("d", "2")
    for entry := range bt.TraversePostorder() {
        entry.Print()
    }
    // Output:
    // a 1
    // d 2
    // n 3
    // c 3
    // b 2
}

//      (b, 2)
//      /    \
// (a, 1)    (c, 3)
//                \
//                (n, 3)
//                /
//           (d, 2)
func ExamplePack() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Insert("d", "2")
    oldBtRoot := bt.root
    bt.Pack()
    for entry := range bt.TraversePreorder() {
        entry.Print()
    }
    // TODO: Implement this in a different test type
    if (oldBtRoot == bt.root) {
        fmt.Println("equal root")
    }
    // Output:
    // b 2
    // a 1
    // c 3
    // n 3
    // d 2
}

func TestSearch(t *testing.T) {
    var bt BinaryTree
    bt.Insert("a", "2")
    bt.Insert("b", "1")
    data1, _ := bt.Search("a")
    if data1 != "2" {
        t.Errorf("bt.Search(\"a\") = %s; want \"2\"", data1)
    }
    bt.Insert("a", "5")
    data2, _ := bt.Search("a")
    if data2 != "5" {
        t.Errorf("bt.Search(\"a\") = %s; want \"5\"", data2)
    }
}
