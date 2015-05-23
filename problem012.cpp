#include <iostream>
#include <cmath>
#define MIN_NUM_DIVISORS 500

using namespace std;

int triangleNumber(int n) {
	return n * (n + 1) / 2;
}

int numDivisors(int n) {
	int numDivisors = 0;
	for (int i = 1; i <= (int)sqrt(n); i++) {
		if (n % i == 0) {
			numDivisors += 2;
		}
	}

	if ((int)sqrt(n) * (int)sqrt(n) == n) {
		numDivisors--;
	}

	return numDivisors;
}

int main() {
	int i = 0;
	int nDivisors = 0;
	int number = 0;
	while (nDivisors <= 500) {
		i++;
		number = triangleNumber(i);
		nDivisors = numDivisors(number);

		if (nDivisors > 500) {
			cout << triangleNumber(i) << endl;
		}
	}
}