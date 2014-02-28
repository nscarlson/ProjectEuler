#include <iostream>

#define LIMIT 1000
using namespace std;

int main() {
	for (int a = 1; a < LIMIT; a++) {
		for (int b = a + 1; b < LIMIT; b++) {
			for (int c = b + 1; c < LIMIT; c++) {
				if (a + b + c == LIMIT && a * a + b * b == c * c) {
					cout << "a = " << a << ", " << "b = " << b << ", " << "c = " << c << endl;
					cout << "The product abc where a, b and c make a pythagorean triplet is: ";
					cout << a * b * c << endl;
				}
			}
		}
	}
}