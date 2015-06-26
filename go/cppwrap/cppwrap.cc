#include <iostream>
using namespace std;

void print()
{
	cout << "Hello World cppwrap\n";
}

extern "C" {
	void CCppwrap() {
		print();
	}
}
