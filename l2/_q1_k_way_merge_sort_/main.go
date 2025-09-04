package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// ----------------------------------------------------------------
// heap
type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].val < h[j].val } // min-heap by value
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ----------------------------------------------------------------

type Node struct {
	val  int
	next *Node
}

func newNode() *Node {
	return &Node{
		val:  0,
		next: nil,
	}
}

func merge(lists []*Node) *Node {
	h := &NodeHeap{}
	heap.Init(h)
	// push all heads to the heap
	for _, head := range lists {
		heap.Push(h, head)
	}

	var merged, tail *Node = nil, nil
	for h.Len() > 0 {
		// .() is for type assertion
		minNode := heap.Pop(h).(*Node)
		if minNode.next != nil {
			heap.Push(h, minNode.next)
		}
		if merged == nil {
			merged = minNode
			tail = merged
		} else {
			tail.next = minNode
			tail = tail.next
		}
	}

	return merged
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	var lists []*Node = make([]*Node, t)

	for i := 0; i < t; i++ {
		var size int
		fmt.Fscan(in, &size)

		var curr *Node = nil
		for ; size > 0; size-- {
			if lists[i] == nil {
				lists[i] = newNode()
				curr = lists[i]
				fmt.Fscan(in, &curr.val)
			} else {
				curr.next = newNode()
				fmt.Fscan(in, &curr.next.val)
				curr = curr.next
			}
		}
	}

	merged := merge(lists)
	var curr *Node = merged
	for curr != nil {
		fmt.Print(curr.val, " ")
		curr = curr.next
	}
}
