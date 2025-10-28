package main

import "fmt"

type MyDeque struct {
	elem []int
}

func NewMyDeque() *MyDeque {
	return &MyDeque{elem: make([]int, 0)}
}

func (d *MyDeque) PushFront(value int) {
	d.elem = append([]int{value}, d.elem...)
}

func (d *MyDeque) PushBack(value int) {
	d.elem = append(d.elem, value)
}

func (d *MyDeque) PopFront() int {
	if len(d.elem) == 0 {
		return 0
	}

	value := d.elem[0]
	d.elem = d.elem[1:]
	return value
}

func (d *MyDeque) PopBack() int {
	if len(d.elem) == 0 {
		return 0
	}

	value := d.elem[len(d.elem)-1]
	d.elem = d.elem[:len(d.elem)-1]
	return value
}

func (d *MyDeque) IsEmpty() bool {
	return len(d.elem) == 0
}

func (d *MyDeque) Size() int {
	return len(d.elem)
}

func (d *MyDeque) Clear() {
	d.elem = make([]int, 0)
}

func (d *MyDeque) String() string {
	return fmt.Sprintf("MyDeque%v", d.elem)
}

func main() {
	deque := NewMyDeque()

	fmt.Println("=== Использование MyDeque ===")

	fmt.Println("\n Добавляем элементы:")
	deque.PushBack(1)
	deque.PushFront(5)
	deque.PushBack(24)
	deque.PushFront(17)

	fmt.Println(deque.String())

	fmt.Println("\n Удаляем элементы:")
	fmt.Printf("PopFront(): %d\n", deque.PopFront())
	fmt.Printf("PopBack(): %d\n", deque.PopBack())

	fmt.Println("После удаления:", deque.String())
	fmt.Printf("Размер: %d\n", deque.Size())

	fmt.Printf("IsEmpty(): %v\n", deque.IsEmpty())

	fmt.Println("\nОчищаем очередь:")
	deque.Clear()
	fmt.Printf("IsEmpty(): %v\n", deque.IsEmpty())
	fmt.Printf("Size(): %d\n", deque.Size())
	fmt.Println(deque.String())

}
