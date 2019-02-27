package gtree

import (
    "testing"
)

func ExampleInsert() {
    bt := NewBinaryTree("b", "2")
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
    bt := NewBinaryTree("a", "2")
    data, _ := bt.Search("a")
    if data != "2" {
        t.Errorf("bt.Search(\"a\") = %s; want \"2\"", data)
    }
}
