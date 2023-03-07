package main

type Node struct {
	data       int
	prev, next *Node
}

// DuLinkList double linked list
type DuLinkList struct {
	first *Node
	last  *Node
	len   int
}

func NewDuLinkList() *DuLinkList {
	return new(DuLinkList).Init()
}

func (n *Node) Prev() *Node {
	return n.prev
}
func (n *Node) Next() *Node {
	return n.next
}
func (l *DuLinkList) First() *Node {
	return l.first
}
func (l *DuLinkList) Front() *Node {
	return l.First()
}
func (l *DuLinkList) Last() *Node {
	return l.last
}
func (l *DuLinkList) Back() *Node {
	return l.Last()
}
func (l *DuLinkList) Len() int {
	return l.len
}
func (l *DuLinkList) Init() *DuLinkList {
	l.first = nil
	l.len = 0
	l.last = nil
	return l
}
func (l *DuLinkList) PushFrontList(lf *DuLinkList) *DuLinkList {
	l.first.prev = lf.last
	lf.last.next = l.first
	l.first = lf.first
	l.len += lf.len
	return l
}

func (l *DuLinkList) PushBackList(lf *DuLinkList) *DuLinkList {
	l.last.next = lf.first
	lf.first.prev = l.last
	l.last = lf.last
	l.len += lf.len
	return l
}
func (l *DuLinkList) PushBack(x int) *Node {
	s := new(Node)
	s.data = x
	if l.len != 0 {
		l.last.next = s
	} else {
		l.first = s
	}

	s.prev = l.last
	l.last = s
	l.len++
	return s
}

func (l *DuLinkList) PushFront(x int) *Node {
	s := new(Node)
	s.data = x
	if l.len != 0 {
		l.first.prev = s
	} else {
		l.last = s
	}
	s.next = l.first
	l.first = s
	l.len++
	return s
}

func (l *DuLinkList) Remove(n *Node) *DuLinkList {
	n.prev.next = n.next
	if n.next == nil {
		l.last = n.prev
	} else if n.prev == nil {
		l.first = n.next
	}
	l.len--
	return l
}

func (l *DuLinkList) PopLast() {
	l.last = l.last.prev
	l.last.next = nil
	l.len--
}

func (l *DuLinkList) PopFirst() {
	l.first = l.first.next
	l.first.prev = nil
	l.len--

}

func (l *DuLinkList) move(n, des *Node) {
	// des destination 目标位置
	if n == des {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev

	n.prev = des
	n.next = des.next
	n.prev.next = n
	n.next.prev = n
}

func (l *DuLinkList) MoveToFront(n *Node) {
	if l.first == n {
		return
	}
	l.move(n, l.first)
}

func (l *DuLinkList) MoveToBack(n *Node) {
	if l.last == n {
		return
	}
	l.move(n, l.last)
}

func (l *DuLinkList) MoveBefore(n, des *Node) {
	l.move(n, des.prev)
}

func (l *DuLinkList) MoveAfter(n, des *Node) {
	l.move(n, des)
}

func (l *DuLinkList) Insert(n, des *Node) *Node {
	n.prev = des
	n.next = des.next
	n.prev.next = n
	n.next.prev = n
	l.len++
	return n
}
