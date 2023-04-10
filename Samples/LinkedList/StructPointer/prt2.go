package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

/*
# Above one is a body of all Node which hold value and next node address.
*/
type LinkedList struct {
	headNode *Node
}

/*
# Above struct use for create a linked list.
*/
func (l *LinkedList) AddNodeToFront(nodeValue int) {
	fmt.Println("nodeValue", nodeValue, "l.headNode", l.headNode)
	/*
		  # pointer can hold a "nil" value,that's why below second field
		    "nextNode: l.headNode[nil]"
			# When first we call this function "l.headNode" is "nil" but if you tried to
			  get another "property",you will be getting "nil pointer error".
			  Example :->
			  	node := &Node{property: nodeValue, nextNode: l.headNode}
		        Next := node.nextNode
		        fmt.Println("***node**",Next.property)
	*/
	node := &Node{property: nodeValue, nextNode: l.headNode}

	/*
		# what happen above is every time we make a new node that node hold new value
		  and previous node address.
	*/
	/*
	   # Here we declare "address" variable to "literal syntax variable" so it
	     automatically change "node" as pointer variable because it hold a "address".
	   # If the value is a pointer,we can invoke their data's directly as we're doing
	     other variable's.
	*/
	fmt.Println("node", node)
	/*
	   # Below we declare the new node to the "l.headNode".
	*/
	l.headNode = node
	/*
			#Here we "l.headNode" is a pointer as well as "node" variable is also a pointer
			 but how it works fine because "node" already hold "address or point to some
			 address" so the "l.headNode" take that point to address for itself.
			# But if "node" Didn't has a "point to address or hold address",we will get a
		 	  "Nil pointer error"
			  Example 1 :-> var node *Node
			              l.headNode = node
			Because above one doesn't "point to any address or hold any address"
			Example 2 :-> var node *Node
			            node = Node{property: nodeValue, nextNode: nil}
						l.headNode = node
			This also getting panic because the "node" hold a normal variable it doesn't
			point to any address that's why.

	*/
}

func main() {
	list := LinkedList{}

	// Add nodes to the front of the list
	list.AddNodeToFront(1)
	fmt.Println("node1", list.headNode)
	list.AddNodeToFront(2)
	fmt.Println("node2", list.headNode)
	list.AddNodeToFront(3)
	fmt.Println("node3", list.headNode)

	// Traverse the list and print the node values
	currentNode := list.headNode
	/*
		# Above one hold first or current node a value and address.
	*/
	for currentNode != nil {
		/*
		 # It will run until the currentNode come to nil value which one has that nil
		   value is last node or our first adding node or Head node.
		*/
		fmt.Println(currentNode.property)
		currentNode = currentNode.nextNode
	}
}
