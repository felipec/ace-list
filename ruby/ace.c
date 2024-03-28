#include <ruby.h>

#include "list.h"

VALUE list_class;
VALUE node_class;

struct node {
	struct list_head list;
	VALUE obj;
};

static void list_free(void *ptr)
{
	struct node *p;
	list_for_each_entry_safe(p, (struct list_head *)ptr, list)
		free(p);
	free(ptr);
}

const rb_data_type_t node_type = {
	.wrap_struct_name = "ace_node",
};

const rb_data_type_t list_type = {
	.wrap_struct_name = "ace_list",
	.function = {
		.dfree = list_free,
	},
};

VALUE rb_list_alloc(VALUE klass)
{
	struct list_head *head = malloc(sizeof(*head));
	list_init(head);
	return rb_data_typed_object_wrap(klass, head, &list_type);
}

VALUE rb_list_is_empty(VALUE self)
{
	struct list_head *head = rb_check_typeddata(self, &list_type);

	return list_is_empty(head) ? Qtrue : Qfalse;
}

VALUE rb_list_add(VALUE self, VALUE obj)
{
	struct list_head *head = rb_check_typeddata(self, &list_type);

	struct node *node = malloc(sizeof(*node));
	node->obj = obj;
	list_add(&node->list, head);

	return rb_data_typed_object_wrap(node_class, node, &node_type);
}

VALUE rb_list_del(VALUE self, VALUE rb_node)
{
	struct node *node = rb_check_typeddata(rb_node, &node_type);

	list_del(&node->list);
	free(node);

	return Qnil;
}

VALUE rb_list_each(VALUE self)
{
	struct list_head *head = rb_check_typeddata(self, &list_type);

	struct node *p;
	list_for_each_entry(p, head, list)
		rb_yield(p->obj);

	return self;
}

void Init_ace(void)
{
	VALUE mod = rb_define_module("Ace");

	node_class = rb_define_class_under(mod, "Node", rb_cObject);
	rb_undef_alloc_func(node_class);

	list_class = rb_define_class_under(mod, "List", rb_cObject);
	rb_define_alloc_func(list_class, rb_list_alloc);
	rb_define_method(list_class, "add", rb_list_add, 1);
	rb_define_method(list_class, "del", rb_list_del, 1);
	rb_define_method(list_class, "each", rb_list_each, 0);
	rb_define_method(list_class, "empty?", rb_list_is_empty, 0);
	rb_include_module(list_class, rb_mEnumerable);
}
