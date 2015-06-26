#include <iostream>
#include <string>
#include "cppsrc.h"
using namespace std;

void print()
{
	cout << "Hello World cppwrap\n";
}

extern "C" {
	void CCppwrap() {
		cppsrc *src = new cppsrc(1);
		cppsrc *src1 = new cppsrc(1);
		string	key = "Krishna ";
		src->set_key(&key, 1);
		key = "Mohan";
		src1->set_key(&key, 1);
	}
}
