package llist

import "unsafe"

/* Circular doubly linked list implementation based on the Linux one */

type Head[T any] struct {
    next, prev *Head[T];
}

type Node[T any] struct {
    head Head[T];
    Data T;
}

func (head *Head[T]) Init() {
    head.next = head;
    head.prev = head;
}

func (head *Head[T]) IsEmpty() bool {
    return head.next == head;
}

func (head *Head[T]) Add(n *Node[T]) {
    h := &n.head;

    head.next.prev = h;
    h.next = head.next;
    h.prev = head;
    head.next = h;
}

func (head *Head[T]) Del(n *Node[T]) {
    h := &n.head;

    h.prev.next = h.next;
    h.next.prev = h.prev;
}

func (head *Head[T]) Each() func(yield func(*T)) {
    return func(yield func(*T)) {
        for p := head.next; p != head; p = p.next {
            // Pretend the *Head is a *Node
            yield(&((*Node[T])(unsafe.Pointer(p))).Data);
        }
    }
}
