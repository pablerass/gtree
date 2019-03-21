package gtree

import (
    "testing"
)

//      (b, 2)
//      /    \
// (a, 1)    (c, 3)
//                \
//                (n, 3)
//                /
//           (d, 2)
func ExampleBSTInsert() {
    var bst BinarySearchTree
    bst.Insert("b", "2")
    bst.Insert("c", "3")
    bst.Insert("n", "3")
    bst.Insert("d", "4")
    bst.Insert("a", "1")
    bst.Insert("d", "2")
    bst.Print()
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
func ExampleBSTDelete() {
    var bst BinarySearchTree
    bst.Insert("b", "2")
    bst.Insert("c", "3")
    bst.Delete("d")
    bst.Insert("n", "3")
    bst.Insert("d", "4")
    bst.Insert("a", "1")
    bst.Delete("b")
    bst.Insert("d", "2")
    bst.Print()
    // Output:
    // a 1
    // c 3
    // d 2
    // n 3
}

//      (b, 2)
//      /    \
// (a, 1)    (c, 3)
//                \
//                (n, 3)
//                /
//           (d, 2)
func ExampleBSTPreorder() {
    var bst BinarySearchTree
    bst.Insert("b", "2")
    bst.Insert("c", "3")
    bst.Insert("n", "3")
    bst.Insert("d", "4")
    bst.Insert("a", "1")
    bst.Insert("d", "2")
    for entry := range bst.TraversePreorder() {
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
func ExampleBSTPostorder() {
    var bst BinarySearchTree
    bst.Insert("b", "2")
    bst.Insert("c", "3")
    bst.Insert("n", "3")
    bst.Insert("d", "4")
    bst.Insert("a", "1")
    bst.Insert("d", "2")
    for entry := range bst.TraversePostorder() {
        entry.Print()
    }
    // Output:
    // a 1
    // d 2
    // n 3
    // c 3
    // b 2
}

func TestBSTSearch(t *testing.T) {
    var bst BinarySearchTree
    bst.Insert("a", "2")
    bst.Insert("b", "1")
    data1, _ := bst.Search("a")
    if data1 != "2" {
        t.Errorf("bst.Search(\"a\") = %s; want \"2\"", data1)
    }
    bst.Insert("a", "5")
    data2, _ := bst.Search("a")
    if data2 != "5" {
        t.Errorf("bst.Search(\"a\") = %s; want \"5\"", data2)
    }
}
