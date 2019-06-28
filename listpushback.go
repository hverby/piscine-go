    
package piscine



//function to insert element in the last position of the list
func ListPushBack(l *List, data interface{}) {

	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}

	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}