#include <iostream>
#define RANGE 20

using namespace std;



int main() {

	int number = 1;

	for (int i = 1; i <= RANGE; i++) {
		while (number % i != 0) {
			number++;
		}
	}

	cout << "Smallest multiple is: " << number;
}