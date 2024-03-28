package llist

import "unsafe"

/* Circular doubly linked list implementation based on the Linux one */

type Head struct {
    next, prev *Head;
}

func (head *Head) Init() {
    head.next = head;
    head.prev = head;
}

func (head *Head) IsEmpty() bool {
    return head.next == head;
}

func (head *Head) Add(n *Head) {
    head.next.prev = n;
    n.next = head.next;
    n.prev = head;
    head.next = n;
}

func (head *Head) Del(entry *Head) {
    entry.prev.next = entry.next;
    entry.next.prev = entry.prev;
}

func ListForEach[T any](head *Head) func(yield func(*T)) {
    return func(yield func(*T)) {
        for p := head.next; p != head; p = p.next {
            yield((*T)(unsafe.Pointer(p)));
        }
    }
}
