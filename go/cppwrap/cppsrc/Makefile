all : libcpp.so
CFLAGS = -fPIC
LDFLAGS = -shared -std=c++11 -fPIC

libcpp.so : cppsrc.o
	g++ -o $@ $^ $(LDFLAGS)

cppsrc.o : cppsrc.cc
	g++ $(CFLAGS) -o $@ -c $<

clean :
	rm *.o *.so *.a


