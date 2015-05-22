#include <iostream>
#define RANGE 20

using namespace std;

bool isMultiple(int n, int range) {
	for (int i = 2; i <= range; i++) {
		if (n % i != 0) {
			//cout << "Checking: " << n << endl;
			return false;
		}
	}
	return true;
}

int main() {

	int smallestMultiple = 0;
	int i = 1;

	while (!isMultiple(i, RANGE)) {
		//cout << "Checkpoint 1" << endl;
		i++;
		//cout << "checkpoint 2" << endl;
		if (isMultiple(i, RANGE)) {
			//cout << "Checkpoint 3" << endl;
			smallestMultiple = i;

			//cout << i << endl;
		}
	}

	cout << smallestMultiple << endl;
}