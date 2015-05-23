#include <iostream>
#include <limits>

#define LIMIT 1000000

using namespace std;

int collatz(long n, int chainLength) {
	if(n <= 1) {
		return chainLength;
	}

	switch (n % 2) {
		case 0: return collatz(n / 2, ++chainLength);
		break;

		case 1: return collatz(3 * n + 1, ++chainLength);
		break;

		default: return 1;
	}
}

int main() {
	int longestChain = 0;
	int startingNum = 0;

	for (int i = 0; i <= LIMIT; i++) {
		int currentChain = collatz(i, 1);

		if (currentChain > longestChain) {
			longestChain = currentChain;
			startingNum = i;
		}
	}

	cout << "The starting number " << startingNum << " produces the longest Collatz chain of length " << longestChain << endl;
}