#include <iostream>
#include <string>
#include <string.h>
#include "cppsrc.h"
#include "cppwrap.h"
using namespace std;

void print()
{
	cout << "Hello World cppwrap\n";
}

#ifdef __cplusplus
extern "C" {
#endif

Kvstore Copen(char *path, int ssize, int msize) {
	string cpp_path = path;
	cppsrc *src = new cppsrc(cpp_path, ssize, msize);
	return src;
}

int Cget_key(Kvstore store, char *key, int *val) {
	cppsrc *src = (cppsrc *)store;
	string	lkey = key;
	src->get_key(&lkey, val);
	strcpy(key, lkey.c_str());
	return 0;
}

int Cset_key(Kvstore store, char *key, int val) {
	cppsrc *src = (cppsrc *)store;
	string	lkey = key;
	src->set_key(&lkey, val);
	return 0;
}

#ifdef __cplusplus	
}

#endif
