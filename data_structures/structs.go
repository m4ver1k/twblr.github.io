package data_structures
// Linked List Implementation

type Node struct{
	data int
	next *Node
}
var top *Node
func Add(elem int) {
	var cur *Node

	if top == nil{
		top = &Node{data:elem}
	}else{
		cur = top
		for cur.next != nil{
			cur = cur.next
		}
		cur.next = &Node{data:elem}
	}
	return
}

func Search(elem int) *Node {
	cur  := top
	for cur !=nil{
		if cur.data==elem{
			return cur
		}else{
			cur=cur.next
		}
	}
	return nil
}
