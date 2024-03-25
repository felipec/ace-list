#ifndef LIST_H
#define LIST_H

#include <stddef.h>

struct list_head {
	struct list_head *next;
	struct list_head *prev;
};

#define LIST_HEAD_INIT(name) { &(name), &(name) }

static inline void list_init(struct list_head *list)
{
	list->next = list;
	list->prev = list;
}

static inline int list_is_head(const struct list_head *list, const struct list_head *head)
{
	return list == head;
}

static inline int list_is_empty(const struct list_head *head)
{
	return head->next == head;
}

static inline void list_add(struct list_head *new, struct list_head *head)
{
	struct list_head *next = head->next;

	next->prev = new;
	new->next = next;
	new->prev = head;
	head->next = new;
}

static inline void list_del(struct list_head *entry)
{
	struct list_head *prev = entry->prev;
	struct list_head *next = entry->next;

	next->prev = prev;
	prev->next = next;
}

#define container_of(ptr, type, member) ({ \
	void *__mptr = (void *)(ptr); \
	((type *)(__mptr - offsetof(type, member))); })

#define list_entry(ptr, type, member) \
	container_of(ptr, type, member)

#define list_for_each(pos, head) \
	for (struct list_head *pos = (head)->next; !list_is_head(pos, (head)); pos = pos->next)

#define list_for_each_safe(pos, head) \
	for (struct list_head *pos = (head)->next, *n = pos->next; \
		!list_is_head(pos, (head)); \
		pos = n, n = pos->next)

#define list_first_entry(ptr, type, member) \
	list_entry((ptr)->next, type, member)

#define list_entry_is_head(pos, head, member) \
	(&pos->member == (head))

#define list_next_entry(pos, member) \
	list_entry((pos)->member.next, typeof(*(pos)), member)

#define list_for_each_entry(pos, head, member) \
	for (typeof(*(pos)) *pos = list_first_entry(head, typeof(*pos), member); \
		!list_entry_is_head(pos, head, member); \
		pos = list_next_entry(pos, member))

#define list_for_each_entry_safe(pos, head, member) \
	for (typeof(*(pos)) *pos = list_first_entry(head, typeof(*pos), member), \
			*n = list_next_entry(pos, member); \
		!list_entry_is_head(pos, head, member); \
		pos = n, n = list_next_entry(n, member))

#endif
