package linklist

import (
	"iter"
)

type Node[V any] struct {
	data     V
	prevNode *Node[V]
	nextNode *Node[V]
}

type List[S any] struct {
	root         *Node[S]
	defaultValue int
	size         int
}

func NewList[S any]() *List[S] {
	var def S
	dummyNode := &List[S]{
		root:         newNode[S](def),
		defaultValue: 0,
		size:         0,
	}

	return dummyNode
}

func (l *List[S]) Size() int {
	return l.size
}

func (l *List[S]) PushFront(data S) error {
	newN := newNode[S](data)
	setNewNode(newN, l.root.nextNode)
	l.size++
	return nil
}

func (l *List[S]) PushBack(data S) error {
	newN := newNode(data)
	setNewNode(newN, l.root)
	l.size++
	return nil
}

func (l *List[S]) Front() *Node[S] {
	if(l.root.nextNode == l.root){
		return nil
	}
	return l.root.nextNode
}

func (l *List[S]) Remove(n *Node[S]) S{
	if(n == l.root || n == nil){
		panic("Error node")
	}
	n.prevNode.nextNode = n.nextNode
	n.nextNode.prevNode = n.prevNode
	n.nextNode = nil
	n.prevNode = nil
	l.size--

	return n.data
}

func (l *List[S]) MoveToBack(n *Node[S]){
	if(n == l.root || n == l.root.prevNode){
		return
	}
	n.prevNode.nextNode = n.nextNode
	n.nextNode.prevNode = n.prevNode
	setNewNode(n, l.root)
}

func (l *List[S]) MoveToFront(n *Node[S]){
	if(n == l.root || n == l.root.nextNode){
		return
	}
	n.prevNode.nextNode = n.nextNode
	n.nextNode.prevNode = n.prevNode
	setNewNode(n, l.root.nextNode)
}

func (l* List[S]) Back() *Node[S]{
	if(l.root.prevNode == l.root){
		return nil
	}
	return l.root.prevNode
}

func (l *List[S]) PutNewValue(n *Node[S], data S) {
	n.data = data
}

func (l *List[S]) All() iter.Seq[S]{
	return func(yield func(S) bool) {
		it := l.root.nextNode
		for ; it != l.root; it = it.nextNode {
			if !yield(it.data){
				return
			}
		}
	}
}

func newNode[S any](data S) *Node[S] {
	n := &Node[S]{
		data:     data,
		prevNode: nil,
		nextNode: nil,
	}
	n.nextNode = n
	n.prevNode = n
	return n
}

func (n *Node[V]) GetData() V{
	return n.data
}

func setNewNode[S any](newNode, oldNode *Node[S]) {
	oldNode.prevNode.nextNode = newNode
	newNode.prevNode = oldNode.prevNode
	oldNode.prevNode = newNode
	newNode.nextNode = oldNode
}

// func(n *Node[S]) Print(){
// 	fmt.Println(n.data)
// }