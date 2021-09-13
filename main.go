package main

import (
	"fmt"

	"github.com/mss/datastructures/singlelinkedlist"
)

func main() {
	linkedList := &singlelinkedlist.LinkedList{}
	linkedList.Insert(7)
	linkedList.Insert(14)
	linkedList.Insert(24)
	linkedList.InsertAtFirst("Welcome home")
	previous := linkedList.FindNode(14)
	linkedList.InsertAfter(previous, 1)
	fmt.Println(linkedList.Traverse()...)

}
