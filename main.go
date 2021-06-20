package main 

import (
	"fmt"
	
)


 type Node struct {
	Previous *Node
	Value int
	Next *Node

 }
//сознание указателя на узел
 var root = new(Node)

 func AddNode(n *Node, a int) int {
	 if root == nil {
		n = &Node{nil,a,nil}
		root = n
		return 0
	 }

	 if n.Value == a {
		 fmt.Println("Node alredy")
		 return -1
	 }

	 if n.Next == nil {
		 n.Next = &Node{n,a,nil}
		 return -2
	 }
	 return AddNode(n.Next,a)
 }


 func Traverse(n *Node) {
	 if n == nil {
		fmt.Println("<-> Empty list")
	 }
// пробегаемся по списку 
	 for n != nil {
		 fmt.Printf(" %d <-> ", n.Value)
		 n = n.Next
	 }
 }
 func Delete(n *Node, element int ) {
	for n != nil {
		if n.Value == element {
			n.Next.Previous = n.Previous
			n.Previous.Next = n.Next
		}
		n = n.Next
	}
 }
// поиск определенного узла
 func LookUpNode(n *Node, v int) bool {
	if root == nil {
		n = &Node{nil,v,nil}
		root = n
		return false
	}
	k := n
	for k != nil {
		if k.Value == v {
			return true
		}
	
		if k.Next == nil {
			return false
		}
		k = k.Next
	}
	return LookUpNode(k,v)
 }


 func Size(n *Node) int {
	 if n == nil {
		 return 0
	 }

	 i := 0
	 for n != nil {
		 i++
		 n = n.Next
	 }
	 return i  
 }
func main() {
	list := &Node{nil,5,nil }
	AddNode(list,5)
	AddNode(list,2)
	AddNode(list,76)
	AddNode(list,53)
	AddNode(list,8)
	Traverse(list)
	fmt.Println()
	fmt.Println(LookUpNode(list,5))
	Delete(list,76)
	Traverse(list)
}