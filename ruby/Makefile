CC := gcc

CFLAGS := -O3 -Wall -Wextra -Wno-unused-parameter

ruby_version := $(shell ruby -e 'require "rbconfig"; puts RbConfig::CONFIG["RUBY_API_VERSION"]')
RUBY_CFLAGS := $(shell pkg-config --cflags ruby-$(ruby_version))
RUBY_LIBS := $(shell pkg-config --libs ruby-$(ruby_version))

override CFLAGS += -I../c

all: ace.so

ace.so: CFLAGS += -fPIC $(RUBY_CFLAGS)
ace.so: ace.o

%.so::
	$(CC) -shared $(LDFLAGS) -o $@ $^ $(LIBS)

ace.o: ../c/list.h
