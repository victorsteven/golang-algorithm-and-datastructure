package singlylinkedlists

import "errors"

var (
	errEmptyList        = errors.New("list is empty")
	errIndexOutOfBounds = errors.New("index is out of bounds")
)

//SSLNode singly linked list node
type SLLNode struct {
	Value interface{}
	Next  *SLLNode
}
type SinglylinkedList struct {
	count int
	head  *SLLNode
	tail  *SLLNode
}

//New factory to generate new singly linked lists
func New(values ...interface{}) *SinglylinkedList {
	list := SinglylinkedList{}
	list.Add(values...)
	return &list
}

//isEmpty checks if the list is empty
func (s *SinglylinkedList) IsEmpty() bool {
	return s.Size() == 0
}

//Size returns size of the list
func (s *SinglylinkedList) Size() int {
	return s.count
}

//Clear clears list
func (s *SinglylinkedList) Clear() {
	current := s.head
	for current != nil {
		temp := current
		current = current.Next
		temp.Next = nil
		temp = nil
	}
	current = nil
	s.count = 0
	s.head = nil
	s.tail = nil
}

//GetValues returns values
func (s *SinglylinkedList) GetValues() []interface{} {

	values := make([]interface{}, 0, s.Size())
	current := s.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	current = nil
	return values
}

//GetIndexOf returns the index of the first occurrence
func (s *SinglylinkedList) GetIndexOf(value interface{}) int {
	index := 0
	current := s.head
	for current != nil {
		if value == current.Value {
			return index
		}
		current = current.Next
		index++
	}
	current = nil
	return -1
}

//Add add to the list
func (s *SinglylinkedList) Add(values ...interface{}) {
	for _, value := range values {
		s.InsertAt(s.Size(), value)
	}
}

//InsertAt insert value at specific index in the list
func (s *SinglylinkedList) InsertAt(index int, value interface{}) error {
	if index < 0 || index > s.count {
		return errIndexOutOfBounds
	}
	element := &SLLNode{Value: value}
	if s.IsEmpty() {
		s.head = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	if index == 0 {
		element.Next = s.head
		s.head = element
		s.count++
		element = nil
		return nil
	}

	if index == s.count {
		s.tail.Next = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	current, err := s.getNode(index - 1)
	if err != nil {
		return err
	}

	element.Next = current.Next
	current.Next = element
	s.count++
	current = nil
	element = nil
	return nil
}

//RemoveAt remove value from the list at specific index
func (s *SinglylinkedList) RemoveAt(index int) (interface{}, error) {
	if index < 0 || index >= s.count {
		return nil, errIndexOutOfBounds
	}
	if s.count == 1 {
		value := s.head.Value
		s.head = nil
		s.tail = nil
		s.count--
		return value, nil
	}

	if index == 0 {
		current := s.head
		s.head = current.Next
		current.Next = nil
		value := current.Value
		s.count--
		current = nil
		return value, nil
	}

	current, err := s.getNode(index - 1)
	if err != nil {
		return nil, err
	}

	value := current.Next.Value
	current.Next = current.Next.Next
	if index == s.count-1 {
		s.tail = current
	}
	s.count--
	current = nil
	return value, nil
}

func (s *SinglylinkedList) getNode(index int) (*SLLNode, error) {
	if index < 0 || index >= s.count {
		return nil, errIndexOutOfBounds
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current, nil
}

//GetValueAt returns vlaue at a specific index in the list
func (s *SinglylinkedList) GetValueAt(index int) (interface{}, error) {
	current, err := s.getNode(index)
	if err != nil {
		return nil, err
	}
	return current.Value, nil
}

//GetFirstValue returns first value in the list
func (s *SinglylinkedList) GetFirstValue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyList
	}
	return s.head.Value, nil
}

//GetLastValue returns last value in the list
func (s *SinglylinkedList) GetLastValue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyList
	}
	return s.tail.Value, nil
}

//GetHead returns head node of the list
func (s *SinglylinkedList) GetHead() *SLLNode {
	return s.head
}
