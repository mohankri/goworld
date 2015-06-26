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
	cout << "Copen called... " << src << "\n";
	return src;
}

int Cget_key(Kvstore store, char *key, int val) {
	cout << "Initialized as ... " << key << " val " << val << "\n";
	cppsrc *src = (cppsrc *)store;
	string	lkey = key;
	src->get_key(&lkey, &val);
	strcpy(key, lkey.c_str());
	cout << "Get_key called... " << key << " val " << val << "\n";
}

int Cset_key(Kvstore store, char *key, int val) {
	//cout << "Set_key called... " << key << store << "\n";
	cppsrc *src = (cppsrc *)store;
	string	lkey = key;
	cout << "Set Key " << lkey << " val " << val << "\n";
	src->set_key(&lkey, val);
}

#ifdef __cplusplus	
}

#endif
