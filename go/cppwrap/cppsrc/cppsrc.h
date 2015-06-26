#include <iostream>
#include <string>
using namespace std;

class cppsrc {
	
	public:
	cppsrc(int arg1) {
		cout << "constructor " << arg1 << "\n";
	}

	int set_key(string *key, int value); 
	int get_key(string *key, int value); 
};
