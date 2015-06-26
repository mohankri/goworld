#include <iostream>
#include <string>
using namespace std;

class cppsrc {
	string key;
	int    ksize;
	public:
	cppsrc(string& path, size_t storage_size, size_t mem_size);
	int set_key(string *key, int value); 
	int get_key(string *key, int *value); 
};
