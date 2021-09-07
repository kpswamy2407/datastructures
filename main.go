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
	fmt.Println(linkedList.FindNode(14))
	fmt.Println(linkedList.Traverse()...)

}
