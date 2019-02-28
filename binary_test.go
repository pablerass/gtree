package gtree

import (
    "testing"
)

func ExampleInsert() {
    var bt BinaryTree
    bt.Insert("b", "2")
    bt.Insert("c", "3")
    bt.Insert("d", "4")
    bt.Insert("a", "1")
    bt.Insert("b", "5")
    bt.Print()
    // Output:
    // a 1
    // b 5
    // c 3
    // d 4
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
