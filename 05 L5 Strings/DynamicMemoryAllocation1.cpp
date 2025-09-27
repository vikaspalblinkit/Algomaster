#include<iostream>
using namespace std; 


int main() {

	// Static Allocation maintain by symbolic table 
	int b[100];
	cout << sizeof(b);

	cout << endl;

	// Dynamic Memory Allocations 
	// Here you get the ptr in static memory which is pointed to heap memory arrays 
	int* ptr = new int[100];
	cout << sizeof(ptr); 


	// Memory leaks = when the reference of the objects is lost then it is called the memory leaks
	delete [] ptr;
	cout << endl;
	cout << sizeof(ptr);
	return 0;
}

/* DP on strings questions */ 