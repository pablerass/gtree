package gtree

import (
    "fmt"
)


//         (b, 2)
//         /    \
//      (c, 3)  (n, 3)
//      /    \  /
// (d, 2)(a, 1) (x, 5)
func ExampleBTInsert() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Insert("d", "2")
    bt.Insert("x", "5")
    bt.Print()
    // Output:
    // d 2
    // c 3
    // a 1
    // b 2
    // x 5
    // n 3
}

//     (c, 3)
//     /    \
//   (d, 2)(a, 1)
//   /    \
// (n, 3) (x, 5)
func ExampleBTDelete() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Delete("d")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Delete("b")
    bt.Insert("d", "2")
    bt.Insert("x", "5")
    bt.Print()
    // Output:
    // n 3
    // d 2
    // x 5
    // c 3
    // a 1
}

func ExampleBTTraverseEmpty() {
    var bte BinaryTree
    bte.Print()
    // Output:
}

//         (b, 2)
//         /    \
//      (c, 3)  (n, 3)
//      /    \  /
// (d, 2)(a, 1) (x, 5)
func ExampleBTPreorder() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Insert("d", "2")
    bt.Insert("x", "5")
    for entry := range bt.TraversePreorder() {
        entry.Print()
    }
    // Output:
    // b 2
    // c 3
    // n 3
    // d 2
    // a 1
    // x 5
}

//         (b, 2)
//         /    \
//      (c, 3)  (n, 3)
//      /    \  /
// (d, 2)(a, 1) (x, 5)
func ExampleBTPostorder() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Insert("n", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Insert("d", "2")
    bt.Insert("x", "5")
    for entry := range bt.TraversePostorder() {
        entry.Print()
    }
    // Output:
    // d 2
    // a 1
    // x 5
    // c 3
    // n 3
    // b 2
}

//         (b, 2)
//         /    \
//      (c, 3)  (n, 3)
//      /    \  /
// (d, 2)(a, 1) (x, 5)
func ExampleBTPack() {
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
    // c 3
    // n 3
    // d 2
    // a 1
    // x 5
}
