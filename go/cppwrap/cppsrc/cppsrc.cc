#include <iostream>
#include <string>
#include "cppsrc.h"
using namespace std;

int
cppsrc::get_key(string *pkey, int *pval)
{
	*pkey = key;
	*pval = ksize;	
	cout << "Key Size " << ksize << "\n";
	cout << __FUNCTION__ << " " << *pkey << " val " << pval << "\n";
}

int
cppsrc::set_key(string *pkey, int pval)
{
	key = *pkey;
	ksize = pval;	
	cout << __FUNCTION__ << " " << key << " val " << pval << "\n";
}

cppsrc::cppsrc(string& path, size_t storage_size, size_t mem_size) 
{
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
