package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	count int
	head  *ListItem
	tail  *ListItem
}

func (l list) Len() int {
	return l.count
}

func (l list) Front() *ListItem {
	if l.count == 0 {
		return nil
	}
	return l.head
}

func (l list) Back() *ListItem {
	if l.count == 0 {
		return nil
	}
	return l.tail
}

func (l *list) addHead(node *ListItem) {
	if node == nil {
		return
	}
	if l.head != nil {
		l.head.Prev = node
		node.Next = l.head
	}
	if l.tail == nil {
		l.tail = node
	}
	l.head = node
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := &ListItem{
		Value: v,
	}
	l.addHead(node)
	l.count++
	return l.head
}

func (l *list) addBack(node *ListItem) {
	if node == nil {
		return
	}
	if l.tail != nil {
		l.tail.Next = node
		node.Prev = l.tail
	}
	if l.head == nil {
		l.head = node
	}
	l.tail = node
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := &ListItem{
		Value: v,
	}
	l.addBack(node)
	l.count++
	return l.tail
}

func (l *list) remove(i *ListItem) {
	if i == nil {
		return
	}
	switch {
	case i == l.head:
		l.head = l.head.Next
		if l.head != nil {
			l.head.Prev = nil
		}
	case i == l.tail:
		l.tail = l.tail.Prev
		if l.tail != nil {
			l.tail.Next = nil
		}
	case i != l.head && i != l.tail:
		prev := i.Prev
		next := i.Next
		prev.Next, next.Prev = next, prev
	default:
		return
	}
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	l.remove(i)
	l.count--
}

func (l *list) moveToHead(i *ListItem) {
	l.remove(i)
	l.addHead(i)
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil {
		return
	}
	l.moveToHead(i)
}

func NewList() List {
	return new(list)
}
