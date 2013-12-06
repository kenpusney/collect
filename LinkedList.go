package collect

type LinkedList struct {
	head   *node
	length int
}

type node struct {
	value interface{}
	next  *node
}

type LinkedListIterator struct {
	current int
	list    *LinkedList
}

func (l *LinkedList) GetIterator() Iterator {
	return &LinkedListIterator{list: l}
}

func (l *LinkedList) Append(v interface{}) bool {
	if l.head == nil {
		l.head = &node{value: v}
		l.length++
		return true
	} else {
		this := l.head
		for this.next != nil {
			this = this.next
		}
		this.next = &node{value: v}
		l.length++
		return true
	}
	return false
}

func (l *LinkedList) Get(i int) interface{} {
	if i < l.length {
		this := l.head
		for i != 0 {
			this = this.next
			i--
		}
		return this.value
	}
	return nil
}

func (l *LinkedList) Put(i int, v interface{}) bool {
	if i < l.length {
		this := l.head
		for i != 1 {
			this = this.next
			i--
		}
		temp := &node{value: v}
		temp.next = this.next
		this.next = temp
		return true
	}
	return false
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedListIterator) HasNext() bool {
	return l.current < l.list.Length()
}

func (l *LinkedListIterator) Next() interface{} {
	value := l.list.Get(l.current)
	l.current++
	return value
}
