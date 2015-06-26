#all : libcppwrap.a
all : libcppwrap.so

CFLAGS = -std=c++11 -fPIC

LDFLAGS = -shared -std=c++11 -fPIC 

libcppwrap.so : cppwrap.o
	g++ -o $@ $^ $(LDFLAGS)

libcppwrap.a : cppwrap.o
	ar -rc $@ $<

cppwrap.o : cppwrap.cc
	g++ $(CFLAGS) -o $@ -c $<	

clean :
	rm libcppwrap.a cppwrap.o libcppwrap.so
