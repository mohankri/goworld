#include <iostream>
#include <string>
#include "cppsrc.h"
using namespace std;

int
cppsrc::get_key(string *key, int val)
{
	cout << __FUNCTION__ << " " << *key << "\n";
}

int
cppsrc::set_key(string *key, int val)
{
	cout << __FUNCTION__ << " " << *key << "\n";
}

/* 
int
main()
{
	string test = "test";
	cppsrc *src = new cppsrc(1);
	cppsrc *src1 = new cppsrc(1);
	src->get_key(&test, 1);
	src->set_key(&test, 1);
	test = "test1";
	src1->set_key(&test, 1);
}
*/
