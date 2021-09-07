package singlelinkedlist

import "errors"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Traverse() []interface{} {

	var res []interface{}
	last := l.head
	for last != nil {
		res = append(res, last.data)
		last = last.next
	}

	return res
}

func (l *LinkedList) Insert(data interface{}) {
	node := &Node{
		data: data,
		next: nil,
	}
	if l.length == 0 {
		l.head = node
		l.length++
		return
	}
	last := l.head
	for i := 1; i <= l.length; i++ {
		if last.next == nil {
			last.next = node
			l.length++
			return
		}
		last = last.next
	}
}

func (l *LinkedList) InsertAtFirst(data interface{}) {
	node := &Node{
		data: data,
		next: nil,
	}
	if l.length == 0 {
		l.head = node
		l.length++
		return
	}
	node.next = l.head
	l.head = node
	l.length++
	return
}

func (l *LinkedList) InsertAfter(prevNode *Node, data interface{}) error {
	if prevNode == nil {
		return errors.New("previous node should not be empty")
	}
	node := &Node{
		data: data,
		next: prevNode,
	}
	prevNode.next = node
	l.length++
	return nil
}

func (l *LinkedList) FindNode(needle interface{}) *Node {
	last := l.head
	for last != nil {
		if last.data == needle {
			return last
		}
		last = last.next
	}
	return nil
}
