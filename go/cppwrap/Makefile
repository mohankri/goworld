#all : libcppwrap.a
all : libcppwrap.so

#CFLAGS = -std=c++11 -fPIC
CFLAGS = -fPIC -I./cppsrc

GOPATH = $(HOME)/goworld/go

LDFLAGS = -shared -std=c++11 -fPIC -Wl,-rpath,$(GOPATH)/cppwrap/cppsrc -L./cppsrc -lcpp

libcppwrap.so : cppwrap.o
	g++ -o $@ $^ $(LDFLAGS)

libcppwrap.a : cppwrap.o
	ar -rc $@ $<

cppwrap.o : cppwrap.cc
	g++ $(CFLAGS) -o $@ -c $<	

clean :
	rm libcppwrap.a cppwrap.o libcppwrap.so
