package main
import (
	"fmt"
	"os"
	"io"
)
type Node struct{
	left *Node
	right *Node
	data int64
}
type Tree struct{
	root *Node
}
func (t *Tree) insertTree(data int64) *Tree{
	if t.root == nil {
		t.root = &Node{left: nil, right: nil, data: data}
	}else{
		t.root.insertNode(data)
	}
	return t
}
func (n *Node) insertNode(data int64){
	if n == nil {
		return
	}else if data < n.data {
		if n.left == nil {
			n.left = &Node{left: nil, right: nil, data: data}
		}else{
			n.left.insertNode(data)
		}
	}else {
		if n.right == nil {
			n.right = &Node{left: nil, right: nil, data: data}
		}else {
			n.right.insertNode(data)
		}
	}
}
func printTree(w io.Writer,node *Node,indent int , nameBranch int)  {
	if node == nil {
        return
    }
  
    for i := 0; i < indent; i++ {
        fmt.Fprint(w, "")
    }
    fmt.Fprintf(w, "%c:%v\n", nameBranch, node.data)
    printTree(w, node.left, indent+2, 'L')
    printTree(w, node.right, indent+2, 'R')
}
func main(){
	tree := &Tree{}
    tree.insertTree(100).
	insertTree(-20).
	insertTree(-50).
	insertTree(-15).
	insertTree(-60).
	insertTree(50).
	insertTree(60).
	insertTree(55).
	insertTree(85).
	insertTree(15).
	insertTree(5).
	insertTree(-10)
    printTree(os.Stdout, tree.root, 0, 'M')
}